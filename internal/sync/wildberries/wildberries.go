package wildberries

import (
	"context"
	"github.com/pkg/errors"
	"github.com/shoppigram-com/marketplace-api/packages/logger"
	"github.com/shoppigram-com/marketplace-api/packages/wildberries/contentapi"
	"github.com/shoppigram-com/marketplace-api/packages/wildberries/pricesapi"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type (
	Config struct {
		RunInterval time.Duration
	}
	// Syncer is a runner process for syncing our system with WildBerries
	Syncer struct {
		contentAPI *contentapi.APIClient
		pricesAPI  *pricesapi.APIClient
		repo       Repository
		l          *zap.Logger
		conf       Config
		ctx        context.Context
		cancelFunc context.CancelFunc
	}
)

// New returns a new Syncer instance
func New(r Repository, l *zap.Logger) *Syncer {
	ctx, cancel := context.WithCancel(context.Background())

	contentConf := contentapi.NewConfiguration()
	contentConf.HTTPClient = &http.Client{
		Timeout: time.Second * 10,
	}

	pricesConf := pricesapi.NewConfiguration()
	pricesConf.HTTPClient = &http.Client{
		Timeout: time.Second * 10,
	}

	return &Syncer{
		contentAPI: contentapi.NewAPIClient(contentConf),
		pricesAPI:  pricesapi.NewAPIClient(pricesConf),
		repo:       r,
		ctx:        ctx,
		cancelFunc: cancel,
		l:          l,
	}
}

// Shutdown cancels the context of the runner to stop it
func (r *Syncer) Shutdown() {
	r.cancelFunc()
	<-r.ctx.Done()
}

// Sync launches the Syncer
func (r *Syncer) Sync() error {
	for {
		select {
		case <-r.ctx.Done():
			return nil
		default:
			job, err := r.repo.GetNextSyncJob(r.ctx)
			if err != nil {
				return errors.Wrap(err, "r.repo.GetNextSyncJob")
			}
			if job == nil {
				time.Sleep(timeout)
				continue
			}

			err = r.sync(*job)
			if err != nil {
				r.l.Error("failed wildberries sync", logger.SilentError(err))
				err2 := r.repo.SetSyncFailure(r.ctx, SetSyncFailureParams{
					JobID:     job.SyncJobID,
					LastError: err.Error(),
				})
				if err2 != nil {
					return errors.Wrap(err2, "r.repo.SetSyncFailure")
				}
				continue
			}

			err = r.repo.SetSyncSuccess(r.ctx, SetSyncSuccessParams{
				JobID: job.SyncJobID,
			})
			if err != nil {
				return errors.Wrap(err, "r.repo.SetSyncSuccess")
			}

			r.l.Info("wildberries sync success", zap.String("shop_id", job.ShopID.String()))
		}
	}
}

func (r *Syncer) sync(shop Job) error {
	sortAscPtr := new(bool)
	*sortAscPtr = true

	withPhotoPtr := new(int32)
	*withPhotoPtr = -1

	cards, _, err := r.contentAPI.DefaultApi.
		ContentV2GetCardsListPost(
			context.WithValue(
				r.ctx,
				contentapi.ContextAPIKeys,
				map[string]contentapi.APIKey{
					"HeaderApiKey": {Key: shop.APIKey},
				},
			)).
		ContentV2GetCardsListPostRequest(
			contentapi.ContentV2GetCardsListPostRequest{
				Settings: &contentapi.ContentV2GetCardsListPostRequestSettings{
					Sort: &contentapi.ContentV2GetCardsListPostRequestSettingsSort{Ascending: sortAscPtr},
					Filter: &contentapi.ContentV2GetCardsListPostRequestSettingsFilter{
						WithPhoto: withPhotoPtr,
					},
					Cursor: &contentapi.ContentV2GetCardsListPostRequestSettingsCursor{
						Limit: &fetchLimit,
					},
				}},
		).
		Execute()
	if err != nil {
		return errors.Wrap(err, "r.contentAPI.DefaultApi.ContentV2CardsErrorListGet")
	}

	extProductsMap := make(map[string]ExternalProduct)
	for _, card := range cards.Cards {
		var p ExternalProduct

		id := card.GetNmID()
		p.ExternalID = strconv.Itoa(int(id))
		p.Name = card.GetTitle()
		p.Category = card.GetSubjectName()
		p.Description = card.GetDescription()

		for _, photo := range card.GetPhotos() {
			p.Photos = append(p.Photos, Photo{
				URL: photo.GetBig(),
			})
		}

		extProductsMap[p.ExternalID] = p
	}

	goods, _, err := r.pricesAPI.DefaultApi.
		ApiV2ListGoodsFilterGet(
			context.WithValue(
				context.Background(),
				pricesapi.ContextAPIKeys,
				map[string]pricesapi.APIKey{
					"HeaderApiKey": {Key: shop.APIKey},
				},
			),
		).
		Limit(fetchLimit).
		Execute()
	if err != nil {
		return errors.Wrap(err, "r.pricesAPI.DefaultApi.ApiV2ListGoodsFilterGet")
	}

	// For now, we will get the average price of all sizes
	for _, good := range goods.Data.ListGoods {
		id := strconv.Itoa(int(good.GetNmID()))
		var (
			sum, count, price int32
		)
		for _, s := range good.Sizes {
			sum += s.GetPrice()
			count++
		}

		if count > 0 {
			price = sum / count
		}

		currency := strings.ToLower(good.GetCurrencyIsoCode4217())
		if currency != supportedCurrency {
			return errors.Errorf("currency %s is not supported", currency)
		}

		if p, ok := extProductsMap[id]; ok {
			p.Price = float64(price)

			// Add WildBerries link
			p.ExternalLinks = append(p.ExternalLinks, ExternalLink{
				URL:   makeWBProductLink(id),
				Label: "WildBerries",
			})

			extProductsMap[id] = p
		}
	}

	var extProducts []ExternalProduct
	for _, p := range extProductsMap {
		extProducts = append(extProducts, p)
	}

	err = r.repo.SyncProducts(r.ctx, SetProductsParams{
		ShopID:           shop.ShopID,
		ExternalProvider: externalProvider,
		Products:         extProducts,
	})
	if err != nil {
		return errors.Wrap(err, "r.repo.SyncProducts")
	}

	return nil
}

func makeWBProductLink(id string) string {
	return "https://www.wildberries.ru/catalog/" + id + "/detail.aspx"
}

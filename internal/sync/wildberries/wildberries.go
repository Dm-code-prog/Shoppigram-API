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
	// Syncer is a runner process for syncing our system with WildBerries
	Syncer struct {
		contentAPI *contentapi.APIClient
		pricesAPI  *pricesapi.APIClient
		repo       Repository
		l          *zap.Logger
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
func (s *Syncer) Shutdown(_ error) {
	s.cancelFunc()
	<-s.ctx.Done()
}

// Sync launches the Syncer
func (s *Syncer) Sync() error {
	exec := func() (err error) {
		defer func(l *zap.Logger) {
			_ = l.Sync()
		}(s.l)

		job, err := s.repo.GetNextSyncJob(s.ctx)
		if err != nil {
			return errors.Wrap(err, "s.repo.GetNextSyncJob")
		}
		if job == nil {
			time.Sleep(timeout)
			return nil
		}

		err = s.sync(*job)
		if err != nil {
			s.l.Error("failed wildberries sync", logger.SilentError(err))
			err2 := s.repo.SetSyncFailure(s.ctx, SetSyncFailureParams{
				JobID:     job.SyncJobID,
				LastError: err.Error(),
			})
			if err2 != nil {
				return errors.Wrap(err2, "s.repo.SetSyncFailure")
			}
			return nil
		}

		err = s.repo.SetSyncSuccess(s.ctx, SetSyncSuccessParams{
			JobID: job.SyncJobID,
		})
		if err != nil {
			return errors.Wrap(err, "s.repo.SetSyncSuccess")
		}

		s.l.Info("wildberries sync success", zap.String("shop_id", job.ShopID.String()))
		return nil
	}

	for {
		select {
		case <-s.ctx.Done():
			return nil
		default:
			err := exec()
			if err != nil {
				return err
			}
		}

	}
}

func (s *Syncer) sync(shop Job) error {
	cards, err := s.getCards(shop.APIKey)
	if err != nil {
		return errors.Wrap(err, "s.getCards")
	}

	extProductsMap := make(map[string]Product)
	for _, card := range cards.Cards {
		var p Product

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

	goods, err := s.getGoods(shop.APIKey)
	if err != nil {
		return errors.Wrap(err, "s.getGoods")
	}

	// For now, we will get the average price of all sizes
	for _, good := range goods.Data.ListGoods {
		currency := strings.ToLower(good.GetCurrencyIsoCode4217())
		if currency != supportedCurrency {
			return errors.Errorf("currency %s is not supported", currency)
		}

		id := strconv.Itoa(int(good.GetNmID()))

		var variants []Variant
		for _, s := range good.Sizes {
			variants = append(variants, Variant{
				Price:           float64(s.GetPrice()),
				DiscountedPrice: float64(s.GetDiscountedPrice()),
				Dimensions: map[string]string{
					"size": s.GetTechSizeName(),
				},
			})
		}

		if p, ok := extProductsMap[id]; ok {
			// Add WildBerries link
			p.ExternalLinks = append(p.ExternalLinks, ExternalLink{
				URL:   makeWBProductLink(id),
				Label: "WildBerries",
			})

			// Add variants
			p.Variants = variants

			extProductsMap[id] = p
		}
	}

	var extProducts []Product
	for _, p := range extProductsMap {
		extProducts = append(extProducts, p)
	}

	err = s.repo.SyncProducts(s.ctx, SetProductsParams{
		ShopID:           shop.ShopID,
		ExternalProvider: externalProvider,
		Products:         extProducts,
	})
	if err != nil {
		return errors.Wrap(err, "s.repo.SyncProducts")
	}

	return nil
}

func (s *Syncer) getCards(apiKey string) (*contentapi.ContentV2GetCardsListPost200Response, error) {
	sortAscPtr := new(bool)
	*sortAscPtr = true

	withPhotoPtr := new(int32)
	*withPhotoPtr = -1

	cards, r, err := s.contentAPI.DefaultApi.
		ContentV2GetCardsListPost(
			context.WithValue(
				s.ctx,
				contentapi.ContextAPIKeys,
				map[string]contentapi.APIKey{
					"HeaderApiKey": {Key: apiKey},
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
		return nil, errors.Wrap(err, "s.contentAPI.DefaultApi.ContentV2CardsErrorListGet")
	}

	if r.StatusCode != http.StatusOK {
		return nil, errors.Errorf("unexpected status code %d", r.StatusCode)
	}

	return cards, nil
}

func (s *Syncer) getGoods(apiKey string) (*pricesapi.ApiV2ListGoodsFilterGet200Response, error) {
	goods, r, err := s.pricesAPI.DefaultApi.
		ApiV2ListGoodsFilterGet(
			context.WithValue(
				s.ctx,
				pricesapi.ContextAPIKeys,
				map[string]pricesapi.APIKey{
					"HeaderApiKey": {Key: apiKey},
				},
			),
		).
		Limit(fetchLimit).
		Execute()
	if err != nil {
		return nil, errors.Wrap(err, "s.pricesAPI.DefaultApi.ApiV2ListGoodsFilterGet")
	}

	if r.StatusCode != http.StatusOK {
		return nil, errors.Errorf("unexpected status code %d", r.StatusCode)
	}

	return goods, nil
}

func makeWBProductLink(id string) string {
	return "https://www.wildberries.ru/catalog/" + id + "/detail.aspx"
}

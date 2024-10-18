package wildberries

import (
	"context"
	"github.com/pkg/errors"
	"github.com/shoppigram-com/marketplace-api/packages/wildberries/contentapi"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

const (
	cursorName = "wildberries_sync"

	externalProvider = "wildberries"
)

type (
	// RunnerConf is the configuration for the Runner
	// that syncs shop data with WildBerries
	RunnerConf struct {
		APIHost     string
		RunInterval time.Duration
	}

	// Runner is a runner process for syncing our system with WildBerries
	Runner struct {
		contentAPI *contentapi.APIClient
		repo       Repository
		l          *zap.Logger
		conf       RunnerConf
		ctx        context.Context
		cancelFunc context.CancelFunc
	}
)

// New returns a new Runner instance
func New(r Repository, conf RunnerConf, l *zap.Logger) *Runner {
	ctx, cancel := context.WithCancel(context.Background())

	contentConf := contentapi.NewConfiguration()
	contentConf.Host = conf.APIHost
	contentConf.HTTPClient = &http.Client{
		Timeout: 10,
	}

	return &Runner{
		contentAPI: contentapi.NewAPIClient(contentConf),
		repo:       r,
		ctx:        ctx,
		cancelFunc: cancel,
		l:          l,
	}
}

// Shutdown cancels the context of the runner to stop it
func (r *Runner) Shutdown() {
	r.cancelFunc()
}

// Run launches the Runner
func (r *Runner) Run() error {
	ticker := time.NewTicker(r.conf.RunInterval)

	for {
		select {
		case <-r.ctx.Done():
			return nil
		case <-ticker.C:
			cursor, err := r.repo.GetCursor(
				r.ctx,
				GetCursorParams{Name: cursorName},
			)
			if err != nil {
				return errors.Wrap(err, "r.repo.GetCursor")
			}

			if cursor == nil {
				err := r.repo.ResetCursor(
					r.ctx,
					ResetCursorParams{
						Name: cursorName,
					},
				)
				if err != nil {
					return errors.Wrap(err, "r.repo.ResetCursor")
				}

				r.l.Info("cursor reset")
				continue
			}

			shop, err := r.repo.GetNextShop(r.ctx, *cursor)
			if err != nil {
				return errors.Wrap(err, "r.repo.GetNextShop")
			}

			err = r.sync(shop)
			if err != nil {
				return errors.Wrap(err, "r.sync")
			}

			err = r.repo.SetCursor(
				r.ctx,
				Cursor{
					Name:      cursorName,
					ID:        shop.ID,
					Timestamp: shop.CursorTimestamp,
				},
			)
		}
	}
}

func (r *Runner) sync(shop NextShop) error {
	auth := context.WithValue(
		r.ctx,
		contentapi.ContextAPIKeys,
		map[string]contentapi.APIKey{
			"HeaderApiKey": {Key: shop.APIKey},
		},
	)

	cards, _, err := r.contentAPI.DefaultApi.
		ContentV2GetCardsListPost(auth).
		Execute()
	if err != nil {
		return errors.Wrap(err, "r.contentAPI.DefaultApi.ContentV2CardsErrorListGet")
	}

	var extProducts []ExternalProduct

	photosMap := make(map[string][]contentapi.ContentV2GetCardsListPost200ResponseCardsInnerPhotosInner)

	for _, card := range cards.Cards {
		var p ExternalProduct

		id := card.GetNmID()
		p.ExternalID = strconv.Itoa(int(id))
		p.Name = card.GetTitle()
		p.Category = card.GetSubjectName()
		p.Description = card.GetDescription()

		photosMap[p.ExternalID] = card.GetPhotos()

		extProducts = append(extProducts, p)
	}

	// TODO: request the price from the price and discounts API

	err = r.repo.SetExternalProducts(r.ctx, SetProductsParams{
		ShopID:           shop.ID,
		ExternalProvider: externalProvider,
		Products:         extProducts,
	})
	if err != nil {
		return errors.Wrap(err, "r.repo.SetExternalProducts")
	}

	_, err = r.repo.GetProducts(r.ctx, GetProductsParams{
		ShopID:           shop.ID,
		ExternalProvider: externalProvider,
	})
	if err != nil {
		return errors.Wrap(err, "r.repo.GetProducts")
	}

	return nil
}

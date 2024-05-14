package notifications

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/shoppigram-com/marketplace-api/internal/notifications/generated"
)

// Pg implements the Repository interface
// using PostgreSQL as the backing store.
type Pg struct {
	gen                      *generated.Queries
	encryptionKey            string
	newOrderFetchLimit       int
	newMarketplaceFetchLimit int
}

// NewPg creates a new Pg
func NewPg(db *pgxpool.Pool, encryptionKey string, newOrderFetchLimit int, newMarketplaceFetchLimit int) *Pg {
	return &Pg{
		gen:                      generated.New(db),
		encryptionKey:            encryptionKey,
		newOrderFetchLimit:       newOrderFetchLimit,
		newMarketplaceFetchLimit: newMarketplaceFetchLimit,
	}
}

// GetAdminsNotificationList gets a list of admins to notify about an order
func (p *Pg) GetAdminsNotificationList(ctx context.Context, webAppID uuid.UUID) ([]int64, error) {
	var adminsNotificationList []int64

	rows, err := p.gen.GetAdminsNotificationList(ctx, pgtype.UUID{
		Bytes: webAppID,
		Valid: true,
	})
	if err != nil {
		return nil, errors.Wrap(err, "p.gen.GetAdminsNotificationList")
	}

	for _, v := range rows {
		adminsNotificationList = append(adminsNotificationList, v.AdminChatID)
	}

	return adminsNotificationList, nil
}

// GetReviewersNotificationList gets a list of reviewers to notify about a new marketplace
func (p *Pg) GetReviewersNotificationList(ctx context.Context, webAppID uuid.UUID) ([]int64, error) {
	var reviewersNotificationList []int64

	rows, err := p.gen.GetReviewersNotificationList(ctx, webAppID)
	if err != nil {
		return nil, errors.Wrap(err, "p.gen.GetReviewersNotificationList")
	}

	for _, v := range rows {
		reviewersNotificationList = append(reviewersNotificationList, v)
	}

	return reviewersNotificationList, nil
}

// GetNotifierCursor gets notifier cursor
func (p *Pg) GetNotifierCursor(ctx context.Context, name string) (Cursor, error) {
	cursor, err := p.gen.GetNotifierCursor(ctx, pgtype.Text{
		String: name,
		Valid:  true,
	})
	if err != nil {
		return Cursor{}, errors.Wrap(err, "p.gen.GetNotifierCursor")
	}
	return Cursor{
		CursorDate:      cursor.CursorDate.Time,
		LastProcessedID: cursor.LastProcessedID.Bytes,
	}, nil
}

// UpdateNotifierCursor updates notifier cursor
func (p *Pg) UpdateNotifierCursor(ctx context.Context, cur Cursor) error {
	err := p.gen.UpdateNotifierCursor(ctx, generated.UpdateNotifierCursorParams{
		Name: pgtype.Text{
			String: cur.Name,
			Valid:  true,
		},
		CursorDate: pgtype.Timestamp{
			Time:  cur.CursorDate,
			Valid: true,
		},
		LastProcessedID: pgtype.UUID{
			Bytes: cur.LastProcessedID,
			Valid: true,
		},
	})
	if err != nil {
		return errors.Wrap(err, "p.gen.UpdateNotifierCursor")
	}
	return nil
}

// GetNotificationsForNewOrdersAfterCursor gets notifcations for orders which were
// created after date specified in cursor
func (p *Pg) GetNotificationsForNewOrdersAfterCursor(ctx context.Context, cur Cursor) ([]NewOrderNotification, error) {
	var newOrderNotifications []NewOrderNotification

	rows, err := p.gen.GetNotificationsForNewOrdersAfterCursor(
		ctx,
		generated.GetNotificationsForNewOrdersAfterCursorParams{
			CreatedAt: pgtype.Timestamp{
				Time:  cur.CursorDate,
				Valid: true,
			},
			Limit: int32(p.newOrderFetchLimit),
		})
	if err != nil {
		return nil, errors.Wrap(err, "p.gen.GetNotificationsForNewOrdersAfterCursor")
	}

	ordersMap := map[string]NewOrderNotification{}

	for _, r := range rows {
		orderID := r.OrderID.String()
		if order, ok := ordersMap[orderID]; ok {
			// If the order exists, append the new product to the existing order's product list
			order.Products = append(order.Products, Product{
				Name:          r.Name,
				Quantity:      int(r.Quantity),
				Price:         r.Price,
				PriceCurrency: r.PriceCurrency,
			})
			// Update the map after modification
			ordersMap[orderID] = order
		} else {
			asUUID, err := r.WebAppID.UUIDValue()
			if err != nil {
				return nil, errors.Wrap(err, "p.gen.GetNotificationsForNewOrdersAfterCursor")
			}

			ordersMap[orderID] = NewOrderNotification{
				ID:              r.OrderID,
				ReadableOrderID: r.ReadableID.Int64,
				CreatedAt:       r.CreatedAt.Time,
				UserNickname:    r.Username.String,
				WebAppID:        asUUID.Bytes,
				WebAppName:      r.WebAppName,
				Products: []Product{{
					Name:          r.Name,
					Quantity:      int(r.Quantity),
					Price:         r.Price,
					PriceCurrency: r.PriceCurrency,
				}},
			}
		}
	}

	for _, order := range ordersMap {
		newOrderNotifications = append(newOrderNotifications, order)
	}

	return newOrderNotifications, nil
}

// GetNotificationsForNewMarketplacesAfterCursor gets notifcations for marketplaces
// which were created after date specified in cursor
func (p *Pg) GetNotificationsForNewMarketplacesAfterCursor(ctx context.Context, cur Cursor) ([]NewMarketplaceNotification, error) {
	var newMarketplaceNotifications []NewMarketplaceNotification

	rows, err := p.gen.GetNotificationsForNewMarketplacesAfterCursor(
		ctx,
		generated.GetNotificationsForNewMarketplacesAfterCursorParams{
			CreatedAt: pgtype.Timestamp{
				Time:  cur.CursorDate,
				Valid: true,
			},
			Limit: int32(p.newMarketplaceFetchLimit),
		})
	if err != nil {
		return nil, errors.Wrap(err, "p.gen.GetNotificationsForNewMarketplacesAfterCursor")
	}

	for _, marketplace := range rows {
		newMarketplaceNotifications = append(newMarketplaceNotifications, NewMarketplaceNotification{
			ID:   marketplace.ID,
			Name: marketplace.Name,
		})
	}

	return newMarketplaceNotifications, nil
}

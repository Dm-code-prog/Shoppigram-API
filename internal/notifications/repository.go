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
	gen             *generated.Queries
	encryptionKey   string
	orderFetchLimit int
}

// NewPg creates a new Pg
func NewPg(db *pgxpool.Pool, encryptionKey string, orderFetchLimit int) *Pg {
	return &Pg{
		gen:             generated.New(db),
		encryptionKey:   encryptionKey,
		orderFetchLimit: orderFetchLimit,
	}
}

// GetAdminsNotificationList gets a list of admins to notificate about an order
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

// GetAdminBotToken gets admin bot token
func (p *Pg) GetAdminBotToken(ctx context.Context, webAppID uuid.UUID) (string, error) {
	token, err := p.gen.GetAdminBotToken(
		ctx,
		generated.GetAdminBotTokenParams{ID: webAppID, EncryptionKey: p.encryptionKey},
	)
	if err != nil {
		return "", errors.Wrap(err, "p.gen.GetAdminBotToken")
	}

	if token == nil {
		return "", errors.New("the admin token is nil")
	}

	return token.(string), err
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
		LastProcessedCreatedAt: cursor.LastProcessedCreatedAt.Time,
		LastProcessedID:        cursor.LastProcessedID.Bytes,
	}, nil
}

// UpdateNotifierCursor updates notifier cursor
func (p *Pg) UpdateNotifierCursor(ctx context.Context, cur Cursor) error {
	err := p.gen.UpdateNotifierCursor(ctx, generated.UpdateNotifierCursorParams{
		Name: pgtype.Text{
			String: cur.Name,
			Valid:  true,
		},
		LastProcessedCreatedAt: pgtype.Timestamp{
			Time:  cur.LastProcessedCreatedAt,
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

// GetNotificationsForOrdersAfterCursor gets notifcations for orders which were
// created after date specified in cursor
func (p *Pg) GetNotificationsForOrdersAfterCursor(ctx context.Context, cur Cursor) ([]OrderNotification, error) {
	var orderNotifications []OrderNotification

	rows, err := p.gen.GetNotificationsForOrdersAfterCursor(
		ctx,
		generated.GetNotificationsForOrdersAfterCursorParams{
			CreatedAt: pgtype.Timestamp{
				Time:  cur.LastProcessedCreatedAt,
				Valid: true,
			},
			Limit: int32(p.orderFetchLimit),
		})
	if err != nil {
		return nil, errors.Wrap(err, "p.gen.GetNotificationsForOrdersAfterCursor")
	}

	ordersMap := map[string]OrderNotification{}

	for _, r := range rows {
		orderID := r.OrderID.String()
		if order, ok := ordersMap[orderID]; ok {
			// If the order exists, append the new product to the existing order's product list
			order.Products = append(order.Products, Product{
				Name:     r.Name,
				Quantity: int(r.Quantity),
				Price:    r.Price,
			})
			// Update the map after modification
			ordersMap[orderID] = order
		} else {
			asUUID, err := r.WebAppID.UUIDValue()
			if err != nil {
				return nil, errors.Wrap(err, "p.gen.GetNotificationsForOrdersAfterCursor")
			}

			ordersMap[orderID] = OrderNotification{
				ID:              r.OrderID,
				ReadableOrderID: r.ReadableID.Int64,
				CreatedAt:       r.CreatedAt.Time,
				UserNickname:    r.Username.String,
				WebAppID:        asUUID.Bytes, // fix this
				Products: []Product{{
					Name:     r.Name,
					Quantity: int(r.Quantity),
					Price:    r.Price,
				}},
			}
		}
	}

	for _, order := range ordersMap {
		orderNotifications = append(orderNotifications, order)
	}

	return orderNotifications, nil
}

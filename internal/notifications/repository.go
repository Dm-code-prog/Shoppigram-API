package notifications

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/shoppigram-com/marketplace-api/internal/notifications/generated"
)

// Pg implements the Repository interface
// using PostgreSQL as the backing store.
type Pg struct {
	gen                           *generated.Queries
	orderFetchLimit               int
	newMarketplaceFetchLimit      int
	verifiedMarketplaceFetchLimit int
}

// NewPg creates a new Pg
func NewPg(db *pgxpool.Pool, newOrderFetchLimit int, newMarketplaceFetchLimit int, verifiedMarketplaceFetchLimit int) *Pg {
	return &Pg{
		gen:                           generated.New(db),
		orderFetchLimit:               newOrderFetchLimit,
		newMarketplaceFetchLimit:      newMarketplaceFetchLimit,
		verifiedMarketplaceFetchLimit: verifiedMarketplaceFetchLimit,
	}
}

// GetAdminsNotificationList gets a list of admins to notify about an order
func (p *Pg) GetAdminsNotificationList(ctx context.Context, webAppID uuid.UUID) ([]adminNotitfication, error) {
	adminsNotificationList, err := p.gen.GetAdminsNotificationList(ctx, pgtype.UUID{
		Bytes: webAppID,
		Valid: true,
	})
	if err != nil {
		return nil, errors.Wrap(err, "p.gen.GetAdminsNotificationList")
	}

	var retNotifs []adminNotitfication
	for _, v := range adminsNotificationList {
		retNotifs = append(retNotifs, adminNotitfication{
			Id:       v.AdminChatID,
			Language: v.LanguageCode.String,
		})
	}

	return retNotifs, nil
}

// GetReviewersNotificationList gets a list of reviewers to notify about a new marketplace
func (p *Pg) GetReviewersNotificationList(ctx context.Context) ([]int64, error) {
	reviewersNotificationList, err := p.gen.GetReviewersNotificationList(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "p.gen.GetReviewersNotificationList")
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

// GetNotificationsForOrders gets notifications for orders.
func (p *Pg) GetNotificationsForOrders(ctx context.Context, cursor Cursor) ([]OrderNotification, error) {
	rows, err := p.gen.GetNotificationsForUpdatedOrders(ctx, generated.GetNotificationsForUpdatedOrdersParams{
		Limit: int32(p.orderFetchLimit),
		UpdatedAt: pgtype.Timestamp{
			Time:  cursor.CursorDate,
			Valid: true,
		},
		ID: cursor.LastProcessedID,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "p.gen.GetNotificationsForUpdatedOrders")
	}

	var notifications []OrderNotification
	for _, r := range rows {
		var products []Product
		err := json.Unmarshal(r.Products, &products)
		if err != nil {
			return nil, errors.Wrap(err, "json.Unmarshal")
		}

		notifications = append(notifications, OrderNotification{
			ID:              r.OrderID,
			ReadableOrderID: r.ReadableID.Int64,
			CreatedAt:       r.CreatedAt.Time,
			BuyerNickname:   r.BuyerUsername.String,
			BuyerLanguage:   r.BuyerLanguageCode.String,
			OwnerLanguage:   r.AdminLanguageCode.String,
			WebAppID:        r.WebAppID.Bytes,
			WebAppName:      r.WebAppName,
			WebAppCurrency:  string(r.Currency),
			Products:        products,
			Status:          r.State,
			Comment:         "",
			PaymentType:     r.PaymentType,
			BuyerExternalID: int64(r.BuyerExternalUserID),
		})
	}

	return notifications, nil
}

// GetNotificationsForNewOrdersAfterCursor gets notifcations for orders which were
// created after date specified in cursor
func (p *Pg) GetNotificationsForNewOrdersAfterCursor(ctx context.Context, cur Cursor) ([]NewOrderNotification, error) {
	var newOrderNotifications []NewOrderNotification

	rows, err := p.gen.GetNotificationsForNewOrdersAfterCursor(
		ctx,
		generated.GetNotificationsForNewOrdersAfterCursorParams{
			UpdatedAt: pgtype.Timestamp{
				Time:  cur.CursorDate,
				Valid: true,
			},
			ID:    cur.LastProcessedID,
			Limit: int32(p.orderFetchLimit),
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
				Name:     r.Name,
				Quantity: int(r.Quantity),
				Price:    r.Price,
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
				UserLanguage:    r.LanguageCode.String,
				OwnerLanguage:   r.AdminLanguageCode.String,
				WebAppID:        asUUID.Bytes,
				WebAppName:      r.WebAppName,
				Status:          r.State,
				PaymentType:     r.PaymentType,
				Products: []Product{{
					Name:     r.Name,
					Quantity: int(r.Quantity),
					Price:    r.Price,
				}},
				ExternalUserID: int64(r.ExternalUserID),
				WebAppCurrency: string(r.Currency),
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
			ID:    cur.LastProcessedID,
			Limit: int32(p.newMarketplaceFetchLimit),
		})
	if err != nil {
		return nil, errors.Wrap(err, "p.gen.GetNotificationsForNewMarketplacesAfterCursor")
	}

	for _, r := range rows {
		newMarketplaceNotifications = append(newMarketplaceNotifications, NewMarketplaceNotification{
			ID:              r.ID,
			Name:            r.Name,
			ShortName:       r.ShortName,
			CreatedAt:       r.CreatedAt.Time,
			OwnerUsername:   r.Username.String,
			OwnerLanguage:   r.LanguageCode.String,
			OwnerExternalID: int64(r.OwnerExternalID),
		})
	}

	return newMarketplaceNotifications, nil
}

// GetNotificationsForVerifiedMarketplacesAfterCursor gets notifcations for marketplaces
// which were verified after date specified in cursor
func (p *Pg) GetNotificationsForVerifiedMarketplacesAfterCursor(ctx context.Context, cur Cursor) ([]VerifiedMarketplaceNotification, error) {
	var verifiedMarketplaceNotifications []VerifiedMarketplaceNotification

	rows, err := p.gen.GetNotificationsForVerifiedMarketplacesAfterCursor(
		ctx,
		generated.GetNotificationsForVerifiedMarketplacesAfterCursorParams{
			VerifiedAt: pgtype.Timestamp{
				Time:  cur.CursorDate,
				Valid: true,
			},
			ID:    cur.LastProcessedID,
			Limit: int32(p.verifiedMarketplaceFetchLimit),
		})
	if err != nil {
		return nil, errors.Wrap(err, "p.gen.GetNotificationsForNewMarketplacesAfterCursor")
	}

	for _, r := range rows {
		verifiedMarketplaceNotifications = append(verifiedMarketplaceNotifications, VerifiedMarketplaceNotification{
			ID:                  r.ID,
			Name:                r.Name,
			ShortName:           r.ShortName,
			VerifiedAt:          r.VerifiedAt.Time,
			OwnerExternalUserID: int64(r.OwnerExternalID.Int32),
			OwnerLanguage:       r.LanguageCode.String,
		})
	}

	return verifiedMarketplaceNotifications, nil
}

// AddUserToNewOrderNotifications creates a new order notification
// list entry for some marketplace
func (p *Pg) AddUserToNewOrderNotifications(ctx context.Context, req AddUserToNewOrderNotificationsRequest) error {
	err := p.gen.AddUserToNewOrderNotifications(ctx, generated.AddUserToNewOrderNotificationsParams{
		WebAppID: pgtype.UUID{
			Bytes: req.WebAppID,
			Valid: true,
		},
		AdminChatID: req.AdminChatID,
	})
	if err != nil {
		return errors.Wrap(err, "p.gen.AddUserToNewOrderNotifications")
	}

	return nil
}

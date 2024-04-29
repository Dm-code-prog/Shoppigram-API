package adminbot

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/shoppigram-com/marketplace-api/internal/adminbot/generated"
)

// Pg implements the Repository interface
// using PostgreSQL as the backing store.
type Pg struct {
	gen           *generated.Queries
	encryptionKey string
}

// NewPg creates a new Pg
func NewPg(db *pgxpool.Pool, encryptionKey string) *Pg {
	return &Pg{gen: generated.New(db), encryptionKey: encryptionKey}
}

// GetAdminsNotificationList gets a list of admins to notificate about an order
func (p *Pg) GetAdminsNotificationList(ctx context.Context, webAppID uuid.UUID) ([]string, error) {
	var adminsList []string

	vals, err := p.gen.GetAdminsNotificationList(ctx, pgtype.UUID{
		Bytes: webAppID,
		Valid: true,
	})
	if err != nil {
		return nil, errors.Wrap(err, "p.gen.GetAdminsNotificationList")
	}

	for _, v := range vals {
		adminsList = append(adminsList, v.String)
	}

	return adminsList, nil
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

	return token.(string), nil
}

package telegram_users

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pkg/errors"
	"github.com/shoppigram-com/marketplace-api/internal/users/generated"
)

// Pg implements the Repository interface
// using PostgreSQL as the backing store.
type Pg struct {
	gen *generated.Queries
}

// NewPg creates a new Pg
func NewPg(gen *generated.Queries) *Pg {
	return &Pg{gen: gen}
}

// CreateOrUpdateTelegramUser creates or updates a user record
func (p *Pg) CreateOrUpdateTelegramUser(ctx context.Context, request CreateOrUpdateTelegramUserRequest) (uuid.UUID, error) {
	id, err := p.gen.CreateOrUpdateTelegramUser(ctx, generated.CreateOrUpdateTelegramUserParams{
		ExternalID: int32(request.User.ExternalId),
		IsBot: pgtype.Bool{
			Bool:  request.User.IsBot,
			Valid: true,
		},
		FirstName: request.User.FirstName,
		LastName: pgtype.Text{
			String: request.User.LastName,
			Valid:  true,
		},
		Username: pgtype.Text{
			String: request.User.Username,
			Valid:  true,
		},
		LanguageCode: pgtype.Text{
			String: request.User.LanguageCode,
			Valid:  true,
		},
		IsPremium: pgtype.Bool{
			Bool:  request.User.IsPremium,
			Valid: true,
		},
		AllowsPm: pgtype.Bool{
			Bool:  request.User.AllowsPm,
			Valid: true,
		},
	})
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "p.gen.CreateOrUpdateTelegramUser")
	}

	return id, nil
}

func (p *Pg) GetEndUserBotToken(ctx context.Context, request CreateOrUpdateTelegramUserRequest) (string, error) {
	// TODO: Add token getting logic
	return "", nil
}

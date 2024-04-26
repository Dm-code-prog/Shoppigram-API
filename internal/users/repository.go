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

// CreateOrUpdateTgUser creates or updates a user record
func (p *Pg) CreateOrUpdateTgUser(ctx context.Context, request CreateOrUpdateTgUserRequest) (uuid.UUID, error) {
	id, err := p.gen.CreateOrUpdateTgUser(ctx, generated.CreateOrUpdateTgUserParams{
		ExternalID: int32(request.ExternalId),
		IsBot: pgtype.Bool{
			Bool:  request.IsBot,
			Valid: true,
		},
		FirstName: request.FirstName,
		LastName: pgtype.Text{
			String: request.LastName,
			Valid:  true,
		},
		Username: pgtype.Text{
			String: request.Username,
			Valid:  true,
		},
		LanguageCode: pgtype.Text{
			String: request.LanguageCode,
			Valid:  true,
		},
		IsPremium: pgtype.Bool{
			Bool:  request.IsPremium,
			Valid: true,
		},
		AllowsPm: pgtype.Bool{
			Bool:  request.AllowsPm,
			Valid: true,
		},
	})
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "p.gen.CreateOrUpdateTgUser")
	}

	return id, nil
}

func (p *Pg) GetEndUserBotToken(ctx context.Context, webAppID uuid.UUID) (string, error) {
	token, err := p.gen.GetEndUserBotToken(ctx, generated.GetEndUserBotTokenParams{ID: webAppID})
	if err != nil {
		return "", errors.Wrap(err, "p.gen.GetEndUserBotToken")
	}

	return token.(string), nil
}

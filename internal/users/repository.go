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
		return uuid.Nil, errors.Wrap(err, "p.gen.CreateOrUpdateTgUser")
	}

	return id, nil
}

func (p *Pg) GetEndUserBotToken(ctx context.Context, request CreateOrUpdateTgUserRequest) (string, error) {
	// TODO: Add token getting logic
	return "", nil
}

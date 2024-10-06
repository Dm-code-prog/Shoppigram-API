package telegram_users

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"

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
func NewPg(db *pgxpool.Pool) *Pg {
	return &Pg{gen: generated.New(db)}
}

// CreateOrUpdateTgUser creates or updates a user record
func (p *Pg) CreateOrUpdateTgUser(ctx context.Context, request CreateOrUpdateTgUserRequest) (uuid.UUID, error) {
	id, err := p.gen.CreateOrUpdateTgUser(ctx, generated.CreateOrUpdateTgUserParams{
		ExternalID: request.ExternalId,
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

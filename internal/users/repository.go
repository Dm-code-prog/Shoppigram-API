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

// TelegramAuthUser creates or updates a user record
func (p *Pg) TelegramAuthUser(ctx context.Context, request TelegramAuthUserRequest) (uuid.UUID, error) {
	tgAuthUserParams := generated.TelegramAuthUserParams{
		ExternalID: int32(request.User.ExternalId),
		IsBot: pgtype.Bool{
			Bool:  request.User.IsBot,
			Valid: true,
		},
		FirstName: request.User.FirstName,
		IsPremium: pgtype.Bool{
			Bool:  request.User.IsPremium,
			Valid: true,
		},
		AllowsPm: pgtype.Bool{
			Bool:  request.User.AllowsPm,
			Valid: true,
		},
	}

	switch request.User.LastName {
	case "":
		tgAuthUserParams.LastName = pgtype.Text{
			String: "",
			Valid:  false,
		}
	default:
		tgAuthUserParams.LastName = pgtype.Text{
			String: request.User.LastName,
			Valid:  true,
		}
	}

	switch request.User.Username {
	case "":
		tgAuthUserParams.Username = pgtype.Text{
			String: "",
			Valid:  false,
		}
	default:
		tgAuthUserParams.Username = pgtype.Text{
			String: request.User.Username,
			Valid:  true,
		}
	}

	switch request.User.LanguageCode {
	case "":
		tgAuthUserParams.LanguageCode = pgtype.Text{
			String: "",
			Valid:  false,
		}
	default:
		tgAuthUserParams.LanguageCode = pgtype.Text{
			String: request.User.LanguageCode,
			Valid:  true,
		}
	}

	id, err := p.gen.TelegramAuthUser(ctx, tgAuthUserParams)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "p.gen.TelegramAuthUser")
	}

	return id, nil
}

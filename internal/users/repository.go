package telegram_users

import (
	"context"

	"github.com/jackc/pgx/v5"
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

// CreateUser creates a user record
func (p *Pg) CreateUser(ctx context.Context, request AuthUserRequest) error {
	err := p.gen.CreateUser(ctx, generated.CreateUserParams{
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
		return errors.Wrap(err, "p.gen.CreateUser")
	}
	return nil
}

// GetUser returns a user record
func (p *Pg) GetUser(ctx context.Context, request AuthUserRequest) (User, error) {
	usr, err := p.gen.GetUser(ctx, int32(request.User.ExternalId))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return User{}, ErrorNotFound
		}
		return User{}, errors.Wrap(err, "p.gen.GetUser")
	}

	return User{
		ID:           usr.ID,
		ExternalId:   int(usr.ExternalID),
		IsBot:        usr.IsBot.Bool,
		FirstName:    usr.FirstName,
		LastName:     usr.LastName.String,
		Username:     usr.Username.String,
		LanguageCode: usr.LanguageCode.String,
		IsPremium:    usr.IsPremium.Bool,
		AllowsPm:     usr.AllowsPm.Bool,
	}, nil
}

// UpdateUser updates a user record
func (p *Pg) UpdateUser(ctx context.Context, request AuthUserRequest) error {
	err := p.gen.UpdateUser(ctx, generated.UpdateUserParams{
		ExternalID: int32(request.User.ExternalId),
		FirstName:  request.User.FirstName,
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
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrorNotFound
		}
		return errors.Wrap(err, "p.gen.UpdateUser")
	}

	return nil
}

// DeleteUser deletes a user record
func (p *Pg) DeleteUser(ctx context.Context, request AuthUserRequest) error {
	err := p.gen.DeleteUser(ctx, int32(request.User.ExternalId))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrorNotFound
		}
		return errors.Wrap(err, "p.gen.DeleteUser")
	}

	return nil
}

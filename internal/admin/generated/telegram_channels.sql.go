// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: telegram_channels.sql

package generated

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const getTelegramChannels = `-- name: GetTelegramChannels :many
select id,
       name,
       title,
       external_id
from telegram_channels
where owner_external_id = $1
`

type GetTelegramChannelsRow struct {
	ID         uuid.UUID
	Name       pgtype.Text
	Title      string
	ExternalID int64
}

func (q *Queries) GetTelegramChannels(ctx context.Context, ownerExternalID int64) ([]GetTelegramChannelsRow, error) {
	rows, err := q.db.Query(ctx, getTelegramChannels, ownerExternalID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTelegramChannelsRow
	for rows.Next() {
		var i GetTelegramChannelsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Title,
			&i.ExternalID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const isUserTheOwnerOfTelegramChannel = `-- name: IsUserTheOwnerOfTelegramChannel :one
select owner_external_id = $1
from telegram_channels
where external_id = $2
`

type IsUserTheOwnerOfTelegramChannelParams struct {
	OwnerExternalID int64
	ExternalID      int64
}

func (q *Queries) IsUserTheOwnerOfTelegramChannel(ctx context.Context, arg IsUserTheOwnerOfTelegramChannelParams) (bool, error) {
	row := q.db.QueryRow(ctx, isUserTheOwnerOfTelegramChannel, arg.OwnerExternalID, arg.ExternalID)
	var column_1 bool
	err := row.Scan(&column_1)
	return column_1, err
}

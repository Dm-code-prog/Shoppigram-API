-- name: CreateOrUpdateTelegramChannel :exec
insert into telegram_channels (external_id, title, name, is_public, owner_external_id)
values ($1, $2, $3, $4, $5)
on conflict (external_id) do update
    set title             = $2,
        name              = $3,
        is_public         = $4,
        owner_external_id = $5;

-- name: GetTelegramChannelOwner :one
select owner_external_id
from telegram_channels
where external_id = $1;

-- name: DeleteTelegramChannel :exec
delete
from telegram_channels
where external_id = $1;
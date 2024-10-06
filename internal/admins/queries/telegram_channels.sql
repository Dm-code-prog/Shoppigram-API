-- name: IsUserTheOwnerOfTelegramChannel :one
select owner_external_id = $1
from telegram_channels
where external_id = $2;

-- name: GetTelegramChannels :many
select id,
       name,
       title,
       external_id
from telegram_channels
where owner_external_id = $1;

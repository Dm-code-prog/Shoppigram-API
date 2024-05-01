-- name: GetAdminsNotificationList :many
select admin_username, admin_chat_id
from notify_list
where web_app_id = $1;

-- name: GetAdminBotToken :one
select pgp_sym_decrypt(admin_bot_encr_token, @encryption_key::text)
from web_apps
where id = $1;

-- name: GetNotifierCursor :one
select last_processed_created_at, last_processed_id
from notifier_cursors
where name = $1;

-- name: UpdateNotifierCursor :exec
update notifier_cursors
set last_processed_created_at = $2,
    last_processed_id = $3
where name = $1;

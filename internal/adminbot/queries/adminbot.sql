-- name: GetAdminsNotificationList :many
select admin_username
from notify_list
where web_app_id = $1;

-- name: GetAdminBotToken :one
select pgp_sym_decrypt(admin_bot_encr_token, @encryption_key::text)
from web_apps
where id = $1;

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
    last_processed_id         = $3
where name = $1;

-- name: GetNotificationsForOrdersAfterCursor :many
with orders_batch as (select id as order_id, created_at, readable_id, web_app_id, external_user_id
                      from orders o
                      where o.created_at > $1
                      order by o.created_at
                      limit $2)
select orders_batch.order_id,
       orders_batch.readable_id,
       orders_batch.created_at,
       p.web_app_id,
       p.name,
       p.price,
       op.quantity,
       p.price_currency,
       u.username
from orders_batch
         join order_products op
              on orders_batch.order_id = op.order_id
         join products p on p.id = op.product_id
         join telegram_users u on external_user_id = u.external_id;


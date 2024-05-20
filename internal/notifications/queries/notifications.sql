-- name: GetNotifierCursor :one
select cursor_date, last_processed_id
from notifier_cursors
where name = $1;

-- name: UpdateNotifierCursor :exec
update notifier_cursors
set cursor_date = $2,
    last_processed_id = $3
where name = $1;

-- name: GetAdminsNotificationList :many
select admin_username, admin_chat_id
from new_order_notifications_list
where web_app_id = $1;

-- name: GetReviewersNotificationList :many
select chat_id
from new_web_apps_notifications_list;

-- name: GetNotificationsForNewOrdersAfterCursor :many
with orders_batch as (select id as order_id, created_at, readable_id, web_app_id, external_user_id
                      from orders o
                      where o.created_at > $1
                      order by o.created_at
                      limit $2)
select orders_batch.order_id,
       orders_batch.readable_id,
       orders_batch.created_at,
       p.web_app_id,
       wa.name as web_app_name,
       p.name,
       p.price,
       p.price_currency,
       op.quantity,
       u.username
from orders_batch
         join order_products op
              on orders_batch.order_id = op.order_id
         join products p on p.id = op.product_id
         join telegram_users u on external_user_id = u.external_id
         join web_apps wa on orders_batch.web_app_id = wa.id
order by orders_batch.created_at, orders_batch.order_id;

-- name: GetNotificationsForNewMarketplacesAfterCursor :many
with marketplaces_batch as (select wa.id,
                                   wa.name,
                                   wa.short_name,
                                   wa.created_at,
                                   wa.owner_external_id
         from web_apps wa
         where wa.is_verified = false
         and wa.created_at > $1
         order by wa.created_at
         limit $2)
select marketplaces_batch.id,
       marketplaces_batch.name,
       marketplaces_batch.short_name,
       marketplaces_batch.created_at,
       u.username
from marketplaces_batch
         join telegram_users u
              on marketplaces_batch.owner_external_id = u.external_id
order by marketplaces_batch.created_at, marketplaces_batch.id;

-- name: GetNotificationsForVerifiedMarketplacesAfterCursor :many
with marketplaces_batch as (select wa.id,
                                   wa.name,
                                   wa.short_name,
                                   wa.verified_at,
                                   wa.owner_external_id
         from web_apps wa
         where wa.is_verified = true
         and wa.verified_at > $1
         order by wa.verified_at
         limit $2)
select marketplaces_batch.id,
       marketplaces_batch.name,
       marketplaces_batch.short_name,
       marketplaces_batch.verified_at,
       u.username
from marketplaces_batch
         join telegram_users u
              on marketplaces_batch.owner_external_id = u.external_id
order by marketplaces_batch.verified_at, marketplaces_batch.id;

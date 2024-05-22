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
select ob.order_id,
       ob.readable_id,
       ob.created_at,
       p.web_app_id,
       wa.name as web_app_name,
       p.name,
       p.price,
       p.price_currency,
       op.quantity,
       u.username
from orders_batch ob
         join order_products op
              on ob.order_id = op.order_id
         join products p on p.id = op.product_id
         join telegram_users u on external_user_id = u.external_id
         join web_apps wa on ob.web_app_id = wa.id
order by ob.created_at, ob.order_id;

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
select mb.id,
       mb.name,
       mb.short_name,
       mb.created_at,
       u.username
from marketplaces_batch mb
         join telegram_users u
              on mb.owner_external_id = u.external_id
order by mb.created_at, mb.id;

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
select mb.id,
       mb.name,
       mb.short_name,
       mb.verified_at,
       mb.owner_external_id
from marketplaces_batch mb
order by mb.verified_at, mb.id;

-- name: GetNotifierCursor :one
select cursor_date, last_processed_id
from notifier_cursors
where name = $1;

-- name: UpdateNotifierCursor :exec
update notifier_cursors
set cursor_date       = $2,
    last_processed_id = $3
where name = $1;

-- name: GetAdminsNotificationList :many
select admin_chat_id
from new_order_notifications_list
where web_app_id = $1;

-- name: GetReviewersNotificationList :many
select chat_id
from new_web_apps_notifications_list;

-- name: GetNotificationsForNewOrdersAfterCursor :many
with orders_batch as (select id as order_id, created_at, readable_id, web_app_id, external_user_id, state, type
                      from orders o
                      where (o.updated_at, o.id) > (@updated_at::timestamp, @id::uuid)
                        and o.state = 'confirmed'
                      order by o.created_at, o.id
                      limit $1)
select ob.order_id,
       ob.readable_id,
       ob.created_at,
	   ob.state::text,
       p.web_app_id,
       wa.name       as web_app_name,
       p.name,
       p.price,
       p.price_currency,
       op.quantity,
       u.username,
	   	 u.language_code,
       u.external_id as external_user_id,
       adm.language_code as admin_language_code,
       ob.state::text as state,
	   ob.type::text as payment_type
from orders_batch ob
         join order_products op
              on ob.order_id = op.order_id
         join products p on p.id = op.product_id
         join telegram_users u on external_user_id = u.external_id
         join web_apps wa on ob.web_app_id = wa.id
         join telegram_users adm on wa.owner_external_id = adm.external_id
where ob.state = 'confirmed'
order by ob.created_at, ob.order_id;

-- name: GetNotificationsForNewMarketplacesAfterCursor :many
with marketplaces_batch as (select wa.id,
                                   wa.name,
                                   wa.short_name,
                                   wa.created_at,
                                   wa.owner_external_id
                            from web_apps wa
                            where wa.is_verified = false
                              and (wa.created_at, wa.id) > (@created_at::timestamp, @id::uuid)
                            order by wa.created_at, wa.id
                            limit $1)
select mb.id,
       mb.name,
       mb.short_name,
       mb.created_at,
       u.username,
	   u.language_code,
       u.external_id as owner_external_id
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
                              and (wa.verified_at, wa.id) > (@verified_at::timestamp, @id::uuid)
                            order by wa.verified_at, wa.id
                            limit $1)
select mb.id,
       mb.name,
       mb.short_name,
       mb.verified_at,
       mb.owner_external_id,
	   u.language_code
from marketplaces_batch mb
	 join telegram_users u on mb.owner_external_id = u.external_id
order by mb.verified_at, mb.id;

-- name: AddUserToNewOrderNotifications :exec
insert into new_order_notifications_list (web_app_id, admin_chat_id)
values ($1,
        $2);

-- name :: GetNotificationsForOrderStatusChangeAfterCursor :many
-- with orders_batch as (select
-- 	 			  	  from orders o
--					  where )

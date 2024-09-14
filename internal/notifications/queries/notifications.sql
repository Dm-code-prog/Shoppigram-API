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
with admins_batch as (select admin_chat_id
                      from new_order_notifications_list
                      where web_app_id = $1)
select ab.admin_chat_id,
       u.language_code
from admins_batch ab
         join telegram_users u on ab.admin_chat_id = u.external_id;

-- name: GetReviewersNotificationList :many
select chat_id
from new_web_apps_notifications_list;

-- name: GetProductCustomMessage :one
select message
from products_custom_messages
where product_id = $1
  and on_order_state = $2
order by created_at desc
limit 1;

-- name: GetNotificationsForUpdatedOrders :many
with orders_batch as (select id as order_id,
                             created_at,
                             readable_id,
                             external_user_id,
                             state,
                             type,
                             web_app_id
                      from orders o
                      where (o.updated_at, o.id) > (@updated_at::timestamp, @id::uuid)
                      order by o.updated_at, o.id
                      limit $1)
select orders_batch.order_id    as order_id,
       orders_batch.readable_id as readable_id,
       orders_batch.created_at  as created_at,
       orders_batch.state::text as state,
       orders_batch.web_app_id  as web_app_id,
       wa.name                  as web_app_name,
       coalesce(
               json_agg(json_build_object(
                       'id', p.id,
                       'name', p.name,
                       'quantity', op.quantity,
                       'price', p.price
                        )
               ),
               '[]'::json
       ) ::json                 as products,
       wa.currency              as currency,
       u.username               as buyer_username,
       u.language_code          as buyer_language_code,
       u.external_id            as buyer_external_user_id,
       adm.language_code        as admin_language_code,
       orders_batch.state::text as state,
       orders_batch.type::text  as payment_type
from orders_batch
         join order_products op
              on orders_batch.order_id = op.order_id
         join products p on p.id = op.product_id
         join telegram_users u on external_user_id = u.external_id
         join web_apps wa on orders_batch.web_app_id = wa.id
         join telegram_users adm on wa.owner_external_id = adm.external_id
group by orders_batch.order_id, orders_batch.readable_id, orders_batch.created_at, orders_batch.state::text,
         orders_batch.web_app_id, wa.name,
         wa.currency, op.quantity, u.username, u.language_code, u.external_id, adm.language_code,
         orders_batch.state::text,
         orders_batch.type::text;


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

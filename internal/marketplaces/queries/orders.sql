-- name: CreateP2POrder :one
insert into orders (web_app_id, external_user_id, type, state)
values ($1,
        $2,
        'p2p',
        'confirmed'::order_state)
returning id,readable_id;

-- name: CreateOnlineOrder :one
insert into orders (web_app_id, external_user_id, type, state)
values ($1,
        $2,
        'online',
        'created'::order_state)
returning id,readable_id;

-- name: SetOrderProducts :batchexec
insert into order_products (order_id, product_id, quantity)
values ($1,
        $2,
        $3);

-- name: GetOrder :many
select p.id,
       p.name,
       op.quantity,
       p.description,
       p.category,
       p.price,
       wa.currency::text as price_currency,
       wa.name           as web_app_name,
       wa.short_name     as web_app_short_name,
       o.readable_id,
       tu.username       as seller_username
from orders o
         join order_products op on o.id = op.order_id
         join products p on op.product_id = p.id
         join web_apps wa on o.web_app_id = wa.id
         join telegram_users tu on wa.owner_external_id = tu.external_id
where o.id = $1
  and (o.external_user_id = $2
    or wa.owner_external_id = $2);


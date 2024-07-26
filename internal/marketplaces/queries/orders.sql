-- name: CreateP2POrder :one
insert into orders (web_app_id, external_user_id, type, state)
values ($1,
        $2,
        'p2p',
        'approved')
returning id,readable_id;

-- name: SetOrderProducts :batchexec
insert into order_products (order_id, product_id, quantity)
values ($1,
        $2,
        $3);
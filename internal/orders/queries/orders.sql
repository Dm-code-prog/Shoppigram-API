-- name: CreateOrder :one
insert into orders (web_app_id, telegram_user_id)
values ($1,
        $2)
returning id;

-- name: SetOrderProducts :batchexec
insert into order_products (order_id, product_id, quantity)
values ($1,
        $2,
        $3);
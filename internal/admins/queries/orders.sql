-- name: GetOrders :many
SELECT o.id                       AS id,
       o.web_app_id               AS marketplace_id,
       o.readable_id              AS readable_id,
       (SELECT SUM(p.price * op.quantity)
        FROM order_products op
                 JOIN products p ON p.id = op.product_id
        WHERE op.order_id = o.id) AS total_price,
       tu.username                AS buyer_username,
       (SELECT json_agg(
                       json_build_object(
                               'id', p.id,
                               'name', p.name,
                               'quantity', op.quantity,
                               'price', p.price,
                               'price_currency', p.price_currency
                       )
               )
        FROM order_products op
                 JOIN products p ON p.id = op.product_id
        WHERE op.order_id = o.id) AS products
FROM orders o
         JOIN
     telegram_users tu ON tu.external_id = o.external_user_id
         join web_apps wa on wa.id = o.web_app_id
where tu.external_id = @owner_external_id::integer
  and wa.owner_external_id = @owner_external_id::integer
  and (o.web_app_id = @marketplace_id::uuid or @marketplace_id is null)
  and (o.state = @state or @state is null)
limit $1 offset $2;

-- name: GetOrders :many
SELECT o.id                       AS id,
       o.web_app_id               AS marketplace_id,
       o.readable_id              AS readable_id,
       o.state                    AS state,
       o.type                     AS type,
       o.created_at               AS created_at,
       o.updated_at               AS updated_at,
       wa.currency                AS currency,
       (SELECT SUM(p.price * op.quantity)::float8
        FROM order_products op
                 JOIN products p ON p.id = op.product_id
        WHERE op.order_id = o.id) AS total_price,
       tu.username                AS buyer_username,
       (SELECT json_agg(
                       json_build_object(
                               'id', p.id,
                               'name', p.name,
                               'quantity', op.quantity,
                               'price', p.price
                       )
               )
        FROM order_products op
                 JOIN products p ON p.id = op.product_id
        WHERE op.order_id = o.id) AS products
FROM orders o
         JOIN
     telegram_users tu ON tu.external_id = o.external_user_id
         join web_apps wa on wa.id = o.web_app_id
where wa.owner_external_id = @owner_external_id
  and (
    case when @state != '' then state = @state::order_state else true end
    )
  and (
    case
        when @marketplace_id != '00000000-0000-0000-0000-000000000000' then web_app_id = @marketplace_id::uuid
        else true end
    )
order by o.created_at desc limit $1
offset $2;

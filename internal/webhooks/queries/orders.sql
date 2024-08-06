-- name: GetOrder :one
select o.id,
       o.updated_at,
       sum(op.quantity::float * p.price)::float as order_sum,
       MAX(p.price_currency)::text              as price_currency,
       o.state
from orders o
         join order_products op on op.order_id = o.id
         join products p on p.id = op.product_id
where o.id = $1
group by o.id;

-- name: SetOrderStateConfirmed :exec
update orders
set state = 'confirmed'::order_state
where id = $1
  and type = 'online'::order_type;


-- name: SavePaymentExtraInfo :exec
insert into payments_extra_info (invoice_id, provider, event_type, extra_info, response, error)
values (@invoice_id,
        @provider,
        @event_type,
        @extra_info,
        @response,
        @error);

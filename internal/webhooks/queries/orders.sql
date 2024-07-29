-- name: GetOrder :one
select o.id, o.updated_at, sum(op.quantity::float * p.price)::float as order_sum, MAX(p.price_currency)::text as price_currency
from orders o
	 join order_products op on op.order_id = o.id
	 join products p on p.id = op.product_id
where o.id = $1
group by o.id;

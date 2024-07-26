-- name: GetOrder :one
select o.id, o.updated_at, SUM(p.price*op.quantity)
from orders o
	 join order_products op on op.order_id = o.id
	 join products p on p.id = op.product_id
where o.id = $1
group by o.id;


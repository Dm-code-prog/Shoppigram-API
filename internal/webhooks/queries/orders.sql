-- name: GetOrder :one
select o.updated_at, o.invoice_id, SUM(p.price*op.quantity)
from orders o
	 join order_products op on op.order_id = o.id
	 join products p on p.id = op.product_id
where o.invoice_id = $1
group by o.id;


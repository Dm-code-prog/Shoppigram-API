-- name: GetBalance :many
select sum((p.price - (p.price * wa.commission_percent) - wa.commission_fixed) * op.quantity) as balance, wa.currency
from web_apps wa
         left join orders o on wa.id = o.web_app_id
         left join order_products op on o.id = op.order_id
         left join products p on op.product_id = p.id
where wa.owner_external_id = $1
  and o.type = 'online'
  and o.state = 'done'
group by currency;
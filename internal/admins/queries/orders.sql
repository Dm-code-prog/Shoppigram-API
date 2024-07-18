-- name: GetOrder :many
select p.id,
       p.name,
       p.description,
       p.category,
       p.price,
       p.price_currency::text as price_currency,
       wa.name        as web_app_name,
       wa.short_name  as web_app_short_name
from orders o
	 join order_products op on o.id = op.order_id
	 join products p on op.product_id = p.id
     join web_apps wa on o.web_app_id = wa.id
where o.id = $1
	  and (o.external_user_id = $2
	  or wa.owner_external_id = $2);

-- name: GetProducts :many
select w.name as web_app_name,
       p.id,
       p.name,
       p.description,
       p.price,
       p.price_currency,
       p.image_url
from products p
         join web_apps w on p.web_app_id = w.id
where w.id = $1;
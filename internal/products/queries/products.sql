-- name: GetProducts :many
select w.name as web_app_name,
       p.id,
       p.name,
       p.description,
       p.category,
       p.price,
       p.price_currency,
       p.image_url
from web_apps w
         join products p on w.id = p.web_app_id
where w.id = $1;

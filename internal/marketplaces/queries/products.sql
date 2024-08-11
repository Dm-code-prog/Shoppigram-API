-- name: GetProducts :many
select p.id,
       p.name,
       p.description,
       p.category,
       p.price,
       p.price_currency::text     as price_currency,
       wa.id                      as web_app_id,
       wa.name                    as web_app_name,
       wa.short_name              as web_app_short_name,
       wa.is_verified             as web_app_is_verified,
       wa.online_payments_enabled as online_payments_enabled
from web_apps wa
         join products p on wa.id = p.web_app_id
where wa.id = $1
  and wa.is_deleted = false;


-- name: GetMarketplaceWithProducts :one
select wa.id,
       wa.name,
       wa.short_name,
       wa.is_verified,
       wa.online_payments_enabled,
       wa.currency,
       json_agg(
               json_build_object(
                       'id', p.id,
                       'name', p.name,
                       'description', p.description,
                       'category', p.category,
                       'price', p.price
               )
       ) as products
from web_apps wa
         left join products p on wa.id = p.web_app_id
where wa.id = $1
  and wa.is_deleted = false
group by wa.id, wa.name, wa.short_name, wa.is_verified, wa.online_payments_enabled, wa.currency;


-- name: GetMarketplaceWithProducts :one
SELECT wa.id,
       wa.name,
       wa.short_name,
       wa.is_verified,
       wa.online_payments_enabled,
       wa.currency,
       wa.type,
       COALESCE(
                       json_agg(
                       json_build_object(
                               'id', p.id,
                               'name', p.name,
                               'description', p.description,
                               'category', p.category,
                               'price', p.price,
                               'price_currency', p.price_currency::text
                       )
                               ) FILTER (WHERE p.id IS NOT NULL),
                       '[]'::json
       )::json AS products
FROM web_apps wa
         LEFT JOIN products p ON wa.id = p.web_app_id and p.is_deleted = false
WHERE (
    case when @web_app_id::uuid != '00000000-0000-0000-0000-000000000000' then wa.id = @web_app_id::uuid else true end
    )
  and (
    case when @short_name::text != '' then wa.short_name = @short_name::text else true end
    )
  AND wa.is_deleted = false
GROUP BY wa.id, wa.name, wa.short_name, wa.is_verified, wa.online_payments_enabled, wa.currency;


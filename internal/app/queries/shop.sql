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
                               'external_links', COALESCE(
                                       (SELECT json_agg(json_build_object('url', pel.url, 'label', pel.label))
                                        FROM product_external_links pel
                                        WHERE pel.product_id = p.id),
                                       '[]'::json),
                               'photos', coalesce(
                                       (select json_agg(json_build_object('url', pp.url))
                                        from product_photos pp
                                        where pp.product_id = p.id),
                                       '[]'::json
                                         ),
                               'variants', coalesce(
                                       (select json_agg(json_build_object('id', pv.id, 'price', pv.price,
                                                                          'discounted_price', pv.discounted_price,
                                                                          'dimensions', pv.dimensions))
                                        from product_variants pv
                                        where pv.product_id = p.id),
                                       '[]'::json
                                           ))
                               ) FILTER (WHERE p.id IS NOT NULL),
                       '[]'::json
       )::json AS products
FROM web_apps wa
         LEFT JOIN products p ON wa.id = p.web_app_id AND p.is_deleted = false
WHERE (
    CASE WHEN @web_app_id::uuid != '00000000-0000-0000-0000-000000000000' THEN wa.id = @web_app_id::uuid ELSE TRUE END
    )
  AND (
    CASE WHEN @short_name::text != '' THEN wa.short_name = @short_name::text ELSE TRUE END
    )
  AND wa.is_deleted = false
GROUP BY wa.id, wa.name, wa.short_name, wa.is_verified, wa.online_payments_enabled, wa.currency, wa.type;


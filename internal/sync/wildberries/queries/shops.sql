-- name: GetNextShop :one
select wa.id, wa.created_at, sec.api_key
from web_apps wa
         left join shop_external_connections sec
                   on wa.id = sec.web_app_id
                       and sec.external_provider = 'wildberries'::external_provider
                       and sec.is_active
where wa.created_at > @cursor_date
  and wa.id > @cursor_id
order by wa.created_at, wa.id
limit 1;
-- name: GetMarketplaces :many
select id, name, logo_url
from web_apps
where owner_external_id = $1;

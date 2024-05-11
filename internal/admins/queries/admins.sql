-- name: GetMarketplaces :many
select id, name, logo_url, is_verified
from web_apps
where owner_external_id = $1;

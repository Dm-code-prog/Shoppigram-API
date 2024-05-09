-- name: GetMarketplacesByUserID :many
select id, name
from web_apps
where owner_external_id = $1;

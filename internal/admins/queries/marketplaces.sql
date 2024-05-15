-- name: GetMarketplaces :many
select id, name, logo_url, is_verified
from web_apps
where owner_external_id = $1;

-- name: CreateMarketplace :one
insert into web_apps (name, owner_external_id, short_name)
values ($1,
        $2,
        $3)
returning id;

-- name: CountUserMarketplaces :one
select count(*)
from web_apps
where owner_external_id = $1;

-- name: UpdateMarketplace :execresult
update web_apps
set name = $1
where id = $2
  and owner_external_id = $3;


-- name: IsUserTheOwnerOfMarketplace :one
select owner_external_id = $1
from web_apps
where id = $2;


-- name: GetMarketplaceShortName :one
select short_name
from web_apps
where id = @id::uuid;
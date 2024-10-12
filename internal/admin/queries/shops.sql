-- name: GetShops :many
select id, name, logo_url, is_verified, short_name, currency, type
from web_apps
where owner_external_id = $1
  and is_deleted = false;

-- name: CreateShop :one
insert into web_apps (name, owner_external_id, short_name, currency, type)
values ($1,
        $2,
        $3,
        $4,
        $5)
returning id;

-- name: SoftDeleteShop :exec
update web_apps
set is_deleted= true
where id = $1;

-- name: CountUserShops :one
select count(*)
from web_apps
where owner_external_id = $1
  and is_deleted = false;

-- name: UpdateShop :execresult
update web_apps
set name = $1
where id = $2
  and owner_external_id = $3;


-- name: IsShopOwner :one
select owner_external_id = $1
from web_apps
where id = $2;


-- name: GetShortname :one
select short_name
from web_apps
where id = @id::uuid;

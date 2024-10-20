-- name: GetShops :many
select id, name, is_verified, short_name, currency, type
from web_apps
where is_deleted = false
  and owner_external_id = $1;

-- name: GetShop :one
select wa.id,
       wa.name,
       wa.is_verified,
       wa.short_name,
       wa.currency,
       wa.type,
       sec.external_provider,
       sec.is_active,
       sec.last_sync_at,
       sec.last_sync_status
from web_apps wa
         left join shop_external_connections sec on wa.id = sec.web_app_id
where wa.id = $1
  and wa.owner_external_id = $2
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

-- name: UpdateShop :execresult
update web_apps
set name = $1
where id = $2
  and owner_external_id = $3
  and is_deleted = false
returning id;

-- name: IsShopSyncSupported :one
select exists(select 1
              from web_apps
              where id = $1
                and type = 'panel'
                and is_deleted = false
                and is_verified = true
                and currency = 'rub'::product_currency);

-- name: EnableShopSync :exec
insert
into shop_external_connections
    (web_app_id, external_provider, api_key, is_active)
values ($1, $2, $3, $4)
on conflict (web_app_id, external_provider)
    do update set api_key   = $3,
                  is_active = $4;


-- name: CountUserShops :one
select count(*)
from web_apps
where owner_external_id = $1
  and is_deleted = false;


-- name: IsShopOwner :one
select owner_external_id = $1
from web_apps
where id = $2;


-- name: GetShortname :one
select short_name
from web_apps
where id = @id::uuid;

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

-- name: CreateProduct :one
insert into products (web_app_id, name, description, price, price_currency, category)
values (@web_app_id::uuid,
        $1,
        nullif(@description::text, ''),
        $2,
        $3,
        nullif(@category::varchar(30), ''))
returning id;

-- name: CountMarketplaceProducts :one
select count(*)
from products
where web_app_id = @web_app_id::uuid;

-- name: UpdateProduct :execresult
update products
set name           = $1,
    description    = nullif(@description::text, ''),
    price          = $2,
    price_currency = $3,
    category       = nullif(@category::varchar(30), '')
where web_app_id = $4
  and id = $5;

-- name: DeleteProduct :execresult
delete
from products
where web_app_id = $1
  and id = $2;

-- name: IsUserTheOwnerOfWebApp :one
select owner_external_id = $1
from web_apps;

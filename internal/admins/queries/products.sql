-- name: CreateProduct :one
insert into products (web_app_id, name, description, price, price_currency, category, image_url, extra_properties)
values (@web_app_id::uuid,
        $1,
        nullif(@description::text, ''),
        $2,
        $3,
        nullif(@category::varchar(30), ''),
        '',
		@extra_properties)
returning id;

-- name: CountMarketplaceProducts :one
select count(*)
from products
where web_app_id = @web_app_id::uuid;

-- name: UpdateProduct :execresult
update products
set name             = $1,
    description      = nullif(@description::text, ''),
    price            = $2,
    price_currency   = $3,
    category         = nullif(@category::varchar(30), ''),
	extra_properties = $4
where web_app_id = @web_app_id::uuid
  and id = $5;

-- name: DeleteProduct :execresult
delete
from products
where web_app_id = @web_app_id::uuid
  and id = @id::uuid;

-- name: IsUserTheOwnerOfWebApp :one
select owner_external_id = $1
from web_apps
where id = $2;

-- name: IsUserTheOwnerOfProduct :one
select wa.owner_external_id = $1 from products p
join web_apps wa on p.web_app_id = wa.id
where p.id = @id::uuid;

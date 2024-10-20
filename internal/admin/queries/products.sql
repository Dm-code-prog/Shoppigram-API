-- name: CreateProduct :one
insert into products (web_app_id, name, description, price, category)
values (@web_app_id::uuid,
        $1,
        nullif(@description::text, ''),
        $2,
        nullif(@category::varchar(30), ''))
returning id;

-- name: SetProductExternalLinks :batchexec
insert into product_external_links (product_id, url, label)
values ($1,
        $2,
        $3);

-- name: RemoveProductExternalLinks :exec
delete
from product_external_links
where product_id = $1;

-- name: CountMarketplaceProducts :one
select count(*)
from products
where web_app_id = @web_app_id::uuid;

-- name: UpdateProduct :execresult
update products
set name        = $1,
    description = nullif(@description::text, ''),
    price       = $2,
    category    = nullif(@category::varchar(30), '')
where web_app_id = @web_app_id::uuid
  and id = $3;

-- name: DeleteProduct :execresult
update products
set is_deleted = true
where web_app_id = @web_app_id::uuid
  and id = $1;

-- name: IsUserTheOwnerOfWebApp :one
select owner_external_id = $1
from web_apps
where id = $2;

-- name: IsUserTheOwnerOfProduct :one
select wa.owner_external_id = $1
from products p
         join web_apps wa on p.web_app_id = wa.id
where p.id = @id::uuid;

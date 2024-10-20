-- name: GetExternalIds :many
select external_id
from products
where web_app_id = @web_app_id
  and external_provider = @external_provider
  and is_deleted = false;

-- name: MarkProductAsDeleted :batchexec
update products
set is_deleted = true
where external_id in ($1);

-- name: CreateOrUpdateProducts :batchexec
insert into products (web_app_id, name, description, price, category, external_provider, external_id)
values (@web_app_id,
        @name,
        @description,
        @price,
        @category,
        @external_provider,
        @external_id)
on conflict (external_id, external_provider) do update
    set name        = excluded.name,
        description = excluded.description,
        price       = excluded.price,
        category    = excluded.category;

-- name: GetProducts :many
select p.id,
       p.external_id,
       p.name,
       p.description,
       p.price,
       p.category
from products p
where p.web_app_id = @web_app_id
  and p.external_provider = @external_provider
  and p.is_deleted = false;

-- name: CreateOrUpdateExternalLinks :batchexec
insert into product_external_links (product_id, url, label)
values (@product_id, @url, @label);

-- name: DeleteExternalLinks :batchexec
delete
from product_external_links
where product_id = @product_id;

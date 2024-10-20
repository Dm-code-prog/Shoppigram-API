-- name: DeletePhotos :batchexec
delete
from product_photos
where id in ($1);

-- name: CreateOrUpdatePhotos :batchexec
insert into product_photos(url, product_id)
values (@url, @product_id)
on conflict (url, product_id) do update
    set url = excluded.url;
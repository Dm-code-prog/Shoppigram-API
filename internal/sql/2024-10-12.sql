create type web_app_type as enum ('shop', 'panel');

alter table web_apps
    add column type web_app_type not null default 'shop';


create table product_external_links
(
    id         uuid               default uuid_generate_v4() primary key,
    product_id uuid      not null references products (id),
    url        text      not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);
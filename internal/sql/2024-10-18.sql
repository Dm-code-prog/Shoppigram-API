create type external_provider as enum ('wildberries');

alter table products
    add column external_provider external_provider;

alter table products
    add column
        external_id text;

alter table products
    add column created_at timestamp default now();

alter table products
    add column updated_at timestamp default now();

alter table products
    drop column if exists extra_properties;

alter table products
    drop column if exists price_currency;

alter table products
    drop column if exists image_url;

create unique index on products (external_id, external_provider);

create or replace function set_updated_at()
    returns trigger as
$$
begin
    new.updated_at = now();
    return new;
end;
$$ language plpgsql;

create trigger set_updated_at
    before update
    on products
    for each row
execute function set_updated_at();

create type extenal_sync_status as enum ('success', 'failure');

create table shop_external_connections
(
    id                uuid              not null default uuid_generate_v4(),
    created_at        timestamp                  default now(),
    updated_at        timestamp                  default now(),
    web_app_id        uuid              not null references web_apps (id),
    is_active         boolean           not null default true,
    external_provider external_provider not null,
    api_key           text              not null,
    last_sync_at      timestamp,
    last_failure_at   timestamp,
    last_sync_status  extenal_sync_status,
    last_error        text,
    primary key (id)
);

create trigger set_updated_at
    before update
    on shop_external_connections
    for each row
execute function set_updated_at();

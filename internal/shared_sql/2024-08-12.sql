alter table products
    add column if not exists extra_properties JSON default '{}'::json not null;

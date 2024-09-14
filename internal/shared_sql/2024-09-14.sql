alter table products
    add column is_deleted boolean not null default false;
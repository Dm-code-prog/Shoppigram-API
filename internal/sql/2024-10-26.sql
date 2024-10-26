create table product_variants
(
    id               uuid           not null primary key default uuid_generate_v4(),
    created_at       timestamp                           default now(),
    updated_at       timestamp                           default now(),
    product_id       uuid           not null references products (id),
    dimensions       jsonb          not null,
    price            numeric(10, 2) not null,
    discounted_price numeric(10, 2),
    is_deleted       boolean        not null             default false
);

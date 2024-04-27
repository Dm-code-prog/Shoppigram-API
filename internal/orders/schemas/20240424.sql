create table orders
(
    id               uuid primary key default uuid_generate_v4(),
    -- This might cause errors in very rare cases, in this case the app might just try again.
    readable_id      bigint unique    default floor(random() * (1000000000000 - 1 + 1)) + 1,
    web_app_id       uuid references web_apps (id),
    external_user_id int references telegram_users (external_id),
    created_at       timestamp        default now(),
    updated_at       timestamp        default now()
);

create table order_products
(
    order_id   uuid references orders (id),
    product_id uuid references products (id),
    quantity   int not null check ( quantity > 0 )
);

create table orders
(
    id               uuid primary key default uuid_generate_v4(),
    web_app_id       uuid references web_apps (id),
    telegram_user_id uuid references telegram_users (id)
);

create table order_products
(
    order_id   uuid references orders (id),
    product_id uuid references products (id),
    quantity   int not null
);
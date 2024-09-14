create table if not exists products_custom_messages
(
    id             uuid default uuid_generate_v4() not null primary key,
    created_at     timestamp default now() not null,
    product_id     uuid                            not null references products (id),
    message        text                            not null,
    on_order_state order_state                     not null
);
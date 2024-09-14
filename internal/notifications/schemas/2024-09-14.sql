create table if not exists products_custom_messages
(
    id             uuid      default uuid_generate_v4() not null primary key,
    created_at     timestamp default now()              not null,
    product_id     uuid                                 not null references products (id),
    message        text                                 not null,
    on_order_state order_state                          not null
);


create table if not exists product_custom_media_forwards
(
    id             uuid      default uuid_generate_v4() not null primary key,
    created_at     timestamp default now()              not null,
    product_id     uuid                                 not null references products (id),
    from_chat_id   bigint                               not null,
    message_id     bigint                               not null,
    on_order_state order_state                          not null
);
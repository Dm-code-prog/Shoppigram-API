create extension if not exists "uuid-ossp";

create table telegram_users
(
    id            uuid primary key default uuid_generate_v4(),
    external_id   int unique   not null,
    is_bot        bool,
    first_name    varchar(255) not null,
    last_name     varchar(255),
    username      varchar(35),
    language_code varchar(3),
    is_premium    bool,
    allows_pm     bool,
    created_at    timestamp        default now(),
    updated_at    timestamp        default now()
);

create table web_apps
(
    id                uuid      default uuid_generate_v4() not null
        primary key,
    name              varchar(50)                          not null,
    owner_external_id integer
        references telegram_users (external_id),
    logo_url          text,
    is_verified       boolean   default false,
    short_name        varchar(30)                          not null
        unique
        constraint web_apps_short_name_check
            check ((short_name)::text ~ '^[a-z]{5,}$'::text),
    verified_at       timestamp,
    created_at        timestamp default now(),
    is_deleted        boolean   default false
);

create index
    on web_apps (created_at, id);

create index
    on web_apps (verified_at, id);


create type product_currency as enum ('rub', 'usd', 'eur');

create table products
(
    id             uuid             default uuid_generate_v4()      not null
        primary key,
    web_app_id     uuid
        references web_apps,
    name           varchar(255)                                      not null,
    description    text,
    price          double precision                                 not null,
    price_currency product_currency default 'rub'::product_currency not null,
    image_url      text,
    category       varchar(255)
);

create table orders
(
    id               uuid      default uuid_generate_v4() not null
        primary key,
    readable_id      bigint    default (floor((random() * ((('1000000000000'::bigint - 1) + 1))::double precision)) +
                                        (1)::double precision)
        unique,
    web_app_id       uuid
        references web_apps,
    external_user_id integer
        references telegram_users (external_id),
    created_at       timestamp default now(),
    updated_at       timestamp default now()
);

create index
    on orders (created_at, id);


create table order_products
(
    order_id   uuid
        references orders,
    product_id uuid
        references products,
    quantity   integer not null
        constraint order_products_quantity_check
            check (quantity > 0)
);
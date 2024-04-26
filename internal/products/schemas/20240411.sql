-- Web apps
create table web_apps
(
    id   uuid primary key default uuid_generate_v4(),
    name varchar(50) not null
);

-- Products shown in the Web app
create table products
(
    id             uuid primary key     default uuid_generate_v4(),
    web_app_id     uuid references web_apps (id),
    name           varchar(50) not null,
    description    text,
    price          float       not null,
    price_currency varchar(10) not null default 'rub',
    image_url      text        not null
);

-- End user bots
create table end_user_bots
(
    id                  uuid primary key default uuid_generate_v4(),
    telegram_name       text  not null,
    encrypted_api_token bytea not null,
    web_app_id          uuid references web_apps (id)
);

-- Admin bots
create table admin_bots
(
    id                  uuid primary key default uuid_generate_v4(),
    telegram_name       text  not null,
    encrypted_api_token bytea not null
);

-- Telegram channels
create table channels
(
    id              uuid primary key default uuid_generate_v4(),
    telegram_name   text not null,
    end_user_bot_id uuid references end_user_bots (id),
    admin_bot_id    uuid references admin_bots (id)
);

create table telegram_channels
(
    id                uuid primary key default uuid_generate_v4(),
    external_id       bigint  not null unique,
    title             text    not null,
    name              text,
    is_public         boolean not null,
    owner_external_id bigint  not null
);

create index on telegram_channels (owner_external_id);
create index on telegram_channels (external_id);

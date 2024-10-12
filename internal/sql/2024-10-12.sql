create type web_app_type as enum ('shop', 'panel');

alter table web_apps
    add column type web_app_type not null default 'shop';
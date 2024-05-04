-- Notifiers
create table notifier_cursors
(
    last_processed_created_at timestamp,
    last_processed_id         uuid references orders (id),
    name                      varchar(50) unique
);

-- Notify list
create table notify_list
(
    web_app_id     uuid references web_apps (id),
    admin_username varchar(35),
    admin_chat_id  bigint not null unique
);


alter table notify_list
    drop constraint notify_list_admin_chat_id_key;

alter table notify_list
    add constraint notify_list_chat_id_web_app_key unique (web_app_id, admin_chat_id);
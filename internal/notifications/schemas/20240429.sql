create table new_order_notifications_list
(
    web_app_id     uuid
        constraint notify_list_web_app_id_fkey
            references web_apps,
    admin_username varchar(35),
    admin_chat_id  bigint not null,
    constraint notify_list_chat_id_web_app_key
        unique (web_app_id, admin_chat_id)
);


create table new_web_apps_notifications_list
(
    chat_id bigint not null
);

create table notifier_cursors
(
    cursor_date       timestamp,
    last_processed_id uuid,
    name              varchar(50)
        unique
);
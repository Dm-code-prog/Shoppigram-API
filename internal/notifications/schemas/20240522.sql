create index new_orders_notifications_index on orders (created_at, id);
create index new_web_apps_notifications_index on web_apps (created_at, id);
create index verified_web_apps_notifications_index on web_apps (verified_at, id);

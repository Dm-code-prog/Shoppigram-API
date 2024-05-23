create index new_orders_notifications_index on orders (created_at, id) using btree;
create index new_web_apps_notifications_index on web_apps (created_at, id) using btree;
create index verified_web_apps_notifications_index on web_apps (verified_at, id) using btree;

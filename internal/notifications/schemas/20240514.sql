alter table notify_list rename to new_order_notifications_list;

-- New Web Apps Notifications
create table new_web_apps_notifications_list (
	web_app_id uuid references web_apps(id) not null,
	chat_id bigint not null
);

alter table notifier_cursors rename column last_processed_created_at to cursor_date;

alter table web_apps add column verified_at timestamp;

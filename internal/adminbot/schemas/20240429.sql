-- Notifiers
create table notifier_cursors (
  last_processed_created_at timestamp primary key,
  last_processed_id uuid references orders(id),
  name varchar(50) unique
);

-- Notify list
create table notify_list (
  web_app_id uuid references web_apps(id),
  admin_username varchar(35)
);

alter table web_apps add column owner_external_id int references telegram_users (external_id);
alter table web_apps add column image_url text;

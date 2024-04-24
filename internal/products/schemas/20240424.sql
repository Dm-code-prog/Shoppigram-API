alter table web_apps add column end_user_bot_name text;
alter table web_apps add column end_user_bot_encr_token bytea;

alter table web_apps add column admin_bot_name text;
alter table web_apps add column admin_bot_encr_token bytea;
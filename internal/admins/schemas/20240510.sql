alter table web_apps drop column image_url;
alter table web_apps add column logo_url text;
alter table web_apps add column is_verified boolean default false;

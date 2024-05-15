alter table web_apps
    add column short_name varchar(30) unique check ( short_name ~ '^[a-z]{5,}$');

alter table web_apps
    alter column short_name set not null;
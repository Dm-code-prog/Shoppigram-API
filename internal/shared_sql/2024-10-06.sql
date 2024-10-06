alter table telegram_users
    alter column external_id type bigint;

alter table web_apps
    alter column owner_external_id type bigint;

alter table orders
    alter column external_user_id type bigint;


select *
from telegram_users;

select *
from web_apps
where owner_external_id = 7237487742::integer;
alter table web_apps
    add column currency product_currency default 'rub'::product_currency not null;

alter table web_apps
    add column commission_percent numeric(5, 2) default 0.0 not null;

alter table web_apps
    add column commission_fixed numeric(10, 2) default 0.0 not null;

alter table products
    alter column price_currency type text;
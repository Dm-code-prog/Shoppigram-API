create type product_currency as enum ('rub', 'usd', 'eur');

ALTER TABLE products
    ALTER COLUMN price_currency DROP DEFAULT;

ALTER TABLE products
    ALTER COLUMN price_currency TYPE product_currency
        USING price_currency::product_currency;

ALTER TABLE products
    ALTER COLUMN price_currency SET DEFAULT 'rub';

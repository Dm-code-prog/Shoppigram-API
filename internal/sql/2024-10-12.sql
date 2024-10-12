create type web_app_type as enum ('shop', 'panel');

alter table web_apps
    add column type web_app_type not null default 'shop';


create table product_external_links
(
    id         uuid               default uuid_generate_v4() primary key,
    product_id uuid      not null references products (id),
    url        text      not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);


alter table web_apps
    drop constraint web_apps_short_name_key;

create unique index on web_apps (short_name)
    where is_deleted = false;


-- Drop existing trigger and function if they exist
DROP TRIGGER IF EXISTS on_web_app_is_verified_update ON web_apps;
DROP FUNCTION IF EXISTS update_verified_at();

-- Create the trigger function
CREATE FUNCTION update_verified_at() RETURNS trigger
    LANGUAGE plpgsql
AS
$$
BEGIN
    -- Update 'verified_at' when 'is_verified' changes from not true to true
    IF NEW.is_verified = TRUE AND COALESCE(OLD.is_verified, FALSE) IS DISTINCT FROM TRUE THEN
        NEW.verified_at := NOW();
    END IF;
    RETURN NEW;
END;
$$;

-- Create the trigger
CREATE TRIGGER on_web_app_is_verified_update
    BEFORE UPDATE
    ON web_apps
    FOR EACH ROW
EXECUTE PROCEDURE update_verified_at();
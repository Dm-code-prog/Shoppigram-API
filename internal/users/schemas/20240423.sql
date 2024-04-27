-- Users
create table telegram_users (
	id uuid primary key default uuid_generate_v4(),
	external_id int unique not null,
	is_bot bool,
	first_name varchar(50) not null,
	last_name varchar(50),
	username varchar(35),
	language_code varchar(3),
	is_premium bool,
	allows_pm bool,
	created_at timestamp default now(),
	updated_at timestamp default now()
);

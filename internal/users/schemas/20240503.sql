alter table telegram_users add column temp_first_name VARCHAR(64);
alter table telegram_users add column temp_last_name VARCHAR(64);
alter table telegram_users add column temp_username VARCHAR(32);

update telegram_users set temp_first_name = first_name;
update telegram_users set temp_last_name = last_name;
update telegram_users set temp_username = username;

alter table telegram_users drop column first_name;
alter table telegram_users drop column last_name;
alter table telegram_users drop column username;

alter table telegram_users rename column temp_first_name to first_name;
alter table telegram_users rename column temp_last_name to last_name;
alter table telegram_users rename column temp_username to username;

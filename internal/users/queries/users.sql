-- name: CreateUser :exec
insert into telegram_users (id,
                   external_id, -- 1
                   is_bot, -- 2
                   first_name, -- 3
                   last_name, -- 4
                   username, -- 5
                   language_code, -- 6
                   is_premium, -- 7
                   allows_pm) -- 8
values (uuid_generate_v4(),
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8);

-- name: UpdateUser :exec
update telegram_users
set first_name = $2,
    last_name = $3,
    username = $4,
    language_code = $5,
    is_premium = $6,
    allows_pm = $7,
    updated_at = now()
where external_id = $1;
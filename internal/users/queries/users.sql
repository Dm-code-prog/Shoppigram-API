-- name: AuthUser :one
insert into telegram_users (external_id,
                            is_bot,
                            first_name,
                            last_name,
                            username,
                            language_code,
                            is_premium,
                            allows_pm)
values ($1, $2, $3, $4, $5, $6, $7, $8)
on conflict (external_id)
do update set first_name = $3,
    last_name = $4,
    username = $5,
    language_code = $6,
    is_premium = $7,
    allows_pm = $8,
    updated_at = now()
returning id;
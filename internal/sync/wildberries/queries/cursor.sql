-- name: GetCursor :one
select cursor_id, cursor_timestamp
from cursors
where name = $1;

-- name: SetCursor :exec
update cursors
set cursor_id        = $2,
    cursor_timestamp = $3
where name = $1;

-- name: ResetCursor :exec
update cursors
set cursor_id        = null,
    cursor_timestamp = null
where name = $1;
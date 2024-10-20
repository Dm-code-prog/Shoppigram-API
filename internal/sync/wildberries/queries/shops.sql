-- name: GetNextShop :one
select id, web_app_id, api_key
from shop_external_connections
where is_active = true
  and external_provider = 'wildberries'::external_provider
  and last_sync_at < now() - @sync_interval
  and last_failure_at < now() - @failure_retry_interval
order by last_sync_at
limit 1;

-- name: SetSyncSuccess :exec
update shop_external_connections
set last_sync_at     = now(),
    last_sync_status = 'success'::extenal_sync_status
where id = $1;

-- name: SetSyncFailure :exec
update shop_external_connections
set last_failure_at  = now(),
    last_sync_status = 'failure'::extenal_sync_status,
    last_error       = $2
where id = $1;
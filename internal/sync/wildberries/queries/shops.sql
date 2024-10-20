-- name: GetNextShop :one
select id, web_app_id, api_key
from shop_external_connections
where is_active = true
  and external_provider = 'wildberries'::external_provider
  -- default value for the frequence of syncs
  and coalesce(last_sync_at, '1970-01-01'::timestamp) < now() - interval '1 hour'
  -- if the last sync was a failure, we wait 3 hours before retrying
  and coalesce(last_failure_at, '1970-01-01'::timestamp) < now() - interval '3 hour'
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
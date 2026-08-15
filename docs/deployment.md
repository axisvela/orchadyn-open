# Deployment

Run the ORCHADYN Planning API service with PostgreSQL 18.

## Prerequisites

- Go 1.25.12+ (for the Core source) or a published `orchadyn-api` binary from
  the [release artifacts](https://github.com/axisrobo/orchadyn-open/releases).
- PostgreSQL 18.

## Database

Apply the schema migrations from the Core repository:

```sh
export DATABASE_URL="postgres://user:password@localhost:5432/orchadyn?sslmode=disable"
orchadyn-migrate
```

The migration tool creates the `orchadyn` schema, its tables, and row-level
security policies.

## Service

Start the API with the required environment:

```sh
export DATABASE_URL="postgres://user:password@localhost:5432/orchadyn?sslmode=disable"
export ORCHADYN_TENANT_ID="production"
# Optional trusted proxy resolution instead of a fixed tenant:
# export ORCHADYN_TRUSTED_PROXY_TOKEN="shared-secret"
export ORCHADYN_ALLOWED_ORIGINS="https://app.example.com"
orchadyn-api
```

Control-plane integrations (capability catalog, authority, operational state)
are optional and enabled together:

```sh
export ORCHADYN_MODUREGIS_RESOLVE_URL="https://moduregis.example.com/v1/resolve"
export ORCHADYN_AEGIVELA_GRANTS_URL="https://aegivela.example.com/v1/grants"
export ORCHADYN_ONTOVELA_STATE_URL="https://ontovela.example.com/v1/state"
export ORCHADYN_CONTROL_PLANE_BEARER_TOKEN="..."
export ORCHADYN_CONTROL_PLANE_CACHE_TTL="30s"
```

Set all three control-plane URLs together or none; the service refuses to start
with a partial configuration. `ORCHADYN_CONTROL_PLANE_CACHE_TTL` controls how
long upstream governed inputs are cached before a refresh.

## Probes

- `GET /healthz` returns 200 when the plan ledger is ready.
- `GET /metrics` returns operational counters and latency in Prometheus text
  format.

## Local compose

The Core repository ships `deploy/compose.yaml` for local development:

```sh
docker compose -f deploy/compose.yaml up
```

## Verification

Verify a deployment with the quickstart request:

```sh
curl --request POST http://localhost:8080/plans:generate \
  --header "Content-Type: application/json" \
  --data @examples/scheduling-plan.json
```

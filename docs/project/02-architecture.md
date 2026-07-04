# Access Control API — Architecture

## Overview

Multi-tenant IAM backend: users, organizations, RBAC, sessions, invitations, audit logs.

## Stack

| Layer | Choice |
|-------|--------|
| HTTP | Chi v5 (`net/http`) |
| Database | PostgreSQL 16 via pgx/v5 |
| Cache / rate limit | Redis 7 via go-redis |
| Auth | JWT access + refresh rotation, Argon2id |
| Migrations | golang-migrate |
| API docs | swaggo/swag |
| Tests | testify + httptest + testcontainers |

## Folder layout

```
cmd/api/main.go
internal/
  config/       — env-based config (Docker + local)
  domain/       — entities, permissions, errors
  auth/         — JWT, Argon2, token hashing
  repository/   — postgres + redis
  service/      — business logic
  handler/      — HTTP handlers
  middleware/   — auth, RBAC, rate limit, request ID
db/migrations/  — SQL up/down
docs/project/   — plan and progress
```

## Middleware chain

```
RequestID → Logger → Recoverer → RateLimit(auth) → AuthJWT → RequireOrgMember → RequirePermission
```

## Auth flow

1. **Signup/Login** — verify Argon2id hash, create session with hashed refresh token
2. **Refresh** — validate refresh token hash, rotate hash, issue new JWT pair; reuse detection revokes all sessions
3. **Logout/Revoke** — mark session revoked in PostgreSQL + Redis TTL key

## RBAC

- **Route guard**: middleware requires permission (e.g. `users.invite`)
- **Resource check**: service verifies org membership and ownership
- **Deny by default**: no permission row = denied

## Role-permission matrix

| Permission | owner | admin | member | viewer |
|------------|-------|-------|--------|--------|
| users.read | yes | yes | yes | yes |
| users.invite | yes | yes | no | no |
| projects.write | yes | yes | yes | no |
| billing.read | yes | yes | no | no |

## Dual setup (Docker / local)

Same binary; connection via env vars:

- `DATABASE_URL` — `localhost:5432` when using Docker Compose published ports, or local Postgres host
- `REDIS_URL` — `localhost:6379` when using Docker Compose, or local Redis

Run infrastructure: `docker compose up -d`

## Security rules

- Store SHA-256 hashes of refresh and invite tokens — never raw values in DB
- Short-lived access JWT (default 15m)
- Rate limit `/auth/*` via Redis
- Audit sensitive actions: login, invite, role change

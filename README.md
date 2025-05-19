# Shoppigram Marketplace API

**Shoppigram** turns Telegram into a fully-featured commerce platform.
This Go service powers two Telegram Mini-Apps:

| Mini-App                 | Purpose                                                                                    |
| ------------------------ | ------------------------------------------------------------------------------------------ |
| **Shop (`type=shop`)**   | The public storefront where customers browse products and place orders.                    |
| **Panel (`type=panel`)** | The private admin panel for merchants to manage inventory, orders, payments and promotion. |

Merchants connect their bot to Shoppigram, pick a short name (`/mybrand`) and instantly get both Mini-Apps deployed under the same bot.

---

## ✨ Product capabilities

* **Product & catalog management** – CRUD endpoints for products, categories and photos .
* **Multi-currency pricing** – enum `product_currency` supports RUB / USD / EUR .
* **External catalog sync** – hourly job pulls stock from Wildberries supplier API (`external_provider = 'wildberries'`) .
* **Order flow** – orders can be `p2p` or `online` and move through `created → confirmed → done / rejected` states .
* **Online payments** – CloudPayments webhooks verify and capture payments for `online` orders .
* **Balance tracking** – earnings per currency are calculated from completed online orders .
* **Telegram channel promotion** – admins can publish shop banners to their channels and pin/unpin them .
* **Multi-language notifications** – ready-made Markdown templates (`internal/notifications/templates/{en,ru}`) for buyers and sellers.
* **Fine-grained access** – authorization middleware validates Telegram Init Data (`internal/auth`) so that every request is tied to a real chat user.

---

## 🏗 High-level architecture

```
┌──────────────────────────────────┐
│ Telegram Bot (HTTPS Webhooks)    │
└──────────────────────────────────┘
              │
              ▼
┌──────────────────────────────────┐
│ marketplace-api (this repo)      │
│  • cmd/api/main.go               │
│  • internal/... (bounded contexts)│
│                                   │
│  ┌──────────────────────────────┐ │
│  │ Notification job runners     │
│  │ Wildberries syncer           │
│  └──────────────────────────────┘ │
└──────────┬───────────┬───────────┘
           │           │
           ▼           ▼
   PostgreSQL      DigitalOcean Spaces
   (primary DB)    (object storage)
           │
           ▼
      CloudWatch
    (logs & metrics)
```

Key technical choices:

| Concern          | Implementation                                                        |
| ---------------- | --------------------------------------------------------------------- |
| Language & style | Go 1.22, **go-kit** endpoint-transport-service layout.                |
| Data layer       | PostgreSQL with **sqlc**-generated type-safe queries.                 |
| Caching          | In-memory \[ristretto] LRU to speed up shop reads.                    |
| Observability    | Structured logs (zap), HTTP metrics middleware, CloudWatch collector. |
| Deployment       | Docker → Amazon ECR → Amazon ECS (dev & prod GitHub Actions) .        |

---

## 📂 Repository layout

| Path                          | Description                                                    |
| ----------------------------- | -------------------------------------------------------------- |
| **cmd/api/**                  | Main binary – wires configuration, routes and background jobs. |
| **internal/app**              | Customer-facing shop logic (products, orders, cache).          |
| **internal/admin**            | Merchant admin panel (shops, sync, promotion, balance).        |
| **internal/auth**             | Telegram user registry and auth middleware.                    |
| **internal/notifications**    | Templated Telegram notifications service.                      |
| **internal/webhooks**         | Telegram & CloudPayments webhook processors.                   |
| **internal/sync/wildberries** | External catalog synchroniser.                                 |
| **sql/**                      | Plain-SQL migrations, ordered by date.                         |
| **aws/**                      | ECS task definitions for dev & production clusters.            |
| **.github/workflows/**        | CI (test & lint) and CD (ECS deploy).                          |
| **packages/**                 | Re-usable plumbing: logging, CORS, CloudWatch, etc.            |

---

## ⚙️ Configuration

All settings are injected through environment variables (parsed in `cmd/api/environment.go`) .

<details>
<summary>Most important variables</summary>

### Core

| Variable                           | Default    | Purpose                                |
| ---------------------------------- | ---------- | -------------------------------------- |
| `LOG_LEVEL`                        | `INFO`     | `DEBUG`, `INFO`, `WARN`, `ERROR`       |
| `HTTP_PORT`                        | `8080`     | Public port of the API container       |
| `POSTGRES_DSN`                     | —          | PostgreSQL DSN including database name |
| `POSTGRES_MIN_CONNS` / `MAX_CONNS` | `1` / `10` | Connection pool limits                 |

### Telegram bot

| Variable                          | Required |
| --------------------------------- | -------- |
| `BOT_TOKEN`, `BOT_ID`, `BOT_NAME` | Yes      |

### AWS / Spaces

| Variable                                                              | Required | Notes                                           |
| --------------------------------------------------------------------- | -------- | ----------------------------------------------- |
| `DIGITALOCEAN_SPACES_ENDPOINT`, `..._BUCKET`, `..._KEY`, `..._SECRET` | Yes      | S3-compatible object storage for product photos |

### Webhooks & payments

| Variable                                        | Purpose                               |
| ----------------------------------------------- | ------------------------------------- |
| `TELEGRAM_WEBHOOK_SECRET`                       | HMAC validation of Telegram callbacks |
| `CLOUDPAYMENTS_LOGIN`, `CLOUDPAYMENTS_PASSWORD` | API credentials                       |
| `CLOUDPAYMENTS_MAX_TRANSACTION_DURATION`        | Expiry window (e.g. `1h`)             |

### Feature flags / jobs

| Variable                                        | Description                         |
| ----------------------------------------------- | ----------------------------------- |
| `NEW_ORDER_NOTIFICATIONS_IS_ENABLED`            | Push order notifications to sellers |
| `NEW_MARKETPLACE_NOTIFICATIONS_IS_ENABLED`      | Notify admins about new shops       |
| `VERIFIED_MARKETPLACE_NOTIFICATIONS_IS_ENABLED` | Notify when a shop becomes verified |
| `WILDBERRIES_SYNC_IS_ENABLED`                   | Run external syncer                 |

</details>

---

## ▶️ Running locally

```bash
# 1. Clone & prepare
git clone https://github.com/your-org/marketplace-api.git
cd marketplace-api
cp .env.sample .env   # create and fill secrets

# 2. Start Postgres (example with Docker)
docker run --name pg -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:16

# 3. Apply SQL migrations
psql "$POSTGRES_DSN" -f sql/2024-04-11-aggregate.sql
psql "$POSTGRES_DSN" -f sql/2024-07-26.sql
# …or use your favourite tool

# 4. Run the API
make run          # wraps `go run`
```

The service now listens on **`http://localhost:8080`** and exposes Open-style JSON endpoints under `/api/v2/*`.

---

## 🐳 Docker

A minimal multi-stage `Dockerfile` builds a scratch binary and ships it in Alpine :

```bash
docker build -t shoppigram/api .
docker run -p 8080:8080 --env-file .env shoppigram/api
```

---

## ☑️ Tests & linting

```bash
go test ./...      # unit tests
make test          # CI-identical workflow in Docker-act
golangci-lint run  # static analysis (config in GitHub Actions)
```

Coverage badges, race detector and strict vet are enforced in CI.

---

## 🚀 Continuous deployment

* **Dev** – every push to `main` builds an image, pushes to Amazon ECR and triggers ECS rolling update of the *dev* cluster .
* **Production** – version tags `v*` follow the same path targeting the *prod* cluster .
* Task-definition blueprints live under `aws/ecs/{dev,production}`.

---

## 🗄 Database schema

All DDL is expressed in timestamped files under `/sql`. Notable objects:

* `web_apps` – shops (+ `type`, `currency`, commission fields) .
* `products`, `order_products` – catalog & line items .
* Enums: `product_currency`, `order_type`, `order_state`, `web_app_type`, `external_provider` .

Use any migration tool (Goose, Atlas, Flyway) to apply them in order.

---

## 💬 API quick reference

| Area         | Base path                                                      | Notes                                                                        |
| ------------ | -------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| **Customer** | `/api/v2/app`                                                  | Browse shop, create orders                                                   |
| **Auth**     | `/api/v2/auth`                                                 | Telegram Init-Data validation                                                |
| **Admin**    | `/api/v2/admin`                                                | Manage shops, products, orders, publish banners, balance, Telegram channels  |
| **Webhooks** | `/api/v1/telegram/webhooks`, `/api/v1/cloud-payments/webhooks` | Bot & payment callbacks                                                      |

All endpoints are JSON; see the service contracts in `internal/*/contract.go`.

---

## 🌐 Internationalisation

Notification templates exist in **English** and **Russian** under
`internal/notifications/templates/{en,ru}` with Markdown bodies ready to be sent via Telegram bots.

---

## 🤝 Contributing

1. Fork & branch off `main`.
2. Write tests for any change.
3. Run `make test sqlc golangci-lint run`.
4. Submit a pull request describing **what** & **why**.

Please avoid adding new external dependencies unless necessary and keep the project Go 1.22-compatible.

---

## 📜 License

Distributed under **Creative Commons Attribution-NonCommercial 4.0**. Commercial use requires a separate licence .

---

> Made with 💚 by the original Shoppigram engineering team – now open-sourced for the community.

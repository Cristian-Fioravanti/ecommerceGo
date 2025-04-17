# LastingDynamics E-commerce Backend
---
## Setup Instructions

```bash
git clone https://github.com/Cristian-Fioravanti/ecommerceGo.git
cd LastingDynamics
docker compose -p ecommercego up -d
go mod tidy
go run main.go
```

##  Project Structure
├── main.go
│   └── Starts the app, sets up router and DB connection.
│
├── controllers/
│   └── API logic (login, users, products, etc.).
│
├── database/
│   └── Connects to the DB, creates schema, runs auto-migrations.
│
├── middleware/
│   └── Auth middleware (JWT protection).
│
├── models/
│   └── DB models like User and Product (used by GORM).
│
├── routes/
│   └── Defines all routes and links them to controllers.
│
├── tokens/
│   └── Generates and validates JWT tokens.
│
├── go.mod / go.sum
│   └── Go dependencies and module config.
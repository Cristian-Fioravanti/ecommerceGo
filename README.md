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
- `main.go` – starts the app  
- `controllers/` – handles the API logic (login, users, products)  
- `database/` – DB connection + creates DB/tables  
- `middleware/` – auth middleware (JWT)  
- `models/` – DB models (User, Product)  
- `routes/` – sets up all the routes  
- `tokens/` – JWT stuff (generate, validate)  
- `go.mod / go.sum` – dependencies

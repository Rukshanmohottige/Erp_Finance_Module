# ERP Finance Module

**IT2305 Continuous Assessment 1 — ERP System (Finance Module)**

> Tech Stack: Vue 3 · Go (Gin) · PostgreSQL · Docker

---

## 🚀 How to Run

### Prerequisites
- [Docker Desktop](https://www.docker.com/products/docker-desktop/) installed

### Steps

```bash
# 1. Clone/extract the project
cd erp_finance_module

# 2. Copy environment config
cp .env.example .env

# 3. Start everything
docker-compose up --build
```

| Service  | URL                        |
|----------|----------------------------|
| Frontend | http://localhost:3000       |
| Backend  | http://localhost:8080       |
| Database | localhost:5432              |

**Default Login:**
- Email: `admin@erp.com`
- Password: `admin123` *(update hash in init.sql for production)*

---

## 📁 Project Structure

```
erp_finance_module/
├── docker-compose.yml
├── .env.example
├── frontend/                  # Vue 3 app
│   ├── Dockerfile
│   ├── nginx.conf
│   ├── src/
│   │   ├── views/
│   │   │   ├── Dashboard/     # Summary + charts
│   │   │   ├── Budget/        # Budget CRUD
│   │   │   ├── Expenses/      # Expense CRUD
│   │   │   ├── Invoices/      # Invoice CRUD
│   │   │   ├── Payroll/       # Payroll CRUD
│   │   │   └── Reports/       # Budget vs Actual
│   │   ├── store/             # Pinia auth store
│   │   ├── router/            # Vue Router
│   │   ├── api/               # Axios instance
│   │   └── components/        # Layout, sidebar
├── backend/                   # Go + Gin
│   ├── Dockerfile
│   ├── go.mod
│   ├── cmd/server/main.go     # Entry point
│   ├── internal/
│   │   ├── config/            # DB + env config
│   │   ├── handlers/          # All route handlers
│   │   ├── models/            # Data structs
│   │   └── middleware/        # JWT auth
│   └── migrations/
│       └── init.sql           # DB schema + sample data
└── docs/
    └── api.md                 # API documentation
```

---

## 🔌 API Endpoints

### Auth
| Method | Endpoint           | Description  |
|--------|--------------------|--------------|
| POST   | /api/v1/auth/login | Login        |

### Budgets
| Method | Endpoint              | Description     |
|--------|-----------------------|-----------------|
| GET    | /api/v1/budgets       | List budgets    |
| POST   | /api/v1/budgets       | Create budget   |
| GET    | /api/v1/budgets/:id   | Get budget      |
| PUT    | /api/v1/budgets/:id   | Update budget   |
| DELETE | /api/v1/budgets/:id   | Delete budget   |

### Expenses
| Method | Endpoint              | Description     |
|--------|-----------------------|-----------------|
| GET    | /api/v1/expenses      | List expenses   |
| POST   | /api/v1/expenses      | Create expense  |
| PUT    | /api/v1/expenses/:id  | Update expense  |
| DELETE | /api/v1/expenses/:id  | Delete expense  |

### Invoices
| Method | Endpoint              | Description     |
|--------|-----------------------|-----------------|
| GET    | /api/v1/invoices      | List invoices   |
| POST   | /api/v1/invoices      | Create invoice  |
| PUT    | /api/v1/invoices/:id  | Update invoice  |
| DELETE | /api/v1/invoices/:id  | Delete invoice  |

### Payroll
| Method | Endpoint              | Description      |
|--------|-----------------------|------------------|
| GET    | /api/v1/payroll       | List payroll     |
| POST   | /api/v1/payroll       | Add entry        |
| PUT    | /api/v1/payroll/:id   | Update entry     |
| DELETE | /api/v1/payroll/:id   | Delete entry     |

### Reports
| Method | Endpoint                            | Description             |
|--------|-------------------------------------|-------------------------|
| GET    | /api/v1/reports/budget-vs-actual    | Budget vs spend         |
| GET    | /api/v1/reports/expense-by-category | Expenses grouped        |

---

## 🔗 Integration with Other Modules

This Finance Module integrates with other ERP modules via REST APIs using shared identifiers:

| Integration            | Shared Key    | Description                                     |
|------------------------|---------------|-------------------------------------------------|
| Finance ↔ HR           | `employee_id` | Payroll pulls employee list from HR module      |
| Finance ↔ Procurement  | `order_id`    | Invoices link to purchase orders                |
| Finance ↔ Inventory    | `budget_id`   | Inventory spend updates budget spent_amount     |

**All APIs use:**
- JSON data format
- Bearer token authentication
- Base URL: `http://localhost:8080/api/v1`

---

## 👥 Team Member Contributions (7 Members)

| Member | Role                   | Tasks                                              |
|--------|------------------------|----------------------------------------------------|
| 1      | Team Lead / Backend    | Go project setup, main.go, routing, Docker         |
| 2      | Backend Developer      | Budget & Expense handlers, DB queries              |
| 3      | Backend Developer      | Invoice & Payroll handlers, JWT middleware         |
| 4      | Frontend Lead          | Vue setup, Layout, Router, Auth store              |
| 5      | Frontend Developer     | Dashboard & Reports views                          |
| 6      | Frontend Developer     | Budget, Expenses, Invoices, Payroll views          |
| 7      | DevOps / DB / Docs     | docker-compose, init.sql, README, Postman/Swagger  |

---

## 🗄️ PostgreSQL Installation (Local Dev)

> **Note:** You do NOT need PostgreSQL installed locally — Docker handles it.
> Only install locally if you want to connect without Docker.

### Windows
1. Download from https://www.postgresql.org/download/windows/
2. Run installer, set password for `postgres` user
3. Add `C:\Program Files\PostgreSQL\15\bin` to PATH

### macOS
```bash
brew install postgresql@15
brew services start postgresql@15
```

### Ubuntu/Linux
```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
sudo systemctl start postgresql
```

---

## 🏆 Bonus Features Implemented
- JWT Authentication & route protection
- Vue + Go alternative stack (bonus marks eligible)
- Docker with volumes and custom networks
- Sample data pre-loaded via init.sql

# ERP Finance Module — API Documentation

Base URL: `http://localhost:8080/api/v1`

All protected routes require header:
```
Authorization: Bearer <token>
```

---

## POST /auth/login
```json
Request:
{ "email": "admin@erp.com", "password": "admin123" }

Response 200:
{ "token": "eyJ...", "user": { "id": 1, "name": "Admin", "role": "admin" } }
```

---

## GET /dashboard/summary
```json
Response:
{
  "total_budget": 1150000,
  "total_spent": 95000,
  "pending_expenses": 2,
  "unpaid_invoices": 2,
  "pending_payroll": 1,
  "budget_utilization": 8.26
}
```

---

## Budgets

### GET /budgets
Returns array of budget objects.

### POST /budgets
```json
{ "department": "Engineering", "fiscal_year": 2025, "allocated_amount": 500000 }
```

### PUT /budgets/:id
Same body as POST.

### DELETE /budgets/:id
Returns `{ "message": "Deleted" }`

---

## Expenses

### GET /expenses
Returns all expenses.

### POST /expenses
```json
{
  "budget_id": 1,
  "description": "Cloud hosting",
  "amount": 5000,
  "category": "Infrastructure",
  "status": "pending",
  "submitted_by": 1
}
```

---

## Invoices

### POST /invoices
```json
{
  "invoice_number": "INV-2025-010",
  "vendor_name": "AWS",
  "amount": 9000,
  "due_date": "2025-07-01",
  "status": "unpaid"
}
```

---

## Payroll

### POST /payroll
```json
{
  "employee_id": "EMP010",
  "employee_name": "Jane Doe",
  "department": "Engineering",
  "basic_salary": 90000,
  "allowances": 5000,
  "deductions": 9000,
  "pay_period": "2025-05",
  "status": "pending"
}
```

---

## Reports

### GET /reports/budget-vs-actual
```json
[{ "department": "Engineering", "allocated": 500000, "spent": 57000 }]
```

### GET /reports/expense-by-category
```json
[{ "category": "Infrastructure", "total": 45000 }]
```

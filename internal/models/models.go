package models

import "time"

type User struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
}

type Budget struct {
	ID              int       `json:"id"`
	Department      string    `json:"department"`
	FiscalYear      int       `json:"fiscal_year"`
	AllocatedAmount float64   `json:"allocated_amount"`
	SpentAmount     float64   `json:"spent_amount"`
	CreatedAt       time.Time `json:"created_at"`
}

type Expense struct {
	ID          int       `json:"id"`
	BudgetID    int       `json:"budget_id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
	Status      string    `json:"status"`
	SubmittedBy int       `json:"submitted_by"`
	CreatedAt   time.Time `json:"created_at"`
}

type Invoice struct {
	ID            int       `json:"id"`
	InvoiceNumber string    `json:"invoice_number"`
	VendorName    string    `json:"vendor_name"`
	Amount        float64   `json:"amount"`
	DueDate       string    `json:"due_date"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}

type Payroll struct {
	ID           int       `json:"id"`
	EmployeeID   string    `json:"employee_id"`
	EmployeeName string    `json:"employee_name"`
	Department   string    `json:"department"`
	BasicSalary  float64   `json:"basic_salary"`
	Allowances   float64   `json:"allowances"`
	Deductions   float64   `json:"deductions"`
	NetSalary    float64   `json:"net_salary"`
	PayPeriod    string    `json:"pay_period"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}

type DashboardSummary struct {
	TotalBudget      float64 `json:"total_budget"`
	TotalSpent       float64 `json:"total_spent"`
	PendingExpenses  int     `json:"pending_expenses"`
	UnpaidInvoices   int     `json:"unpaid_invoices"`
	PendingPayroll   int     `json:"pending_payroll"`
	BudgetUtilization float64 `json:"budget_utilization"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}


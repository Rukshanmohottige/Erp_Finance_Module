package handlers

import (
	"database/sql"
	"erp_finance/internal/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	DB *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{DB: db}
}

// ── Auth ──────────────────────────────────────────────────────────────────────

func (h *Handler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	row := h.DB.QueryRow("SELECT id, name, email, password_hash, role FROM users WHERE email=$1", req.Email)
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.Role); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenStr, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	c.JSON(http.StatusOK, models.LoginResponse{Token: tokenStr, User: user})
}

// ── Dashboard ─────────────────────────────────────────────────────────────────

func (h *Handler) GetDashboardSummary(c *gin.Context) {
	var s models.DashboardSummary
	h.DB.QueryRow(`SELECT COALESCE(SUM(allocated_amount),0), COALESCE(SUM(spent_amount),0) FROM budgets`).
		Scan(&s.TotalBudget, &s.TotalSpent)
	h.DB.QueryRow(`SELECT COUNT(*) FROM expenses WHERE status='pending'`).Scan(&s.PendingExpenses)
	h.DB.QueryRow(`SELECT COUNT(*) FROM invoices WHERE status='unpaid'`).Scan(&s.UnpaidInvoices)
	h.DB.QueryRow(`SELECT COUNT(*) FROM payroll WHERE status='pending'`).Scan(&s.PendingPayroll)
	if s.TotalBudget > 0 {
		s.BudgetUtilization = (s.TotalSpent / s.TotalBudget) * 100
	}
	c.JSON(http.StatusOK, s)
}

// ── Budgets ───────────────────────────────────────────────────────────────────

func (h *Handler) GetBudgets(c *gin.Context) {
	rows, err := h.DB.Query(`SELECT id, department, fiscal_year, allocated_amount, spent_amount, created_at FROM budgets ORDER BY id`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var budgets []models.Budget
	for rows.Next() {
		var b models.Budget
		rows.Scan(&b.ID, &b.Department, &b.FiscalYear, &b.AllocatedAmount, &b.SpentAmount, &b.CreatedAt)
		budgets = append(budgets, b)
	}
	c.JSON(http.StatusOK, budgets)
}

func (h *Handler) GetBudgetByID(c *gin.Context) {
	id := c.Param("id")
	var b models.Budget
	err := h.DB.QueryRow(`SELECT id, department, fiscal_year, allocated_amount, spent_amount, created_at FROM budgets WHERE id=$1`, id).
		Scan(&b.ID, &b.Department, &b.FiscalYear, &b.AllocatedAmount, &b.SpentAmount, &b.CreatedAt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Budget not found"})
		return
	}
	c.JSON(http.StatusOK, b)
}

func (h *Handler) CreateBudget(c *gin.Context) {
	var b models.Budget
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.DB.QueryRow(`INSERT INTO budgets (department, fiscal_year, allocated_amount) VALUES ($1,$2,$3) RETURNING id`,
		b.Department, b.FiscalYear, b.AllocatedAmount).Scan(&b.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, b)
}

func (h *Handler) UpdateBudget(c *gin.Context) {
	id := c.Param("id")
	var b models.Budget
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.DB.Exec(`UPDATE budgets SET department=$1, fiscal_year=$2, allocated_amount=$3 WHERE id=$4`,
		b.Department, b.FiscalYear, b.AllocatedAmount, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Updated"})
}

func (h *Handler) DeleteBudget(c *gin.Context) {
	id := c.Param("id")
	h.DB.Exec(`DELETE FROM budgets WHERE id=$1`, id)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}

// ── Expenses ──────────────────────────────────────────────────────────────────

func (h *Handler) GetExpenses(c *gin.Context) {
	rows, _ := h.DB.Query(`SELECT id, budget_id, description, amount, category, status, submitted_by, created_at FROM expenses ORDER BY id`)
	defer rows.Close()
	var list []models.Expense
	for rows.Next() {
		var e models.Expense
		rows.Scan(&e.ID, &e.BudgetID, &e.Description, &e.Amount, &e.Category, &e.Status, &e.SubmittedBy, &e.CreatedAt)
		list = append(list, e)
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) CreateExpense(c *gin.Context) {
	var e models.Expense
	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.QueryRow(`INSERT INTO expenses (budget_id, description, amount, category, status, submitted_by) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id`,
		e.BudgetID, e.Description, e.Amount, e.Category, e.Status, e.SubmittedBy).Scan(&e.ID)
	c.JSON(http.StatusCreated, e)
}

func (h *Handler) UpdateExpense(c *gin.Context) {
	id := c.Param("id")
	var e models.Expense
	c.ShouldBindJSON(&e)
	h.DB.Exec(`UPDATE expenses SET description=$1, amount=$2, category=$3, status=$4 WHERE id=$5`,
		e.Description, e.Amount, e.Category, e.Status, id)
	c.JSON(http.StatusOK, gin.H{"message": "Updated"})
}

func (h *Handler) DeleteExpense(c *gin.Context) {
	id := c.Param("id")
	h.DB.Exec(`DELETE FROM expenses WHERE id=$1`, id)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}

// ── Invoices ──────────────────────────────────────────────────────────────────

func (h *Handler) GetInvoices(c *gin.Context) {
	rows, _ := h.DB.Query(`SELECT id, invoice_number, vendor_name, amount, due_date, status, created_at FROM invoices ORDER BY id`)
	defer rows.Close()
	var list []models.Invoice
	for rows.Next() {
		var inv models.Invoice
		rows.Scan(&inv.ID, &inv.InvoiceNumber, &inv.VendorName, &inv.Amount, &inv.DueDate, &inv.Status, &inv.CreatedAt)
		list = append(list, inv)
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) CreateInvoice(c *gin.Context) {
	var inv models.Invoice
	c.ShouldBindJSON(&inv)
	h.DB.QueryRow(`INSERT INTO invoices (invoice_number, vendor_name, amount, due_date, status) VALUES ($1,$2,$3,$4,$5) RETURNING id`,
		inv.InvoiceNumber, inv.VendorName, inv.Amount, inv.DueDate, inv.Status).Scan(&inv.ID)
	c.JSON(http.StatusCreated, inv)
}

func (h *Handler) UpdateInvoice(c *gin.Context) {
	id := c.Param("id")
	var inv models.Invoice
	c.ShouldBindJSON(&inv)
	h.DB.Exec(`UPDATE invoices SET vendor_name=$1, amount=$2, due_date=$3, status=$4 WHERE id=$5`,
		inv.VendorName, inv.Amount, inv.DueDate, inv.Status, id)
	c.JSON(http.StatusOK, gin.H{"message": "Updated"})
}

func (h *Handler) DeleteInvoice(c *gin.Context) {
	id := c.Param("id")
	h.DB.Exec(`DELETE FROM invoices WHERE id=$1`, id)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}

// ── Payroll ───────────────────────────────────────────────────────────────────

func (h *Handler) GetPayroll(c *gin.Context) {
	rows, _ := h.DB.Query(`SELECT id, employee_id, employee_name, department, basic_salary, allowances, deductions, net_salary, pay_period, status, created_at FROM payroll ORDER BY id`)
	defer rows.Close()
	var list []models.Payroll
	for rows.Next() {
		var p models.Payroll
		rows.Scan(&p.ID, &p.EmployeeID, &p.EmployeeName, &p.Department, &p.BasicSalary, &p.Allowances, &p.Deductions, &p.NetSalary, &p.PayPeriod, &p.Status, &p.CreatedAt)
		list = append(list, p)
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) CreatePayroll(c *gin.Context) {
	var p models.Payroll
	c.ShouldBindJSON(&p)
	h.DB.QueryRow(`INSERT INTO payroll (employee_id, employee_name, department, basic_salary, allowances, deductions, pay_period, status) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id`,
		p.EmployeeID, p.EmployeeName, p.Department, p.BasicSalary, p.Allowances, p.Deductions, p.PayPeriod, p.Status).Scan(&p.ID)
	c.JSON(http.StatusCreated, p)
}

func (h *Handler) UpdatePayroll(c *gin.Context) {
	id := c.Param("id")
	var p models.Payroll
	c.ShouldBindJSON(&p)
	h.DB.Exec(`UPDATE payroll SET employee_name=$1, department=$2, basic_salary=$3, allowances=$4, deductions=$5, pay_period=$6, status=$7 WHERE id=$8`,
		p.EmployeeName, p.Department, p.BasicSalary, p.Allowances, p.Deductions, p.PayPeriod, p.Status, id)
	c.JSON(http.StatusOK, gin.H{"message": "Updated"})
}

func (h *Handler) DeletePayroll(c *gin.Context) {
	id := c.Param("id")
	h.DB.Exec(`DELETE FROM payroll WHERE id=$1`, id)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}

// ── Reports ───────────────────────────────────────────────────────────────────

func (h *Handler) GetBudgetVsActual(c *gin.Context) {
	rows, _ := h.DB.Query(`SELECT department, allocated_amount, spent_amount FROM budgets`)
	defer rows.Close()
	var result []gin.H
	for rows.Next() {
		var dept string
		var allocated, spent float64
		rows.Scan(&dept, &allocated, &spent)
		result = append(result, gin.H{
			"department": dept, "allocated": allocated, "spent": spent,
		})
	}
	c.JSON(http.StatusOK, result)
}

func (h *Handler) GetExpenseByCategory(c *gin.Context) {
	rows, _ := h.DB.Query(`SELECT category, SUM(amount) as total FROM expenses GROUP BY category`)
	defer rows.Close()
	var result []gin.H
	for rows.Next() {
		var cat string
		var total float64
		rows.Scan(&cat, &total)
		result = append(result, gin.H{"category": cat, "total": total})
	}
	c.JSON(http.StatusOK, result)
}

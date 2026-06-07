
package main

import (
	"erp_finance/internal/config"
	"erp_finance/internal/handlers"
	"erp_finance/internal/middleware"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	db, err := config.ConnectDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://frontend"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	h := handlers.NewHandler(db)

	api := r.Group("/api/v1")
	{
		api.POST("/auth/login", h.Login)

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware(cfg.JWTSecret))
		{
			protected.GET("/dashboard/summary", h.GetDashboardSummary)

			protected.GET("/budgets", h.GetBudgets)
			protected.POST("/budgets", h.CreateBudget)
			protected.GET("/budgets/:id", h.GetBudgetByID)
			protected.PUT("/budgets/:id", h.UpdateBudget)
			protected.DELETE("/budgets/:id", h.DeleteBudget)

			protected.GET("/expenses", h.GetExpenses)
			protected.POST("/expenses", h.CreateExpense)
			protected.PUT("/expenses/:id", h.UpdateExpense)
			protected.DELETE("/expenses/:id", h.DeleteExpense)

			protected.GET("/invoices", h.GetInvoices)
			protected.POST("/invoices", h.CreateInvoice)
			protected.PUT("/invoices/:id", h.UpdateInvoice)
			protected.DELETE("/invoices/:id", h.DeleteInvoice)

			protected.GET("/payroll", h.GetPayroll)
			protected.POST("/payroll", h.CreatePayroll)
			protected.PUT("/payroll/:id", h.UpdatePayroll)
			protected.DELETE("/payroll/:id", h.DeletePayroll)

			protected.GET("/reports/budget-vs-actual", h.GetBudgetVsActual)
			protected.GET("/reports/expense-by-category", h.GetExpenseByCategory)
		}
	}

	log.Printf("ERP Finance backend running on port %s", cfg.Port)
	r.Run(":" + cfg.Port)
}

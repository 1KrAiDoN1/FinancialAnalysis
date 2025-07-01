package routes

import (
	"finance/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.RouterGroup, authHandler handler.AuthHandlerInterface) {
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", authHandler.SignUp)
		auth.POST("/sign-in", authHandler.SignIn)
		auth.POST("/logout", authHandler.Logout)
	}
}

func SetupExpenseRoutes(router *gin.RouterGroup, expenseHandler handler.ExpenseHandlerInterface) {
	expenses := router.Group("/categories/:category_id/expenses")
	{
		expenses.POST("", expenseHandler.CreateExpense)
		expenses.GET("", expenseHandler.GetExpenses)
		expenses.GET("/:expense_id", expenseHandler.GetExpense)
		expenses.DELETE("/:expense_id", expenseHandler.DeleteExpense)
		expenses.GET("/analytics", expenseHandler.GetAnalytics)
	}
}
func SetupCategoryRoutes(router *gin.RouterGroup, categoryHandler handler.CategoryHandlerInterface) {
	categories := router.Group("/categories")
	{
		categories.POST("", categoryHandler.CreateCategory)
		categories.GET("", categoryHandler.GetCategories)
		categories.GET("/:category_id", categoryHandler.GetCategoryByID)
		categories.GET("/top", categoryHandler.GetMostUsedCategories)
		categories.DELETE("/:category_id", categoryHandler.DeleteCategory)
		categories.GET("/analytics/:category_id", categoryHandler.GetAnalyticsByCategory)
	}
}

func SetupBudgetRoutes(router *gin.RouterGroup, budgetHandler handler.BudgetHandlerInterface) {
	budgets := router.Group("/categories/:category_id/budgets")
	{
		budgets.POST("", budgetHandler.CreateBudget)
		budgets.GET("", budgetHandler.GetBudgets)
		//budgets.PUT("/:budget_id", budgetHandler.UpdateBudget)
		budgets.DELETE("/:budget_id", budgetHandler.DeleteBudget)
	}
}

func SetupUserRoutes(router *gin.RouterGroup, userHandler handler.UserHandlerInterface) {
	users := router.Group("/user")
	{
		users.GET("/profile", userHandler.GetProfile)
		users.DELETE("/account", userHandler.DeleteAccount)
		users.GET("/stats", userHandler.GetStats)
	}
}

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
		//auth.POST("/refresh", authHandler.RefreshToken)
		auth.POST("/logout", authHandler.Logout)
	}
}

func SetupExpenseRoutes(router *gin.RouterGroup, expenseHandler handler.ExpenseHandlerInterface) {
	expenses := router.Group("/expenses")
	{
		expenses.POST("", expenseHandler.CreateExpense)
		expenses.GET("", expenseHandler.GetExpenses)
		expenses.GET("/:id", expenseHandler.GetExpense)
		expenses.DELETE("/:id", expenseHandler.DeleteExpense)
		expenses.GET("/analytics", expenseHandler.GetAnalytics)
	}
}

func SetupCategoryRoutes(router *gin.RouterGroup, categoryHandler handler.CategoryHandlerInterface) {
	categories := router.Group("/categories")
	{
		categories.POST("", categoryHandler.CreateCategory)
		categories.GET("", categoryHandler.GetCategories)
		categories.GET("/:id", categoryHandler.GetCategoryByID)
		categories.DELETE("/:id", categoryHandler.DeleteCategory)
	}
}

func SetupBudgetRoutes(router *gin.RouterGroup, budgetHandler handler.BudgetHandlerInterface) {
	budgets := router.Group("/budgets")
	{
		budgets.POST("", budgetHandler.CreateBudget)
		budgets.GET("", budgetHandler.GetBudgets)
		budgets.PUT("/:id", budgetHandler.UpdateBudget)
		budgets.DELETE("/:id", budgetHandler.DeleteBudget)
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

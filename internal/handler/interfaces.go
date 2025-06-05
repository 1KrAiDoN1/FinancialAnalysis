package handler

import (
	"github.com/gin-gonic/gin"
)

type AuthHandlerInterface interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
	Logout(c *gin.Context)
}

type BudgetHandlerInterface interface {
	CreateBudget(c *gin.Context)
	GetBudgets(c *gin.Context)
	UpdateBudget(c *gin.Context)
	DeleteBudget(c *gin.Context)
}

type CategoryHandlerInterface interface {
	CreateCategory(c *gin.Context)
	GetCategories(c *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
}

type ExpenseHandlerInterface interface {
	CreateExpense(c *gin.Context)
	GetExpenses(c *gin.Context)
	GetExpense(c *gin.Context)
	UpdateExpense(c *gin.Context)
	DeleteExpense(c *gin.Context)
	GetAnalytics(c *gin.Context)
}

type UserHandlerInterface interface {
	GetProfile(c *gin.Context)
	UpdateProfile(c *gin.Context)
	GetStats(c *gin.Context)
	DeleteAccount(c *gin.Context)
}

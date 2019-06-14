package routes

import (
	"github.com/asahasrabuddhe/rest-api/expenses/handlers"
	"github.com/asahasrabuddhe/rest-api/expenses/middleware"
	"github.com/asahasrabuddhe/rest-api/router"
)

func InitRoutes() {
	router.Post("/expenses", handlers.CreateExpense, nil)
	router.Get("/expenses", handlers.ListAllExpense, nil)
	router.Get("/expenses/{id}", handlers.ListOneExpense, middleware.ExpenseContext)
	router.Put("/expenses/{id}", handlers.UpdateExpense, middleware.ExpenseContext)
	router.Delete("/expenses/{id}", handlers.DeleteExpense, middleware.ExpenseContext)
}

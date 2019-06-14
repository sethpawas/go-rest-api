package handlers

import (
	"encoding/json"
	"github.com/asahasrabuddhe/rest-api/expenses"
	"net/http"
)

func ListOneExpense(writer http.ResponseWriter, request *http.Request) {
	encoder := json.NewEncoder(writer)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	expense := request.Context().Value("expense").(expenses.Expense)
	_ = encoder.Encode(expense)
}

func ListAllExpense(writer http.ResponseWriter, _ *http.Request) {
	encoder := json.NewEncoder(writer)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_ = encoder.Encode(expenses.Exp)
}

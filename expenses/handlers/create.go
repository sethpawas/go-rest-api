package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/asahasrabuddhe/rest-api/expenses"
	"github.com/asahasrabuddhe/rest-api/expenses/requests"
	"github.com/asahasrabuddhe/rest-api/logger"
	"github.com/go-chi/render"
	"net/http"
)

func CreateExpense(writer http.ResponseWriter, request *http.Request) {
	var req requests.CreateExpenseRequest

	err := render.Bind(request, &req)
	if err != nil {
		logger.LogEntrySetField(request, "error", err)
		return
	}

	req.Expense.Id = len(expenses.Exp) + 1
	expenses.Exp = append(expenses.Exp, *req.Expense)

	j, _ := json.Marshal(req.Expense)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)

	_, _ = fmt.Fprintf(writer, `{"success": true, "data": %v}`, string(j))
}

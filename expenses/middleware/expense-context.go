package middleware

import (
	"context"
	"github.com/asahasrabuddhe/rest-api/expenses"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func ExpenseContext (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if id := chi.URLParam(r, "id"); id != "" {
			if idInt, err := strconv.Atoi(id); err != nil {
				// error
			} else {
				for index, expense := range expenses.Exp {
					if index == idInt {
						ctx := context.WithValue(r.Context(), "expense", expense)
						next.ServeHTTP(w, r.WithContext(ctx))
					}
				}

				// not found
			}
		}
	})
}

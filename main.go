package main

import (
	"encoding/json"
	"fmt"
	"github.com/asahasrabuddhe/rest-api/requests"
	"github.com/asahasrabuddhe/rest-api/types"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

var expenses types.Expenses

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/expenses", func(r chi.Router) {
		r.Post("/", CreateExpense)
		r.Get("/", ListAllExpense)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", ListOneExpense)
			r.Put("/", UpdateExpense)
			r.Delete("/", DeleteExpense)
		})
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}

func CreateExpense(writer http.ResponseWriter, request *http.Request) {
	var req requests.CreateExpenseRequest

	err := render.Bind(request, &req)
	if err != nil {
		log.Println(err)
		return
	}

	expenses = append(expenses, *req.Expense)

	j, _ := json.Marshal(req.Expense)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)

	_, _ = fmt.Fprintf(writer, `{"success": true, "data": %v}`, string(j))
}

func ListOneExpense(writer http.ResponseWriter, request *http.Request) {

}

func ListAllExpense(writer http.ResponseWriter, request *http.Request) {
	encoder := json.NewEncoder(writer)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_ = encoder.Encode(expenses)
}

func UpdateExpense(writer http.ResponseWriter, request *http.Request) {

}

func DeleteExpense(writer http.ResponseWriter, request *http.Request) {

}

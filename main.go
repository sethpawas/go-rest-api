package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/asahasrabuddhe/rest-api/logger"
	"github.com/asahasrabuddhe/rest-api/renderers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
	"google.golang.org/genproto/googleapis/type/date"
	"net/http"
)

type Expense struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Amount      float64   `json:"amount"`
	CreatedOn   date.Date `json:"created_on" `
	UpdatedOn   date.Date `json:"updated_on"`
}

func (e *Expense) Bind(r *http.Request) error {
	return nil
}

func (e *Expense) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type Expenses []Expense

var expenses Expenses

func main() {
	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{
		DisableTimestamp: true,
		PrettyPrint:      true,
	})

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(logger.NewStructuredLogger(log))
	r.Use(middleware.Recoverer)

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
	var expense Expense

	err := render.Bind(request, &expense)
	if err != nil {
		_ = render.Render(writer, request, &renderers.ErrorResponse{
			HTTPStatusCode: http.StatusUnprocessableEntity,
			Err:            errors.New("cound not parse json request"),
		})
		return
	}

	expenses = append(expenses, expense)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)

	_, _ = fmt.Fprintln(writer, `{"success": true}`)
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

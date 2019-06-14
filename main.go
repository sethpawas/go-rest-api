package main

import (
	"github.com/asahasrabuddhe/rest-api/expenses"
	expenseRoutes "github.com/asahasrabuddhe/rest-api/expenses/routes"
	"github.com/asahasrabuddhe/rest-api/logger"
	"github.com/asahasrabuddhe/rest-api/router"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
	"net/http"
)

var exp expenses.Expenses

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
	r.Use(render.SetContentType(render.ContentTypeJSON))

	expenseRoutes.InitRoutes()
	router.RegisterRoutes(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}

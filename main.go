package main

import (
	"fmt"
	. "github.com/asahasrabuddhe/rest-api/logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
	"html"
	"net/http"
	"time"
)

func main() {
	logger := logrus.New()

	r := chi.NewRouter()
	logger.Formatter = &logrus.JSONFormatter{
		// disable, as we set our own
		DisableTimestamp: true,
	}

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(NewStructuredLogger(logger))
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	r.Get("/wait", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		LogEntrySetField(r, "wait", true)
		_, _ = fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("oops")
	})

	logger.Fatal(http.ListenAndServe(":8080", r))
}

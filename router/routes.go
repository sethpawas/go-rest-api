package router

import (
	"github.com/go-chi/chi"
	"net/http"
)

type route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Middleware  func(next http.Handler) http.Handler
}

var routes []route

func Get(pattern string, handlerFunc http.HandlerFunc, middleware func(next http.Handler) http.Handler) {
	var route route

	route.Method = "GET"
	route.Pattern = pattern
	route.HandlerFunc = handlerFunc
	route.Middleware = middleware

	routes = append(routes, route)
}

func Post(pattern string, handlerFunc http.HandlerFunc, middleware func(next http.Handler) http.Handler) {
	var route route

	route.Method = "POST"
	route.Pattern = pattern
	route.HandlerFunc = handlerFunc
	route.Middleware = middleware

	routes = append(routes, route)
}

func Put(pattern string, handlerFunc http.HandlerFunc, middleware func(next http.Handler) http.Handler) {
	var route route

	route.Method = "PUT"
	route.Pattern = pattern
	route.HandlerFunc = handlerFunc
	route.Middleware = middleware

	routes = append(routes, route)
}

func Patch(pattern string, handlerFunc http.HandlerFunc, middleware func(next http.Handler) http.Handler) {
	var route route

	route.Method = "PATCH"
	route.Pattern = pattern
	route.HandlerFunc = handlerFunc
	route.Middleware = middleware

	routes = append(routes, route)
}

func Delete(pattern string, handlerFunc http.HandlerFunc, middleware func(next http.Handler) http.Handler) {
	var route route

	route.Method = "DELETE"
	route.Pattern = pattern
	route.HandlerFunc = handlerFunc
	route.Middleware = middleware

	routes = append(routes, route)
}

func Options(pattern string, handlerFunc http.HandlerFunc, middleware func(next http.Handler) http.Handler) {
	var route route

	route.Method = "OPTIONS"
	route.Pattern = pattern
	route.HandlerFunc = handlerFunc
	route.Middleware = middleware

	routes = append(routes, route)
}

func RegisterRoutes(m *chi.Mux) {
	for _, route := range routes {
		if route.Middleware != nil {
			m.With(route.Middleware).MethodFunc(route.Method, route.Pattern, route.HandlerFunc)
		} else {
			m.MethodFunc(route.Method, route.Pattern, route.HandlerFunc)
		}
	}
}

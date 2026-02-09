package rest

import (
	"ecommerce/rest/handlers"
	"ecommerce/rest/middleware"
	"net/http"
)

func initRoutes(manager *middleware.Manager, mux *http.ServeMux) {
mux.Handle(
	"GET /rahim",
	manager.With(
		http.HandlerFunc(handlers.Test),
		middleware.Arekta,
	),
)

mux.Handle(
	"GET /products",
	manager.With(
		http.HandlerFunc(handlers.GetProducts),
		middleware.Arekta,
	),
)

mux.Handle(
	"POST /create-products",
	manager.With(
		http.HandlerFunc(handlers.CreateProduct),
		middleware.Arekta,
	),
)
}
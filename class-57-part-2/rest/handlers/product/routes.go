package product

import (
	"ecommerce/rest/handlers"
	"ecommerce/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(manager *middleware.Manager, mux *http.ServeMux) {


mux.Handle(
	"GET /products",
	manager.With(
		http.HandlerFunc(h.GetProducts),
		
	),
)

mux.Handle(
	"POST /products",
	manager.With(
		http.HandlerFunc(h.CreateProduct),
		h.middlewares.AuthenticateJWT,
	),
)
mux.Handle(
	"GET /products/{id}",
	manager.With(
		http.HandlerFunc(h.GetProduct),
		
	),
)
mux.Handle(
	"PUT /products/{id}",
	manager.With(
		http.HandlerFunc(h.UpdateProduct),
		
	),
)

mux.Handle(
	"DELETE /products/{id}",
	manager.With(
		http.HandlerFunc(h.DeleteProduct),
		
	),
)

}
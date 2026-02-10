package user

import (
	"ecommerce/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(manager *middleware.Manager, mux *http.ServeMux) {



mux.Handle(
	"POST /users",
	manager.With(
		http.HandlerFunc(h.CreateUser),
		
	),
)
mux.Handle(
	"POST /login",
	manager.With(
		http.HandlerFunc(h.Login),
		
	),
)

}
package product

import (
	"ecommerce/rest/middleware"
	// middleware "ecommerce/rest/middleware" // if code not work try it
)

type Handler struct {
	middlewares *middleware.Middlewares
	svc Service
}

func NewHandler(
	middlewares *middleware.Middlewares,
	svc Service,
	) *Handler {
	return &Handler{
		middlewares:middlewares,
		svc: svc,
	}
}
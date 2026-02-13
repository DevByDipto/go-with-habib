package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	"fmt"
	"net/http"
	"strconv"
)

type Server struct{
	cnf *config.Config
	productHandler *product.Handler
	userHandler *user.Handler
}

func NewServer( 
	cnf *config.Config,
	productHandler *product.Handler,
	userHandler *user.Handler,
	) *Server{ 
	return &Server{
		cnf : cnf,
		productHandler: productHandler,
		userHandler: userHandler,
	}
}

func (server *Server) Start() { // aikhane pointer keno use hoise ??
	mux := http.NewServeMux()

	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	wrappedMux := manager.WrapMux(mux)

	// initRoutes(manager, mux)
	server.productHandler.RegisterRoutes(manager, mux)
	server.userHandler.RegisterRoutes(manager, mux)

	addr := ":" + strconv.Itoa(server.cnf.HttpPort)
	fmt.Println("server runing", addr)

	err := http.ListenAndServe(addr, wrappedMux)

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}

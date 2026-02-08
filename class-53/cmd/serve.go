package cmd

import (
	"ecommerce/global_router"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()

manager.Use(middleware.Logger, middleware.Hudai)

mux := http.NewServeMux()

initRoutes(manager,mux)
	// mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts)) // route
	// mux.Handle("POST /create-products", http.HandlerFunc(handlers.CreateProduct))

	fmt.Println("server runing 8080")

	globalRouter := global_router.GlobalRouter(mux)
	err := http.ListenAndServe(":8080", globalRouter)

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
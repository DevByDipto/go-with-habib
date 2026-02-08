package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	wrappedMux := manager.WrapMux(mux)

    initRoutes(manager,mux)
	// mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts)) // route
	// mux.Handle("POST /create-products", http.HandlerFunc(handlers.CreateProduct))

	fmt.Println("server runing 8080")

	// globalRouter := middleware.CorsWithPreflight(mux)
	err := http.ListenAndServe(":8080", wrappedMux)

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
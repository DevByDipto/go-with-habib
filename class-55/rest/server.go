package rest

import (
	"ecommerce/config"
	"ecommerce/rest/middleware"
	"fmt"
	"net/http"
	"strconv"
)

func Start(cnf config.Config) {
	mux := http.NewServeMux()

	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	wrappedMux := manager.WrapMux(mux)

	initRoutes(manager, mux)

	addr := ":" + strconv.Itoa(cnf.HttpPort)
	fmt.Println("server runing", addr)

	err := http.ListenAndServe(addr, wrappedMux)

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}

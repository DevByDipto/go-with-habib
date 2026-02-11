package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()
	
	dbCon, err := db.NewConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Use dbCon here...
	fmt.Println("Successfully connected to the database!", dbCon)

middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares)
	userHandler := user.NewHandler()
	server := rest.NewServer(cnf,productHandler,userHandler)
	server.Start()
}
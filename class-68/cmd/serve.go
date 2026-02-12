package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	productHandler "ecommerce/rest/handlers/product"
	userHandler "ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	"ecommerce/user"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()
	
	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
err = db.MigrateDB(dbCon, "./migrations")
if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Use dbCon here...
	fmt.Println("Successfully connected to the database!", dbCon)

middlewares := middleware.NewMiddlewares(cnf)

// repos
productRepo := repo.NewProductRepo(dbCon)
userRepo := repo.NewUserRepo(dbCon)

// domains
usrSvc := user.NewService(userRepo)


	productHandler := productHandler.NewHandler(middlewares,productRepo)
	userHandler := userHandler.NewHandler(cnf,usrSvc)

	server := rest.NewServer(cnf,productHandler,userHandler)
	server.Start()
}
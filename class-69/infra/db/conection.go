package db

import (
	"ecommerce/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(cnf *config.DBConfig) string {
	connString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s",
		cnf.User,
		cnf.Password,
		cnf.Host,
		cnf.Port,
		cnf.Name,
	
	)

	// Append SSL mode if disabled
	if !cnf.EnableSSLMODE {
		connString += " sslmode=disable"
	}

	return connString
}

func NewConnection(cnf *config.DBConfig) (*sqlx.DB, error) {
	// Replace with your actual connection string logic
	dbSource := GetConnectionString(cnf) 
	
	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return dbCon, nil
}
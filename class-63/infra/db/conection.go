package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString()string{
	return "user=postgres password=123456 host=localhost port=5432 dbname=test_db sslmode=disable"
}

func NewConnection() (*sqlx.DB, error) {
	// Replace with your actual connection string logic
	dbSource := GetConnectionString() 
	
	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return dbCon, nil
}
package main

import (
	"log" // https://pkg.go.dev/log

	"database/sql"

	_ "github.com/lib/pq"
)

type DatabaseConnector struct {
	database *sql.DB
}

func (c *DatabaseConnector) Connect(connStr string) error {
	var err error
	c.database, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err) // @NOTE !!!
	} else {
		log.Println("SUCCEESS Open postgres database:", connStr)
	}
	return err
}

func (c *DatabaseConnector) Setup() {
	c.CreateTable_Customers()
	c.CreateTable_Categories()
	c.CreateTable_Products()
}

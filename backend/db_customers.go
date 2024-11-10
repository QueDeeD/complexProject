package main

import (
	"fmt"
	"log" // https://pkg.go.dev/log

	_ "github.com/lib/pq"
)

type __TableCustomersFields struct {
	id       string
	username string
	password string
}

type __TableCustomers struct {
	name   string
	fields __TableCustomersFields
}

var TableCustomers = __TableCustomers{
	name: "customers",
	fields: __TableCustomersFields{
		id:       "id",
		username: "username",
		password: "password",
	},
}

func (c *DatabaseConnector) CreateTable_Customers() error {
	rows, err := c.database.Query(
		fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			%s BIGSERIAL PRIMARY KEY,
			%s text UNIQUE NOT NULL,
			%s text NOT NULL);`,
			TableCustomers.name,
			TableCustomers.fields.id,
			TableCustomers.fields.username,
			TableCustomers.fields.password,
		),
	)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer rows.Close()
	log.Println("SUCCEESS CreateTable_Customers")
	return err
}

func (c *DatabaseConnector) IsExistingCustomer(username string, password string) bool {

	var itemExists = false

	rows, err := c.database.Query(
		fmt.Sprintf(`SELECT EXISTS(SELECT 1 from %s WHERE %s=$1 AND %s=$2);`,
			TableCustomers.name,
			TableCustomers.fields.username,
			TableCustomers.fields.password,
		),
		username,
		password,
	)
	if err != nil {
		log.Println("!!! DB Query failed:", err)
		return false
	}
	defer rows.Close()

	if !rows.Next() {
		log.Println("!!! Empty SELECT !!!")
		return false
	}
	if err = rows.Scan(&itemExists); err != nil {
		log.Println("!!! Rows Scan failed:", err)
		return false
	}
	return itemExists
}

func (c *DatabaseConnector) InsertCustomer(username string, password string) (bool, error) {

	var itemExists = false

	rowssel, err := c.database.Query(
		fmt.Sprintf(`SELECT EXISTS(SELECT 1 from %s WHERE %s=$1);`,
			TableCustomers.name,
			TableCustomers.fields.username,
		),
		username,
	)
	if err != nil {
		log.Println("!!! DB Query failed:", err)
		return itemExists, err
	}
	defer rowssel.Close()

	if !rowssel.Next() {
		return itemExists, fmt.Errorf("Empty SELECT")
	}
	if err = rowssel.Scan(&itemExists); err != nil {
		log.Println("!!! Rows Scan failed:", err)
		return itemExists, err
	}
	if itemExists {
		return itemExists, nil // @NOTE User Already Exists
	}

	rows, err := c.database.Query(
		fmt.Sprintf(`INSERT INTO %s (%s, %s) VALUES ($1, $2) ON CONFLICT (%s) DO NOTHING;`,
			TableCustomers.name,
			TableCustomers.fields.username,
			TableCustomers.fields.password,
			TableCustomers.fields.username),
		username,
		password,
	)
	if err != nil {
		log.Println("!!! DB Query failed:", err)
		return itemExists, err
	}
	defer rows.Close()

	log.Println("SUCCEESS InsertCustomer")
	return itemExists, nil
}

package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log" // https://pkg.go.dev/log
	"strings"

	_ "github.com/lib/pq"
)

type Category struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Image       string `json:"image"`
}

type __TableCategoriesFields struct {
	id          string
	name        string
	displayName string
	image       string
}

type __TableCategories struct {
	name   string
	fields __TableCategoriesFields
}

var TableCategories = __TableCategories{
	name: "categories",
	fields: __TableCategoriesFields{
		id:          "id",
		name:        "name",
		displayName: "displayName",
		image:       "image",
	},
}

func (c *DatabaseConnector) CreateTable_Categories() error {
	rows, err := c.database.Query(
		fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			%s BIGSERIAL PRIMARY KEY,
			%s text UNIQUE NOT NULL,
			%s text NOT NULL,
			%s text NOT NULL);`,
			TableCategories.name,
			TableCategories.fields.id,
			TableCategories.fields.name,
			TableCategories.fields.displayName,
			TableCategories.fields.image,
		),
	)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer rows.Close()
	log.Println("SUCCEESS CreateTable_Categories")
	return nil
}

func (c *DatabaseConnector) InsertCategory(body io.Reader) error {

	var category Category
	if err := json.NewDecoder(body).Decode(&category); err != nil {
		return err
	}

	hashImage := base64.StdEncoding.EncodeToString([]byte(category.Image))

	rows, err := c.database.Query(
		fmt.Sprintf(`INSERT INTO %s (%s, %s, %s) VALUES ($1, $2, $3) ON CONFLICT (%s) DO UPDATE SET
				%s = EXCLUDED.%s,
				%s = EXCLUDED.%s,
				%s = EXCLUDED.%s
			;`,
			TableCategories.name,
			TableCategories.fields.name,
			TableCategories.fields.displayName,
			TableCategories.fields.image,
			TableCategories.fields.name,

			TableCategories.fields.name, TableCategories.fields.name,
			TableCategories.fields.displayName, TableCategories.fields.displayName,
			TableCategories.fields.image, TableCategories.fields.image,
		),
		category.Name,
		category.DisplayName,
		hashImage,
	)
	if err != nil {
		log.Println("InsertProduct :: !!! DB Query failed:", err)
		return err
	}
	defer rows.Close()

	// https://go.dev/doc/database/execute-transactions

	log.Println("SUCCEESS InsertCategory")
	return nil
}

func (c *DatabaseConnector) IsExistingCategory(name string) (uint64, error) {

	var catID uint64

	rows, err := c.database.Query(
		fmt.Sprintf(`SELECT %s from %s WHERE %s=$1;`,
			TableCategories.fields.id,
			TableCategories.name,
			TableCategories.fields.name,
		),
		name,
	)
	if err != nil {
		log.Println("IsExistingCategory :: !!! DB Query failed:", err)
		return 0, err
	}
	defer rows.Close()

	if !rows.Next() {
		return 0, fmt.Errorf("Empty SELECT")
	}
	if err = rows.Scan(&catID); err != nil {
		log.Println("IsExistingCategory :: !!! Rows Scan failed:", err)
		return 0, err
	}
	return catID, nil
}

func parseRowCategory(rows *sql.Rows) (Category, error) {

	var category Category

	var bytes []uint8
	if err := rows.Scan(&bytes); err != nil {
		return category, err
	}

	bytes = bytes[1 : len(bytes)-1]
	text := string(bytes)
	log.Println(text)

	parts := strings.Split(text, `,`)
	log.Println(parts)

	if len(parts) == 4 {

		fmt.Sscan(parts[0], &category.Id)
		category.Name = parts[1]
		category.DisplayName = parts[2]
		if bytes, err := base64.StdEncoding.DecodeString(parts[3]); err == nil {
			category.Image = string(bytes)
		} else {
			log.Println(err)
		}
		return category, nil
	}
	return category, fmt.Errorf("Invalid Row")
}

func (c *DatabaseConnector) RetrieveCategories() ([]Category, error) {

	var categories []Category

	rows, err := c.database.Query(
		fmt.Sprintf(`SELECT (%s, %s, %s, %s) from %s;`,
			TableCategories.fields.id,
			TableCategories.fields.name,
			TableCategories.fields.displayName,
			TableCategories.fields.image,
			TableCategories.name,
		),
	)
	if err != nil {
		log.Println("RetrieveCategories :: !!! DB Query failed:", err)
		return categories, err
	}
	defer rows.Close()

	for rows.Next() {
		if category, err := parseRowCategory(rows); err == nil {
			categories = append(categories, category)
		} else {
			log.Println("RetrieveCategories :: !!! Failed to parse row:", err)
			continue
		}
	}
	return categories, nil
}

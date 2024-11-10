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

type __TableProductsFields struct {
	id          string // PK
	category_id string // FK
	vendor      string
	model       string
	price       string
	image       string
	descr       string
}

type __TableProducts struct {
	name   string
	fields __TableProductsFields
}

var TableProducts = __TableProducts{
	name: "products",
	fields: __TableProductsFields{
		id:          "id",
		category_id: "category_id",
		vendor:      "vendor",
		model:       "model",
		price:       "price",
		image:       "image",
		descr:       "descr",
	},
}

type Product struct {
	Id       uint64  `json:"id"`
	Category string  `json:"category"`
	Vendor   string  `json:"vendor"`
	Model    string  `json:"model"`
	Price    float64 `json:"price"`
	Image    string  `json:"image"`
	Descr    string  `json:"descr"`
}

func (c *DatabaseConnector) CreateTable_Products() error {
	rows, err := c.database.Query(
		fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			%s BIGSERIAL PRIMARY KEY,
			%s BIGINT,
			%s text NOT NULL,
			%s text UNIQUE NOT NULL,
			%s text NOT NULL,
			%s text NOT NULL,
			%s text NOT NULL,
			CONSTRAINT fk_%s
      			FOREIGN KEY(%s)
        		REFERENCES %s(%s));`,
			TableProducts.name,
			TableProducts.fields.id,          // PK
			TableProducts.fields.category_id, // FK
			TableProducts.fields.vendor,
			TableProducts.fields.model,
			TableProducts.fields.price,
			TableProducts.fields.image,
			TableProducts.fields.descr,
			TableCategories.name,
			TableProducts.fields.category_id, // FK
			TableCategories.name,
			TableCategories.fields.id,
		),
	)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer rows.Close()
	log.Println("SUCCEESS CreateTable_Products")
	return nil
}

func (c *DatabaseConnector) InsertProduct(body io.Reader) error {

	var product Product
	if err := json.NewDecoder(body).Decode(&product); err != nil {
		return err
	}

	categoryID, err := c.IsExistingCategory(product.Category)
	if err != nil {
		log.Println("InsertProduct :: !!!", err)
		return fmt.Errorf("Category does not exist")
	}

	hashModel := base64.StdEncoding.EncodeToString([]byte(product.Model))
	hashImage := base64.StdEncoding.EncodeToString([]byte(product.Image))
	hashDescr := base64.StdEncoding.EncodeToString([]byte(product.Descr))

	rows, err := c.database.Query(
		fmt.Sprintf(`INSERT INTO %s (%s, %s, %s, %s, %s, %s) VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT (%s) DO UPDATE SET
			%s = EXCLUDED.%s,
			%s = EXCLUDED.%s,
			%s = EXCLUDED.%s,
			%s = EXCLUDED.%s,
			%s = EXCLUDED.%s,
			%s = EXCLUDED.%s
		;`,
			TableProducts.name,
			TableProducts.fields.category_id,
			TableProducts.fields.vendor,
			TableProducts.fields.model,
			TableProducts.fields.price,
			TableProducts.fields.image,
			TableProducts.fields.descr,
			TableProducts.fields.model,

			TableProducts.fields.category_id, TableProducts.fields.category_id,
			TableProducts.fields.vendor, TableProducts.fields.vendor,
			TableProducts.fields.model, TableProducts.fields.model,
			TableProducts.fields.price, TableProducts.fields.price,
			TableProducts.fields.image, TableProducts.fields.image,
			TableProducts.fields.descr, TableProducts.fields.descr,
		),
		categoryID,
		product.Vendor,
		hashModel,
		product.Price,
		hashImage,
		hashDescr,
	)
	if err != nil {
		log.Println("InsertProduct :: !!! DB Query failed:", err)
		return err
	}
	defer rows.Close()
	log.Println("SUCCEESS InsertProduct")
	return nil
}

func parseRowProduct(rows *sql.Rows) (Product, error) {

	var product Product

	var bytes []uint8
	if err := rows.Scan(&bytes); err != nil {
		return product, err
	}

	bytes = bytes[1 : len(bytes)-1]
	text := string(bytes)
	log.Println(text)

	parts := strings.Split(text, `,`)
	log.Println(parts)

	if len(parts) == 6 {

		fmt.Sscan(parts[0], &product.Id)
		product.Vendor = parts[1]
		if bytes, err := base64.StdEncoding.DecodeString(parts[2]); err == nil {
			product.Model = string(bytes)
		} else {
			log.Println(err)
		}
		fmt.Sscan(parts[3], &product.Price)
		if bytes, err := base64.StdEncoding.DecodeString(parts[4]); err == nil {
			product.Image = string(bytes)
		} else {
			log.Println(err)
		}
		if bytes, err := base64.StdEncoding.DecodeString(parts[5]); err == nil {
			product.Descr = string(bytes)
		} else {
			log.Println(err)
		}
		return product, nil
	}
	return product, fmt.Errorf("Invalid Row")
}

func (c *DatabaseConnector) RetrieveProductsFromCategory(category_id uint64) ([]Product, error) {

	var products []Product

	rows, err := c.database.Query(
		fmt.Sprintf(`SELECT (%s, %s, %s, %s, %s, %s) FROM %s WHERE %s=$1;`,
			TableProducts.fields.id,
			TableProducts.fields.vendor,
			TableProducts.fields.model,
			TableProducts.fields.price,
			TableProducts.fields.image,
			TableProducts.fields.descr,
			TableProducts.name,
			TableProducts.fields.category_id,
		),
		category_id,
	)
	if err != nil {
		log.Println("RetrieveProductsFromCategory :: !!! DB Query failed:", err)
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		if product, err := parseRowProduct(rows); err == nil {
			products = append(products, product)
		} else {
			log.Println("RetrieveProductsFromCategory :: !!! Failed to parse row:", err)
			continue
		}
	}
	return products, nil
}

func (c *DatabaseConnector) RetrieveProduct(product_id uint64) (Product, error) {

	var product Product

	rows, err := c.database.Query(
		fmt.Sprintf(`SELECT (%s, %s, %s, %s, %s, %s) FROM %s WHERE %s=$1;`,
			TableProducts.fields.id,
			TableProducts.fields.vendor,
			TableProducts.fields.model,
			TableProducts.fields.price,
			TableProducts.fields.image,
			TableProducts.fields.descr,
			TableProducts.name,
			TableProducts.fields.id,
		),
		product_id,
	)
	if err != nil {
		log.Println("RetrieveProduct :: !!! DB Query failed:", err)
		return product, err
	}
	defer rows.Close()

	if rows.Next() {
		if product, err := parseRowProduct(rows); err == nil {
			return product, nil
		} else {
			log.Println("RetrieveProductsFromCategory :: !!! Failed to parse row:", err)
		}
	}
	return product, fmt.Errorf("Empty SELECT")
}

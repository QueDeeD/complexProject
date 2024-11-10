package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func perform(client *http.Client, meth string, url string, b *bytes.Buffer) error {
	req, err := http.NewRequest(meth, url, b)
	if err != nil {
		log.Println(err)
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	log.Println("RESPONSE:", buf.String())
	return nil
}

type Category struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Image       string `json:"image"`
}

func CreateCategories(file string) {

	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatalln(err)
		return
	}

	var categories []Category
	if err := json.Unmarshal(data, &categories); err != nil {
		log.Fatalln(err)
		return
	}

	client := &http.Client{}
	for id, category := range categories {
		log.Println("CATEGORY #", id, category)

		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(category)
		if err = perform(client, "POST", "http://localhost:8090/category/create", b); err != nil {
			break
		}
	}
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

func AddProducts(file string) {

	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatalln(err)
		return
	}

	var products []Product
	if err := json.Unmarshal(data, &products); err != nil {
		log.Fatalln(err)
		return
	}

	client := &http.Client{}
	for id, product := range products {
		log.Println("PRODUCT #", id, product)

		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(product)
		if err = perform(client, "POST", "http://localhost:8090/product/add", b); err != nil {
			break
		}
	}
}

var (
	mode string
	file string
)

func main() {
	flag.StringVar(&mode, "mode", "default", "help message")
	flag.StringVar(&file, "file", "default", "help message")
	flag.Parse()

	fmt.Println("mode value is: ", mode)
	fmt.Println("file value is: ", file)

	if mode == "cat" {
		CreateCategories(file)

	} else if mode == "prod" {
		AddProducts(file)
	}
}

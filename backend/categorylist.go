package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log" // https://pkg.go.dev/log
	"net/http"

	_ "github.com/lib/pq"
)

/**
 * \brief
 */
type CategoryListHandler struct {
	dbconn *DatabaseConnector
}

func (h CategoryListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("HTTP Request `CategoryList`")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	categories, err := h.dbconn.RetrieveCategories()
	if err != nil {
		io.WriteString(w, "{}")
		return
	}

	if bytes, err := json.Marshal(categories); err != nil {
		// @NOTE Should NEVER happen !!!
		io.WriteString(w, "{}")
	} else {
		log.Println(string(bytes))
		io.Writer.Write(w, bytes)
	}
}

/**
 * \brief
 */
type CategoryListByIDHandler struct {
	dbconn *DatabaseConnector
}

func (h CategoryListByIDHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("HTTP Request `CategoryListByID`")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	page := r.PathValue("page")

	var category_id uint64
	fmt.Sscan(page, &category_id)

	products, err := h.dbconn.RetrieveProductsFromCategory(category_id)
	if err != nil {
		io.WriteString(w, "{}")
		return
	}

	if bytes, err := json.Marshal(products); err != nil {
		// @NOTE Should NEVER happen !!!
		io.WriteString(w, "{}")
	} else {
		log.Println(string(bytes))
		io.Writer.Write(w, bytes)
	}
}

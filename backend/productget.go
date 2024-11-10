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
type ProductGetHandler struct {
	dbconn *DatabaseConnector
}

func (h ProductGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("HTTP Request `ProductGet`")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	page := r.PathValue("page")

	var product_id uint64
	fmt.Sscan(page, &product_id)

	product, err := h.dbconn.RetrieveProduct(product_id)
	if err != nil {
		io.WriteString(w, "{}")
		return
	}

	if bytes, err := json.Marshal(product); err != nil {
		// @NOTE Should NEVER happen !!!
		io.WriteString(w, "{}")
	} else {
		log.Println(string(bytes))
		io.Writer.Write(w, bytes)
	}
}

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
type ProductAddHandler struct {
	dbconn *DatabaseConnector
}

type ProductAddResponse struct {
	StatusCode ErrorCode `json:"statusCode"`
	Message    string    `json:"message"`
}

func ProductAddResponse_Success() ProductAddResponse {
	return ProductAddResponse{
		StatusCode: SUCCEEDED,
		Message:    "OK",
	}
}
func ProductAddResponse_Error(err error) ProductAddResponse {
	return ProductAddResponse{
		StatusCode: FAILED,
		Message:    err.Error(),
	}
}

func (r *ProductAddResponse) write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	if bytes, err := json.Marshal(r); err != nil {
		// @NOTE Should NEVER happen !!!
		io.WriteString(w, fmt.Sprintf("{\"statusCode\":%d,\"message\":\"%s\"}", r.StatusCode, r.Message))
	} else {
		log.Println(string(bytes))
		io.Writer.Write(w, bytes)
	}
}

func (h ProductAddHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("HTTP Request `ProductAdd`")

	var resp ProductAddResponse
	defer resp.write(w)

	if err := h.dbconn.InsertProduct(r.Body); err != nil {
		resp = ProductAddResponse_Error(err)
		return
	}
	resp = ProductAddResponse_Success()
}

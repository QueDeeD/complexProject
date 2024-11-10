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
type CategoryCreateHandler struct {
	dbconn *DatabaseConnector
}

type CategoryCreateResponse struct {
	StatusCode ErrorCode `json:"statusCode"`
	Message    string    `json:"message"`
}

func (r *CategoryCreateResponse) write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	if bytes, err := json.Marshal(r); err != nil {
		// @NOTE Should NEVER happen !!!
		io.WriteString(w, fmt.Sprintf("{\"statusCode\":%d,\"message\":\"%s\"}", r.StatusCode, r.Message))
	} else {
		log.Println(string(bytes))
		io.Writer.Write(w, bytes)
	}
}

func CategoryCreateResponse_Success() CategoryCreateResponse {
	return CategoryCreateResponse{
		StatusCode: SUCCEEDED,
		Message:    "OK",
	}
}
func CategoryCreateResponse_Error(err error) CategoryCreateResponse {
	return CategoryCreateResponse{
		StatusCode: FAILED,
		Message:    err.Error(),
	}
}

func (h CategoryCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("HTTP Request `CategoryCreate`")

	var resp CategoryCreateResponse
	defer resp.write(w)

	if err := h.dbconn.InsertCategory(r.Body); err != nil {
		resp = CategoryCreateResponse_Error(err)
	} else {
		resp = CategoryCreateResponse_Success()
	}
}

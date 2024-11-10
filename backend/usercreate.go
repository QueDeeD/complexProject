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
type UserCreateHandler struct {
	dbconn *DatabaseConnector
}

type UserCreateResponse struct {
	StatusCode ErrorCode `json:"statusCode"`
	Message    string    `json:"message"`
}

func (r *UserCreateResponse) write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	if bytes, err := json.Marshal(r); err != nil {
		// @NOTE Should NEVER happen !!!
		io.WriteString(w, fmt.Sprintf("{\"statusCode\":%d,\"message\":\"%s\"}", r.StatusCode, r.Message))
	} else {
		log.Println(string(bytes))
		io.Writer.Write(w, bytes)
	}
}

func UserCreateResponse_Success() UserCreateResponse {
	return UserCreateResponse{
		StatusCode: SUCCEEDED,
		Message:    "OK",
	}
}
func UserCreateResponse_Error(err error) UserCreateResponse {
	return UserCreateResponse{
		StatusCode: FAILED,
		Message:    err.Error(),
	}
}
func UserCreateResponse_Exists() UserCreateResponse {
	return UserCreateResponse{
		StatusCode: EXISTS_ALREADY,
		Message:    "User already exists",
	}
}

func (h UserCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("HTTP Request `UserCreate`")

	var resp UserCreateResponse
	defer resp.write(w)

	params, err := ParseQuery(r)
	if err != nil {
		resp = UserCreateResponse_Error(err)
		return
	}

	username, unameFound := params[TableCustomers.fields.username]
	password, paswdFound := params[TableCustomers.fields.password]

	if unameFound && paswdFound {
		var itemExists bool
		if itemExists, err = h.dbconn.InsertCustomer(username[0], password[0]); err != nil {
			resp = UserCreateResponse_Error(err)

		} else {
			if itemExists {
				resp = UserCreateResponse_Exists()
			} else {
				resp = UserCreateResponse_Success()
			}
		}
		return
	}
	resp = UserCreateResponse_Error(fmt.Errorf("Invalid Query"))
}

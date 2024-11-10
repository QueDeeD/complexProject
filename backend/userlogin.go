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
type UserLoginHandler struct {
	dbconn *DatabaseConnector
}

type UserLoginResponse struct {
	StatusCode  ErrorCode `json:"statusCode"`
	Message     string    `json:"message"`
	DisplayName string    `json:"displayName"`
	AccessToken string    `json:"accessToken"`
}

func (r *UserLoginResponse) write(w http.ResponseWriter) {
	if bytes, err := json.Marshal(r); err != nil {
		// @NOTE Should NEVER happen !!!
		io.WriteString(w, fmt.Sprintf("{\"statusCode\":%d,\"message\":\"%s\",\"displayName\":\"%s\",\"accessToken\":\"%s\"}",
			r.StatusCode, r.Message, r.DisplayName, r.AccessToken))
	} else {
		log.Println(string(bytes))
		io.Writer.Write(w, bytes)
	}
}

func UserLoginResponse_Success(username string, token string) UserLoginResponse {
	return UserLoginResponse{
		StatusCode:  SUCCEEDED,
		Message:     "OK",
		DisplayName: username,
		AccessToken: token,
	}
}
func UserLoginResponse_Error(username string, err error) UserLoginResponse {
	return UserLoginResponse{
		StatusCode:  FAILED,
		Message:     err.Error(),
		DisplayName: username,
		AccessToken: "",
	}
}
func UserLoginResponse_NotExists(username string) UserLoginResponse {
	return UserLoginResponse{
		StatusCode:  NOT_EXISTS,
		Message:     "User does not exists",
		DisplayName: username,
		AccessToken: "",
	}
}

func (h UserLoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("HTTP Request `UserLogin`")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var resp UserLoginResponse
	defer resp.write(w)

	params, err := ParseQuery(r)
	if err != nil {
		resp = UserLoginResponse_Error("", err)
		return
	}
	if username, found := params[TableCustomers.fields.username]; found {
		log.Println("UserLogin ", username)
		if password, found := params[TableCustomers.fields.password]; found {
			log.Println("UserLogin ", password) // @TODO !!! Encrypt Password !!!

			if h.dbconn.IsExistingCustomer(username[0], password[0]) {
				if token, err := CreateToken(username[0]); err != nil {
					resp = UserLoginResponse_Error(username[0], err)
				} else {
					resp = UserLoginResponse_Success(username[0], token)
				}
			} else {
				resp = UserLoginResponse_NotExists(username[0])
			}
			return
		}
	}
	resp = UserLoginResponse_Error("", fmt.Errorf("Invalid Query"))
}

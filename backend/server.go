package main

import (
	"log" // https://pkg.go.dev/log
	"net/http"
	"net/url"

	_ "github.com/lib/pq"
)

func ParseQuery(r *http.Request) (url.Values, error) {
	err := r.ParseForm()
	if err != nil {
		log.Println("!!! Parse Query failed:", err)
		return url.Values{}, err
	}
	params := r.Form
	log.Println("ParseQuery params count:", len(params))
	for key, values := range params {
		log.Println("ParseQuery param", key, ":", values)
	}
	return params, nil
}

type ErrorCode int

const (
	SUCCEEDED      ErrorCode = 0
	EXISTS_ALREADY ErrorCode = 1
	NOT_EXISTS     ErrorCode = 2
	FAILED         ErrorCode = -1
)

func main() {

	var dbconn DatabaseConnector
	connStr := "postgres://CyberStoreVueApp-develop:qqq@localhost/CyberStoreVueApp?sslmode=disable"
	dbconn.Connect(connStr)
	dbconn.Setup()

	userCreateHandler := UserCreateHandler{
		dbconn: &dbconn,
	}
	userLoginHandler := UserLoginHandler{
		dbconn: &dbconn,
	}

	categoryCreateHandler := CategoryCreateHandler{
		dbconn: &dbconn,
	}
	categoryListHandler := CategoryListHandler{
		dbconn: &dbconn,
	}
	categoryListByIDHandler := CategoryListByIDHandler{
		dbconn: &dbconn,
	}

	productAddHandler := ProductAddHandler{
		dbconn: &dbconn,
	}
	productGetHandler := ProductGetHandler{
		dbconn: &dbconn,
	}

	mux := http.NewServeMux()
	mux.Handle("/user/create", userCreateHandler)
	mux.Handle("/user/login", userLoginHandler)

	mux.Handle("POST /category/create", categoryCreateHandler)
	mux.Handle("GET /categories", categoryListHandler)
	mux.Handle("GET /category/list/{page}", categoryListByIDHandler)

	mux.Handle("POST /product/add", productAddHandler)
	mux.Handle("GET /product/get/{page}", productGetHandler)

	http.ListenAndServe(":8090", mux)
}

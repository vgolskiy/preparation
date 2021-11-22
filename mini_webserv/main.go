package main

import (
	H "./handlers" //H - handlers for request/response
	M "./model"    //M - model
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strings"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "vladimirgolskiy"
	password = "ozon"
	dbname   = "ozon"
)

func splitFn(c rune) bool {
	return c == '/'
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		parts := strings.FieldsFunc(r.URL.Path, splitFn)
		switch len(parts) {
		case 1:
			H.HandleFactAny(w)
		case 2:
			H.HandleFactById(w, parts[1])
		default:
			http.Error(w, "404 not found", http.StatusNotFound)
		}
	case "POST":
		switch r.URL.Path {
		case "/fact":
			H.HandleFactsInsertion(w, r)
		default:
			http.Error(w, "404 not found", http.StatusNotFound)
		}
	default:
		http.Error(w, "Only GET and POST methods are supported", http.StatusMethodNotAllowed)
	}

}

func main() {
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	M.DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer M.DB.Close()

	err = M.DB.Ping()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

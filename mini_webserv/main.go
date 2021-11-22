package main

import (
	M "./model"
	R "./support"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
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
	fmt.Println(r.URL, r.Method, r.Body)
	switch r.Method {
	case "GET":
		parts := strings.FieldsFunc(r.URL.Path, splitFn)
		switch len(parts) {
		case 1:
			fact, err := M.FactLastOne()
			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			} else {
				R.ReturnFact(w, fact)
			}
		case 2:
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				http.Error(w, "Bad request: fact identifier should be an integer, got " + parts[1], http.StatusBadRequest)
			} else {
				fact, err := M.FactByID(id)
				if err != nil {
					http.Error(w, "Bad request: "+err.Error(), http.StatusBadRequest)
				} else {
					R.ReturnFact(w, fact)
				}
			}
		default:
			http.Error(w, "404 not found", http.StatusNotFound)
		}
	case "POST":
		switch r.URL.Path {
		case "/fact":
			R.ReadFacts(w, r)
		default:
			http.Error(w, "404 not found", http.StatusNotFound)
		}
	default:
		http.Error(w, "Only GET and POST methods are supported", http.StatusMethodNotAllowed)
	}

}

func main() {
	var err error
	// DB connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// DB connection string validation
	M.DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// DB connection closing deferring
	defer M.DB.Close()

	// Creating request to DB
	err = M.DB.Ping()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

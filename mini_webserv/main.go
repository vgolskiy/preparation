package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type SFact struct {
	Id          int      `json:"id,omitempty"`
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Links       []string `json:"links,omitempty"`
}

type PsqlFact struct {
	SFact
	Links string
}

type InFacts struct {
	Facts []SFact
}

var DB *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "vladimirgolskiy"
	password = "ozon"
	dbname   = "ozon"
)

func returnFact(w http.ResponseWriter, fact SFact) {
	jData, err := json.Marshal(fact)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(jData))
}

func readFacts(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var facts InFacts
	err := decoder.Decode(&facts)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	jData, err := json.Marshal(facts)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(jData))
}

func convertToFact(fact *SFact, tmp *PsqlFact) {
	fact.Id = tmp.Id
	fact.Title = tmp.Title
	fact.Description = tmp.Description
	fact.Links = strings.Split(tmp.Links, ",")
}

func factLastOne() (SFact, error) {
	var fact SFact
	var tmp PsqlFact

	script := `
	SELECT *
	FROM facts
	ORDER BY id DESC
	LIMIT 1`
	row := DB.QueryRow(script)
	if err := row.Scan(&tmp.Id, &tmp.Title, &tmp.Description, &tmp.Links); err != nil {
		if err == sql.ErrNoRows {
			return fact, fmt.Errorf("DB does not contain any facts")
		}
		return fact, fmt.Errorf("facts: %v", err)
	}
	convertToFact(&fact, &tmp)
	return fact, nil
}

func factByID(id int) (SFact, error) {
	var fact SFact
	var tmp PsqlFact

	script := `
	SELECT *
	FROM facts
	WHERE id = $1`
	row := DB.QueryRow(script, id)
	if err := row.Scan(&tmp.Id, &tmp.Title, &tmp.Description, &tmp.Links); err != nil {
		if err == sql.ErrNoRows {
			return fact, fmt.Errorf("DB does not contain such fact %d", id)
		}
		return fact, fmt.Errorf("facts %d: %v", id, err)
	}
	convertToFact(&fact, &tmp)
	return fact, nil
}

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
			fact, err := factLastOne()
			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}
			returnFact(w, fact)
		case 2:
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				http.Error(w, "Bad request: "+err.Error(), http.StatusBadRequest)
			}
			fact, err := factByID(id)
			if err != nil {
				http.Error(w, "Bad request: "+err.Error(), http.StatusBadRequest)
			}
			returnFact(w, fact)
		default:
			http.Error(w, "404 not found", http.StatusNotFound)
		}
	case "POST":
		switch r.URL.Path {
		case "/fact":
			readFacts(w, r)
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
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// DB connection closing deferring
	defer DB.Close()

	// Creating request to DB
	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

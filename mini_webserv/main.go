package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type SFact struct {
	Id          int      `json:"id,omitempty"`
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Links       []string `json:"links,omitempty"`
}

type InFacts struct {
	Facts []SFact
}

var fact = SFact{
	Id:          1,
	Title:       "високосная секунда",
	Description: "Дополнительная (висоќосная, скачущая) сеќунда[4][5][6], или сеќунда координ́ации[7] (англ. leap second) — секунда, иногда добавляемая (теоретически возможно и вычитание) в шкалу всемирного координированного времени (UTC) для согласования его со средним солнечным временем UT1.",
	Links:       []string{"shorturl.at/dwzK9"},
}

const (
	host     = "localhost"
	port     = 5432
	user     = "vladimirgolskiy"
	password = "ozon"
	dbname   = "ozon"
)

func getFact(w http.ResponseWriter) {
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

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL, r.Method, r.Body)
	switch r.Method {
	case "GET":
		switch r.URL.Path {
		case "/fact":
			getFact(w)
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
	// DB connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// DB connection string validation
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// DB connection closing deferring
	defer db.Close()

	// Creating request to DB
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

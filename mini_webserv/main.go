package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Sfact struct {
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Links       []string `json:"links,omitempty"`
}

type Infacts struct {
	Facts []Sfact
}

var fact = Sfact{
	Title:       "високосная секунда",
	Description: "Дополнительная (висоќосная, скачущая) сеќунда[4][5][6], или сеќунда координ́ации[7] (англ. leap second) — секунда, иногда добавляемая (теоретически возможно и вычитание) в шкалу всемирного координированного времени (UTC) для согласования его со средним солнечным временем UT1.",
	Links:       []string{"shorturl.at/dwzK9"},
}

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
	var facts Infacts
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
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

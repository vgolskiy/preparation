package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Sfact struct {
	Id          int
	Title       string
	Description string
	Links       []string
}

var fact = Sfact{
	Id:          1,
	Title:       "високосная секунда",
	Description: "Дополнительная (висоќосная, скачущая) сеќунда[4][5][6], или сеќунда координ́ации[7] (англ. leap second) — секунда, иногда добавляемая (теоретически возможно и вычитание) в шкалу всемирного координированного времени (UTC) для согласования его со средним солнечным временем UT1.",
	Links:       []string{"shorturl.at/dwzK9"},
}

type test_struct struct {
	Name	string
	Address string
}

func getFact(w http.ResponseWriter) {
	jData, err := json.Marshal(fact)
	if err != nil {
		//TODO
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(jData))
}

func fill_fact(w http.ResponseWriter, r *http.Request) {
	// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func read_from_body(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t test_struct
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "\n%s %s\n", t.Name, t.Address)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL, r.Method, r.Body)
	switch r.Method {
	case "GET":
		switch r.URL.Path {
		case "/":
			http.ServeFile(w, r, "./mini_webserv/form.html")
		case "/fact":
			getFact(w)
			}
	case "POST":
		switch r.URL.Path {
		case "/fill_fact":
			fill_fact(w, r)
		case "/fact":
			read_from_body(w, r)
		default:
			http.Error(w, "404 not found.", http.StatusNotFound)
		}
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

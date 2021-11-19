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

func getFact(w http.ResponseWriter, r *http.Request) {
	jData, err := json.Marshal(fact)
	if err != nil {
		//TODO
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(jData))
}

func main() {
	http.HandleFunc("/fact", getFact)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

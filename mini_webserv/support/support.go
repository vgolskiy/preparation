package support

import (
	S "../structs"
	"encoding/json"
	"fmt"
	"net/http"
)

func ReturnFact(w http.ResponseWriter, fact S.SFact) {
	jData, err := json.Marshal(fact)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(jData))
}

func ReadFacts(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var facts S.InFacts
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

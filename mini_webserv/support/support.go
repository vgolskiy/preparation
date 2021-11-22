package support

import (
	S "../structs"
	"encoding/json"
	"fmt"
	"net/http"
)

func ReturnData(w http.ResponseWriter, data interface{}) {
	jData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(jData))
}

func ReadFacts(r *http.Request) ([]S.SFact, error) {
	decoder := json.NewDecoder(r.Body)
	var facts S.InFacts
	err := decoder.Decode(&facts)
	if err != nil {
		return []S.SFact{}, err
	}
	return facts.Facts, nil
}

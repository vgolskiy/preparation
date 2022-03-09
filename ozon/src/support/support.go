package support

import (
	S "ozon/src/structs"
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
	var fact S.SFact
	decoder := json.NewDecoder(r.Body)
	var facts S.InFacts
	err := decoder.Decode(&facts)
	if err != nil {
		return []S.SFact{}, err
	}
	for _, fact = range facts {
		if len(fact.Title) == 0 || len(fact.Description) == 0 {
			err = fmt.Errorf("Required field is empty")
			return []S.SFact{}, err
		}
	}
	return facts.Facts, nil
}

func ReadFact(r *http.Request) (S.SFact, error) {
	decoder := json.NewDecoder(r.Body)
	var fact S.SFact
	err := decoder.Decode(&fact)
	if err != nil {
		return fact, err
	}
	if len(fact.Title) == 0 || len(fact.Description) == 0 {
		err = fmt.Errorf("Required field is empty")
		return fact, err
	}
	return fact, nil
}

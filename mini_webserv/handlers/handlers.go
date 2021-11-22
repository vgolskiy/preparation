package handlers

import (
	M "../model"   //M - model
	S "../structs" //S - structs
	R "../support" //R - request/response
	"net/http"
	"strconv"
)

func HandleFactAny(w http.ResponseWriter) {
	fact, err := M.FactLastOne()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	} else {
		R.ReturnData(w, fact)
	}
}

func HandleFactById(w http.ResponseWriter, s string) {
	id, err := strconv.Atoi(s)
	if err != nil {
		http.Error(w, "Bad request: fact identifier should be an integer, got " + s, http.StatusBadRequest)
	} else {
		fact, err := M.FactByID(id)
		if err != nil {
			http.Error(w, "Bad request: "+err.Error(), http.StatusBadRequest)
		} else {
			R.ReturnData(w, fact)
		}
	}
}

func HandleFactsInsertion(w http.ResponseWriter, r *http.Request) {
	var (
		res S.WriteFactResult
		err error
		facts []S.SFact
	)

	facts, err = R.ReadFacts(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		res.Ids, err = M.FactInsert(facts)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		R.ReturnData(w, res)
	}
}

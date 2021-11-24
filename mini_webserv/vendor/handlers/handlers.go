package handlers

import (
	M "model"   //M - model
	S "structs" //S - structs
	R "support" //R - request/response
	"net/http"
	"strconv"
	"strings"
)

func handleFactAny(w http.ResponseWriter) {
	fact, err := M.FactLastOne()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	} else {
		R.ReturnData(w, fact)
	}
}

func handleFactById(w http.ResponseWriter, s string) {
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

func handleFactsInsertion(w http.ResponseWriter, r *http.Request) {
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

func handleFactUpdate(w http.ResponseWriter, r *http.Request, s string) {
	var (
		fact	S.SFact
		err		error
		id		int
	)

	fact, err = R.ReadFact(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		id, err = strconv.Atoi(s)
		if err != nil {
			http.Error(w, "Bad request: fact identifier should be an integer, got "+s, http.StatusBadRequest)
		} else {
			fact, err := M.FactUpdate(id, fact)
			if err != nil {
				http.Error(w, "Bad request: "+err.Error(), http.StatusBadRequest)
			} else {
				R.ReturnData(w, fact)
			}
		}
	}
}

func splitFn(c rune) bool {
	return c == '/'
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		parts := strings.FieldsFunc(r.URL.Path, splitFn)
		switch len(parts) {
		case 1:
			handleFactAny(w)
		case 2:
			handleFactById(w, parts[1])
		default:
			http.Error(w, "404 not found", http.StatusNotFound)
		}
	case "POST":
		switch r.URL.Path {
		case "/fact":
			handleFactsInsertion(w, r)
		default:
			http.Error(w, "404 not found", http.StatusNotFound)
		}
	case "PUT":
		parts := strings.FieldsFunc(r.URL.Path, splitFn)
		switch len(parts) {
		case 2:
			handleFactUpdate(w, r, parts[1])
		default:
			http.Error(w, "404 not found", http.StatusNotFound)
		}
	default:
		http.Error(w, "Only GET, POST, PUT methods are supported", http.StatusMethodNotAllowed)
	}
}

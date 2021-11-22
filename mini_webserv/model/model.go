package model

import (
	S "../structs"
	"database/sql"
	"fmt"
	"strings"
)

var DB *sql.DB
var SEPARATOR = ", "

func convertToFact(fact *S.SFact, tmp *S.PsqlFact) {
	fact.Id = tmp.Id
	fact.Title = tmp.Title
	fact.Description = tmp.Description
	if len(tmp.Links) > 0 {
		fact.Links = strings.Split(tmp.Links, SEPARATOR)
	}
}

func FactLastOne() (S.SFact, error) {
	var fact	S.SFact
	var tmp		S.PsqlFact

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

func FactByID(id int) (S.SFact, error) {
	var fact	S.SFact
	var tmp		S.PsqlFact

	script := `
	SELECT *
	FROM facts
	WHERE id = $1`
	row := DB.QueryRow(script, id)
	if err := row.Scan(&tmp.Id, &tmp.Title, &tmp.Description, &tmp.Links); err != nil {
		if err == sql.ErrNoRows {
			return fact, fmt.Errorf("DB does not contain fact #%d", id)
		}
		return fact, fmt.Errorf("facts %d: %v", id, err)
	}
	convertToFact(&fact, &tmp)
	return fact, nil
}

func FactInsert(facts []S.SFact) ([]int, error) {
	var (
		id		int
		ids		[]int
		err		error
		fact	S.SFact
		s		string
	)

	script := `
	INSERT INTO facts (title, description, links)
	VALUES ($1, $2, $3)
	RETURNING id`

	for _, fact = range facts {
		if len(fact.Links) > 0 {
			s = strings.Join(fact.Links, SEPARATOR)
		}
		err = DB.QueryRow(script, fact.Title, fact.Description, s).Scan(&id)
		if err != nil {
			return ids, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

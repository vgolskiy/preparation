package structs

type SFact struct {
	Id          int      `json:"id,omitempty"`
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Links       []string `json:"links,omitempty"`
}

type InFacts struct {
	Facts []SFact
}

type PsqlFact struct {
	SFact
	Links string
}

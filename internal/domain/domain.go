package domain

// Domain
type InfoBook struct {
	Nome          string `json:"nome"`
	Autor         string `json:"autor"`
	Edicao        string `json:"edicao"`
	AnoLancamento int    `json:"anoLancamento"`
}

type Book struct {
	Id            string `json:"id"`
	Nome          string `json:"nome"`
	Autor         string `json:"autor"`
	Edicao        string `json:"edicao"`
	AnoLancamento int    `json:"anoLancamento"`
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	log.Println("Loading in server...")
	router.Post("/book", CreateBook)
	router.Put("/book/{id}", UpdateBook)
	router.Delete("/book/{id}", DeleteBook)
	router.Get("/book/{id}", DescribeBook)
	router.Get("/books", ListBook)

	log.Fatal(http.ListenAndServe(":8080", router))
}

// Domain
type Book struct {
	Id            string `json:"id"`
	Nome          string `json:"nome"`
	Autor         string `json:"autor"`
	Edicao        string `json:"edicao"`
	AnoLancamento int    `json:"ano lancamento"`
}

type InfoBook struct {
	Nome          string `json:"nome"`
	Autor         string `json:"autor"`
	Edicao        string `json:"edicao"`
	AnoLancamento int    `json:"ano lancamento"`
}

// Controller
func CreateBook(w http.ResponseWriter, r *http.Request) {
	infoBook := InfoBook{}
	err := json.NewDecoder(r.Body).Decode(&infoBook)
	if err != nil {
		log.Printf("Erro ao Cadastrar livro %v", err)

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("json", "content-type")
		return
	}
	book := Book{
		Id:            "01",
		Nome:          infoBook.Nome,
		Autor:         infoBook.Autor,
		Edicao:        infoBook.Edicao,
		AnoLancamento: infoBook.AnoLancamento,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
	w.WriteHeader(http.StatusCreated)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	infoBook := InfoBook{}
	id := chi.URLParam(r, "id")
	err := json.NewDecoder(r.Body).Decode(&infoBook)
	if err != nil {
		log.Printf("Erro ao atualizar livro %v", err)

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("json", "content-type")
		return
	}
	book := Book{
		Id:            id,
		Nome:          infoBook.Nome,
		Autor:         infoBook.Autor,
		Edicao:        infoBook.Edicao,
		AnoLancamento: infoBook.AnoLancamento,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
	w.WriteHeader(http.StatusOK)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	resp := fmt.Sprint("Livro deletado com sucesso", id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
	w.WriteHeader(http.StatusNoContent)
}

func DescribeBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	book := Book{
		Id:            id,
		Nome:          "O Diário de Anne Frank",
		Autor:         "JJ",
		Edicao:        "1990.3",
		AnoLancamento: 2000,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
	w.WriteHeader(http.StatusOK)
}

func ListBook(w http.ResponseWriter, r *http.Request) {
	listBook := []Book{
		{
			Id:            "01",
			Nome:          "A Culpa é das Estrelas",
			Autor:         "John Green",
			Edicao:        "2014.3",
			AnoLancamento: 2014,
		},
		{
			Id:            "02",
			Nome:          "O Diário de Anne Frank",
			Autor:         "Ane Frank",
			Edicao:        "1995.3",
			AnoLancamento: 1995,
		},
		{
			Id:            "03",
			Nome:          "O Código Da Vinci (Robert Langdon)",
			Autor:         "Dan Brown",
			Edicao:        "2000-1",
			AnoLancamento: 2003,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(listBook)
	w.WriteHeader(http.StatusOK)
}

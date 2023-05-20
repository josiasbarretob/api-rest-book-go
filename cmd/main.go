package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Rotas
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
		log.Printf("Error ao Cadastrar livro %v", err)

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("json", "content-type")
		return
	}

	book, err := CreateBookService(infoBook)
	if err != nil {
		log.Printf("Error ao criar livro")
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
		log.Printf("Error ao atualizar livro %v", err)

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("json", "content-type")
		return
	}
	book, err := UpdateBookService(id, infoBook)
	if err != nil {
		log.Printf("Error ao atualizar livro")
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
	w.WriteHeader(http.StatusOK)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := DeleteBookService(id)
	if err != nil {
		log.Printf("Error ao deletar livro")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func DescribeBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	book, err := DescribeBookService(id)
	if err != nil {
		log.Printf("Error ao descrever o livro")
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
	w.WriteHeader(http.StatusOK)
}

func ListBook(w http.ResponseWriter, r *http.Request) {

	listBook, err := ListBookService()
	if err != nil {
		log.Print("Error ao listar livros")
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(listBook)
	w.WriteHeader(http.StatusOK)
}

// Service
func CreateBookService(infoBook InfoBook) (Book, error) {
	book := Book{
		Id:            "01",
		Nome:          infoBook.Nome,
		Autor:         infoBook.Autor,
		Edicao:        infoBook.Edicao,
		AnoLancamento: infoBook.AnoLancamento,
	}
	return book, nil
}

func UpdateBookService(id string, infoBook InfoBook) (Book, error) {
	book := Book{
		Id:            id,
		Nome:          infoBook.Nome,
		Autor:         infoBook.Autor,
		Edicao:        infoBook.Edicao,
		AnoLancamento: infoBook.AnoLancamento,
	}
	return book, nil
}

func DeleteBookService(id string) error {
	fmt.Println("Livro deletado com sucesso", id)
	return nil
}

func DescribeBookService(id string) (Book, error) {
	book := Book{
		Id:            id,
		Nome:          "O Diário de Anne Frank",
		Autor:         "JJ",
		Edicao:        "1990.3",
		AnoLancamento: 2000,
	}
	return book, nil
}

func ListBookService() ([]Book, error) {
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
	return listBook, nil
}


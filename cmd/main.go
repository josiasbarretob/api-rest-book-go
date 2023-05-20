package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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

// Controller
func CreateBook(w http.ResponseWriter, r *http.Request) {
	infoBook := InfoBook{}
	err := json.NewDecoder(r.Body).Decode(&infoBook)
	if err != nil {
		log.Printf("Error ao Cadastrar livro %v", err)
	}

	book, err := CreateBookService(infoBook)
	if err != nil {
		log.Printf("Error ao criar livro")

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("json", "content-type")
		return
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

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
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

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func DescribeBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	book, err := DescribeBookService(id)
	if err != nil {
		log.Printf("Error ao descrever o livro")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
	w.WriteHeader(http.StatusOK)
}

func ListBook(w http.ResponseWriter, r *http.Request) {

	listBook, err := ListBookService()
	if err != nil {
		log.Print("Error ao listar livros")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(listBook)
	w.WriteHeader(http.StatusOK)
}

// Service
func CreateBookService(infoBook InfoBook) (*Book, error) {
	book, err := CreateBookRepository(infoBook)
	return book, err
}

func UpdateBookService(id string, infoBook InfoBook) (*Book, error) {
	book, err := UpdateBookRepository(id, infoBook)
	return book, err
}

func DeleteBookService(id string) error {
	err := DeleteBookRepository(id)
	return err
}

func DescribeBookService(id string) (Book, error) {
	book, err := DescribeBookRepository(id)
	return book, err
}

func ListBookService() ([]Book, error) {
	listBook, err := ListBookRepository()
	return listBook, err
}

// Repository
var bookMap = make(map[string]Book)

func CreateBookRepository(infoBook InfoBook) (*Book, error) {
	numberRandon := rand.Intn(10)
	id := strconv.Itoa(numberRandon)
	if _, ok := bookMap[id]; ok {
		return nil, fmt.Errorf("livro já existe")
	}

	book := Book{
		Id:            id,
		Nome:          infoBook.Nome,
		Autor:         infoBook.Autor,
		Edicao:        infoBook.Edicao,
		AnoLancamento: infoBook.AnoLancamento,
	}

	bookMap[id] = book
	return &book, nil
}

func UpdateBookRepository(id string, infoBook InfoBook) (*Book, error) {
	if _, ok := bookMap[id]; ok {
		book := Book{
			Id:            id,
			Nome:          infoBook.Nome,
			Autor:         infoBook.Autor,
			Edicao:        infoBook.Edicao,
			AnoLancamento: infoBook.AnoLancamento,
		}
		bookMap[id] = book
		return &book, nil
	}
	return nil, fmt.Errorf("ID não encontrado")
}

func DeleteBookRepository(id string) error {
	if _, ok := bookMap[id]; ok {
		delete(bookMap, id)
		return nil
	}
	return fmt.Errorf("ID não encontrado")
}

func DescribeBookRepository(id string) (Book, error) {
	if _, ok := bookMap[id]; ok {
		return bookMap[id], nil
	}
	return bookMap[id], fmt.Errorf("ID não encontrado")
}

func ListBookRepository() ([]Book, error) {
	listBook := convertMapToSlice(bookMap)
	return listBook, nil
}

func convertMapToSlice(m map[string]Book) []Book {
	var listBook []Book
	for _, v := range m {
		listBook = append(listBook, v)
	}
	return listBook
}

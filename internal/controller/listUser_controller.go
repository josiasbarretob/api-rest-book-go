package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/service"
)

func ListBook(w http.ResponseWriter, r *http.Request) {

	listBook, err := service.ListBookService()
	if err != nil {
		log.Print("Error ao listar livros")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(listBook)
	w.WriteHeader(http.StatusOK)
}
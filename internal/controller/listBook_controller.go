package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/service"
)

type listBookController struct {
	listService service.ListServiceInterface
}

func NewListBookController(listService service.ListServiceInterface)*listBookController{
	return &listBookController{
		listService: listService,
	}
}
func (l listBookController) ListBook(w http.ResponseWriter, r *http.Request) {

	listBook, err := l.listService.ListBookService()
	if err != nil {
		log.Print("Error ao listar livros")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(listBook)
	w.WriteHeader(http.StatusOK)
}

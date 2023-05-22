package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/domain"
	"github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/service"
)

type updateBookController struct {
	updateService service.UpdateServiceInterface
}

func NewUpdateBookController(updateService service.UpdateServiceInterface)*updateBookController{
	return &updateBookController{updateService: updateService}
}

func (u updateBookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	infoBook := domain.InfoBook{}
	id := chi.URLParam(r, "id")
	err := json.NewDecoder(r.Body).Decode(&infoBook)
	if err != nil {
		log.Printf("Error ao atualizar livro %v", err)

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("json", "content-type")
		return
	}
	book, err := u.updateService.UpdateBookService(id, infoBook)
	if err != nil {
		log.Printf("Error ao atualizar livro")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
	w.WriteHeader(http.StatusOK)
}

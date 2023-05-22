package controller

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/service"
)

type deleteBookController struct {
	deleteService service.DeleteServiceInterface
}

func NewDeleteBookController(deleteService service.DeleteServiceInterface) *deleteBookController {
	return &deleteBookController{deleteService}
}

func (d deleteBookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := d.deleteService.DeleteBookService(id)
	if err != nil {
		log.Printf("Error ao deletar livro")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

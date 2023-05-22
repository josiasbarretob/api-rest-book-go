package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/service"
)

type describeBookController struct{
	describeService service.DescribeServiceInterface
}

func NewDescribeBookController(describeService service.DescribeServiceInterface)*describeBookController{
	return &describeBookController{
		describeService: describeService,
	}
}

func (d describeBookController)DescribeBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	book, err := d.describeService.DescribeBookService(id)
	if err != nil {
		log.Printf("Error ao descrever o livro")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
	w.WriteHeader(http.StatusOK)
}

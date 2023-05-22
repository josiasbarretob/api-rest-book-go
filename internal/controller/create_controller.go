package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/domain"
	"github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/service"
)

type createBookController struct{
	createService service.CreateServiceInterface
}

func NewCreateBookController(createService service.CreateServiceInterface)*createBookController{
	return &createBookController{
		createService: createService,
	}
}

func (c createBookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	infoBook := domain.InfoBook{}
	err := json.NewDecoder(r.Body).Decode(&infoBook)
	if err != nil {
		log.Printf("Error ao Cadastrar livro %v", err)
	}

	book, err := c.createService.CreateBookService(infoBook)
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

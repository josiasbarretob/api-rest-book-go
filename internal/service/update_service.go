package service

import (
	"github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/domain"
	"github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/repository"
)

func UpdateBookService(id string, infoBook domain.InfoBook) (*domain.Book, error) {
	book, err := repository.UpdateBookRepository(id, infoBook)
	return book, err
}

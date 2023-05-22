package service

import (
	"github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/domain"
	"github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/repository"
)

// Service
func CreateBookService(infoBook domain.InfoBook) (*domain.Book, error) {
	book, err := repository.CreateBookRepository(infoBook)
	return book, err
}
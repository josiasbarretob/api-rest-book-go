package service

import (
	"github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/domain"
	"github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/repository"
)

func ListBookService() ([]domain.Book, error) {
	listBook, err := repository.ListBookRepository()
	return listBook, err
}

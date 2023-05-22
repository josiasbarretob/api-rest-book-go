package service

import (
	"github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/domain"
	"github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/repository"
)

func DescribeBookService(id string) (domain.Book, error) {
	book, err := repository.DescribeBookRepository(id)
	return book, err
}

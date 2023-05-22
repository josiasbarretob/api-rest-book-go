package repository

import "github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/domain"

type ListBookRepository interface{
	List() ([]domain.Book, error)
}
package repository

import "github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/domain"

type CreateBookRepository interface{
	Create(infoBook domain.InfoBook) (*domain.Book, error)
}
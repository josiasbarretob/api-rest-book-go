package repository

import "github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/domain"

type UpdateBookRepository interface {
	Update(id string, infoBook domain.InfoBook) (*domain.Book, error)
}

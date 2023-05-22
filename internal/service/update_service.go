package service

import "github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/domain"

type UpdateServiceInterface interface {
	UpdateBookService(id string, infoBook domain.InfoBook) (*domain.Book, error)
}

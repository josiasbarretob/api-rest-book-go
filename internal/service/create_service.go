package service

import "github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/domain"

type CreateServiceInterface interface {
	CreateBookService(infoBook domain.InfoBook) (*domain.Book, error)
}

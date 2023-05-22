package service

import "github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/domain"

type ListServiceInterface interface{
	ListBookService()([]domain.Book, error)
}
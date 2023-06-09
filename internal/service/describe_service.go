package service

import "github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/domain"

type DescribeServiceInterface interface{
	DescribeBookService(id string) (domain.Book, error)
}
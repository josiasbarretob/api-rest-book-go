package repository

import "github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/domain"

type DescribeBookRepository interface{
	Describe(id string) (domain.Book, error)
}


package repository

import (
	"fmt"

	"github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/domain"
)

func DescribeBookRepository(id string) (domain.Book, error) {
	if _, ok := bookMap[id]; ok {
		return bookMap[id], nil
	}
	return bookMap[id], fmt.Errorf("ID não encontrado")
}

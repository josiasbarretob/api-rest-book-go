package repository

import (
	"fmt"

	"github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/domain"
)

func UpdateBookRepository(id string, infoBook domain.InfoBook) (*domain.Book, error) {
	if _, ok := bookMap[id]; ok {
		book := domain.Book{
			Id:            id,
			Nome:          infoBook.Nome,
			Autor:         infoBook.Autor,
			Edicao:        infoBook.Edicao,
			AnoLancamento: infoBook.AnoLancamento,
		}
		bookMap[id] = book
		return &book, nil
	}
	return nil, fmt.Errorf("ID não encontrado")
}

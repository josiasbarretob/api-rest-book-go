package repository

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/domain"
)

// Repository
var bookMap = make(map[string]domain.Book)

func CreateBookRepository(infoBook domain.InfoBook) (*domain.Book, error) {
	numberRandon := rand.Intn(10)
	id := strconv.Itoa(numberRandon)
	if _, ok := bookMap[id]; ok {
		return nil, fmt.Errorf("livro já existe")
	}

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

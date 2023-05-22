package memory

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/domain"
)

type memoryMap struct {
	bookMap map[string]domain.Book
}

func NewMemoryMapRepository() memoryMap {
	var bookMap = make(map[string]domain.Book)
	memory := memoryMap{bookMap: bookMap}
	return memory
}

func (m memoryMap) Create(infoBook domain.InfoBook) (*domain.Book, error) {
	numberRandon := rand.Intn(10)
	id := strconv.Itoa(numberRandon)
	if _, ok := m.bookMap[id]; ok {
		return nil, fmt.Errorf("livro já existe")
	}

	book := domain.Book{
		Id:            id,
		Nome:          infoBook.Nome,
		Autor:         infoBook.Autor,
		Edicao:        infoBook.Edicao,
		AnoLancamento: infoBook.AnoLancamento,
	}

	m.bookMap[id] = book
	return &book, nil
}

func (m memoryMap) Update(id string, infoBook domain.InfoBook) (*domain.Book, error) {
	if _, ok := m.bookMap[id]; ok {
		book := domain.Book{
			Id:            id,
			Nome:          infoBook.Nome,
			Autor:         infoBook.Autor,
			Edicao:        infoBook.Edicao,
			AnoLancamento: infoBook.AnoLancamento,
		}
		m.bookMap[id] = book
		return &book, nil
	}
	return nil, fmt.Errorf("ID não encontrado")
}

func (m memoryMap) Describe(id string) (domain.Book, error) {
	if _, ok := m.bookMap[id]; ok {
		return m.bookMap[id], nil
	}
	return m.bookMap[id], fmt.Errorf("ID não encontrado")
}

func (m memoryMap) Delete(id string) error {
	if _, ok := m.bookMap[id]; ok {
		delete(m.bookMap, id)
		return nil
	}
	return fmt.Errorf("ID não encontrado")
}

func (m memoryMap) List() ([]domain.Book, error) {
	listBook := convertMapToSlice(m.bookMap)
	return listBook, nil
}

func convertMapToSlice(m map[string]domain.Book) []domain.Book {
	var listBook []domain.Book
	for _, v := range m {
		listBook = append(listBook, v)
	}
	return listBook
}

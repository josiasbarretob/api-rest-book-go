package repository

import "github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/domain"

func ListBookRepository() ([]domain.Book, error) {
	listBook := convertMapToSlice(bookMap)
	return listBook, nil
}

func convertMapToSlice(m map[string]domain.Book) []domain.Book {
	var listBook []domain.Book
	for _, v := range m {
		listBook = append(listBook, v)
	}
	return listBook
}

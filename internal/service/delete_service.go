package service

import "github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/repository"

func DeleteBookService(id string) error {
	err := repository.DeleteBookRepository(id)
	return err
}

package repository

import "fmt"

func DeleteBookRepository(id string) error {
	if _, ok := bookMap[id]; ok {
		delete(bookMap, id)
		return nil
	}
	return fmt.Errorf("ID não encontrado")
}

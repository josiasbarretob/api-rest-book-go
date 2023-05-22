package repository

type DeleteBookRepository interface {
	Delete(id string) error
}

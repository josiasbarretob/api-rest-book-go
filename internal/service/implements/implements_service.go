package implements

import (
	"github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/domain"
	"github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/repository"
)

type bookService struct {
	createBookRepository   repository.CreateBookRepository
	updateBookRepository   repository.UpdateBookRepository
	deleteBookRepository   repository.DeleteBookRepository
	describeBookRepository repository.DescribeBookRepository
	listBookRepository     repository.ListBookRepository
}

// Construtor
func NewBookService(
	createBookRepo repository.CreateBookRepository,
	updateBookRepo repository.UpdateBookRepository,
	deleteBookRepo repository.DeleteBookRepository,
	describeBookRepo repository.DescribeBookRepository,
	listBookRepo repository.ListBookRepository,
) *bookService {
	return &bookService{
		createBookRepository:   createBookRepo,
		updateBookRepository:   updateBookRepo,
		deleteBookRepository:   deleteBookRepo,
		describeBookRepository: describeBookRepo,
		listBookRepository:     listBookRepo,
	}
}

func (b bookService) CreateBookService(infoBook domain.InfoBook) (*domain.Book, error) {
	book, err := b.createBookRepository.Create(infoBook)
	return book, err
}

func (b bookService) UpdateBookService(id string, infoBook domain.InfoBook) (*domain.Book, error) {
	book, err := b.updateBookRepository.Update(id, infoBook)
	return book, err
}

func (b bookService) DeleteBookService(id string) error {
	err := b.deleteBookRepository.Delete(id)
	return err
}

func (b bookService) DescribeBookService(id string) (domain.Book, error) {
	book, err := b.describeBookRepository.Describe(id)
	return book, err
}

func (b bookService) ListBookService() ([]domain.Book, error) {
	listBook, err := b.listBookRepository.List()
	return listBook, err
}

package service

type DeleteServiceInterface interface{
	DeleteBookService(id string) error
}
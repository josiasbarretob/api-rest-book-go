package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/controller"
	"github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/repository/memory"
	"github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/service/implements"
)

var repository = memory.NewMemoryMapRepository()
var implementsBookService = implements.NewBookService(repository, repository, repository, repository, repository)

var (
	createController   = controller.NewCreateBookController(implementsBookService)
	updateController   = controller.NewUpdateBookController(implementsBookService)
	deleteController   = controller.NewDeleteBookController(implementsBookService)
	describeController = controller.NewDescribeBookController(implementsBookService)
	listController     = controller.NewListBookController(implementsBookService)
)

func main() {
	router := chi.NewRouter()
	log.Println("Loading in server...")

	router.Post("/book", createController.CreateBook)
	router.Put("/book/{id}", updateController.UpdateBook)
	router.Delete("/book/{id}", deleteController.DeleteBook)
	router.Get("/book/{id}", describeController.DescribeBook)
	router.Get("/books", listController.ListBook)

	log.Fatal(http.ListenAndServe(":8080", router))
}

package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mercadolibre/api-rest-book-golang/api-rest-book-go/internal/controller"
)

// Rotas
func main() {
	router := chi.NewRouter()
	log.Println("Loading in server...")
	router.Post("/book", controller.CreateBook)
	router.Put("/book/{id}", controller.UpdateBook)
	router.Delete("/book/{id}", controller.DeleteBook)
	router.Get("/book/{id}", controller.DescribeBook)
	router.Get("/books", controller.ListBook)

	log.Fatal(http.ListenAndServe(":8080", router))
}
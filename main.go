package main

import (
	"net/http"

	"golang_crud/controller"
	"golang_crud/db"
	"golang_crud/repository"
	"golang_crud/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	dbConn := db.NewDB()
	defer dbConn.Close()

	validate := *validator.New()

	bookRepository := repository.NewBookRepository(dbConn)
	bookService := service.NewBookService(bookRepository, &validate)
	bookController := controller.NewBookController(bookService)

	router := httprouter.New()
	router.POST("/api/v1/books", bookController.Save)
	router.GET("/api/v1/books", bookController.FindAll)
	router.GET("/api/v1/books/:idBook", bookController.FindById)
	router.DELETE("/api/v1/books/:idBook", bookController.Delete)
	router.PUT("/api/v1/books/:idBook", bookController.Update)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}

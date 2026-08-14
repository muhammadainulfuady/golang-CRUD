package main

import (
	"net/http"

	"golang_crud/db"
	"golang_crud/handler"

	"github.com/julienschmidt/httprouter"
)

func main() {
	dbConn := db.NewDB()
	defer dbConn.Close()

	bookHandler := handler.Book{
		DB: dbConn,
	}

	router := httprouter.New()

	router.POST("/api/v1/book/", bookHandler.Create)
	serve := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	serve.ListenAndServe()
}

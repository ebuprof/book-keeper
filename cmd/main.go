package main

import (
	"log"
	"net/http"

	"github.com/book_keeper_go/pkg/db"
	"github.com/book_keeper_go/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	//
	DB := db.Init()
	h := handlers.New(DB)

	router := mux.NewRouter()

	router.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", h.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books", h.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", h.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", h.DeleteBook).Methods(http.MethodDelete)
	router.HandleFunc("/users", h.AddNewPerson).Methods(http.MethodPost)
	router.HandleFunc("/users{id}", h.UpdatePerson).Methods(http.MethodPut)

	log.Println("API is running")
	http.ListenAndServe(":9080", router)
}

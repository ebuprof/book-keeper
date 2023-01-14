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
	router.HandleFunc("/people", h.AddNewPerson).Methods(http.MethodPost)
	router.HandleFunc("/people/{id}", h.UpdatePerson).Methods(http.MethodPut)
	router.HandleFunc("/people", h.GetPeople).Methods(http.MethodGet)
	router.HandleFunc("/people/{id}", h.GetPerson).Methods(http.MethodGet)
	router.HandleFunc("/register", h.Register).Methods(http.MethodPost)
	router.HandleFunc("/login", h.Login).Methods(http.MethodPost)

	log.Println("API is running")
	http.ListenAndServe(":9080", router)
}

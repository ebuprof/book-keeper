package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/book_keeper_go/pkg/models"
)

func (h handler) AddBook(w http.ResponseWriter, r *http.Request) {
	// read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	// save book
	if result := h.DB.Create(&book); result.Error != nil {
		log.Fatal(result.Error)
	}

	// send response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("New book created")
}

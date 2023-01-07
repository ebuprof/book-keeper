package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/book_keeper_go/pkg/models"
)

func (h handler) AddNewPerson(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var person models.Person
	json.Unmarshal(body, &person)

	if result := h.DB.Create(&person); result.Error != nil {
		log.Fatal(result.Error)
	}

	//response to client
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("User Created")
}

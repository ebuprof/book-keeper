package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/book_keeper_go/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var updatedPerson models.Person
	json.Unmarshal(body, &updatedPerson)

	var person models.Person

	if result := h.DB.First(&person, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	// map data update from request
	person.Email = updatedPerson.Email
	person.Model = updatedPerson.Model
	person.Name = updatedPerson.Name

	h.DB.Save(&person)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Person Updated")
}

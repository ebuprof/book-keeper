package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/book_keeper_go/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) GetPerson(w http.ResponseWriter, r *http.Request) {
	// get the id value from request
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var person models.Person

	// fetch by id
	if result := h.DB.First(&person, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "apolication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(person)
}

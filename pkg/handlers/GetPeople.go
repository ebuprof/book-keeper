package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/book_keeper_go/pkg/models"
)

func (h handler) GetPeople(w http.ResponseWriter, r *http.Request) {
	var people models.Person

	// fetch all persons

	if result := h.DB.Find(&people); result != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(people)
}

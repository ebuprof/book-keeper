package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/book_keeper_go/pkg/models"
	"github.com/book_keeper_go/pkg/util"
)

func (h handler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// hash password
	hashPassword, err := util.HashPassword(user.Password)
	//util.CheckError(err)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	user.Password = hashPassword

	if result := h.DB.Create(&user); result.Error != nil {
		http.Error(w, "Database operation error", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("User registered successfully")
}

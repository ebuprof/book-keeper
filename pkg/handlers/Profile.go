package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/book_keeper_go/pkg/models"
)

func (h handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		return
	}

	username := r.Context().Value("username").(string)

	var user models.User
	if err := h.DB.Where("email = ?", username).First(&user).Error; err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

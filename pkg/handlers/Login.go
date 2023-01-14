package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/book_keeper_go/pkg/dto"
	"github.com/book_keeper_go/pkg/models"
	"github.com/book_keeper_go/pkg/util"
)

func (h handler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var login dto.LoginDto
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := h.DB.Where("email = ?", login.Email).First(&user).Error; err != nil {
		http.Error(w, "invalid email or password", http.StatusUnauthorized)
		return
	}

	if err := util.ValidatePassword(login.Password, user.Password); err != nil {
		http.Error(w, "invalid email or password", http.StatusUnauthorized)
		return
	}

	tokenString, err := util.CreateToken(login.Email, user.ID.String())
	if err != nil {
		http.Error(w, "Error signing token", http.StatusInternalServerError)
		return
	}

	//w.Write([]byte(tokenString))
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode([]byte(tokenString))

}

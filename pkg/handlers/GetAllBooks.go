package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/book_keeper_go/pkg/dto"
	"github.com/book_keeper_go/pkg/models"
)

func (h handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {

	var books []models.Book
	var person models.Person
	var pResponses []dto.PersonResponse
	var pResponse dto.PersonResponse

	if result := h.DB.Find(&books); result.Error != nil {
		fmt.Println(result.Error)
	}

	// iterate and get every person linked to a book from list of books, and map to the response object
	for _, book := range books {
		if res := h.DB.First(&person, book.PersonID); res.Error != nil {
			fmt.Println(res.Error)
		}
		pResponse.Author = book.Author
		pResponse.Title = book.Title
		pResponse.Description = book.Desc
		pResponse.PhoneNumber = book.PhoneNumber
		pResponse.PersonName = person.Name
		pResponse.PersonEmail = person.Email

		pResponses = append(pResponses, pResponse)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pResponses)
}

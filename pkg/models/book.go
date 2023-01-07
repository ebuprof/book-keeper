package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model

	Title       string `json:"title"`
	Author      string `json:"author"`
	PhoneNumber string `gorm:"unique_index"`
	Desc        string `json:"desc"`
	PersonID    int
}

/*var BooksData = []Book{
	{Title: "Understanding Nigeria's Law", Author: "Ebuka", PhoneNumber: "08038785674", Desc: "Understanding Nigeria's Law", PersonID: 1},
	{Title: "Philosophy of Life", Author: "Ebuka", PhoneNumber: "08038785675", Desc: "Philosophy of Life", PersonID: 1},
}**/

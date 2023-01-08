package db

import (
	"fmt"
	"log"
	"os"

	"github.com/book_keeper_go/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	person = &models.Person{Name: "Chukwuebuka", Email: "ebukaprof@gmail.com"}
	book   = &models.Book{Title: "Understanding Nigeria's Law", Author: "Ebuka", PhoneNumber: "08038785674", Desc: "Understanding Nigeria's Law", PersonID: 1}
)

func Init() *gorm.DB {
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	dbPassword := os.Getenv("PASSWORD")

	dbUrl := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, dbPassword, dbPort)
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Successfully connected to database")
	}

	db.AutoMigrate(&models.Book{})
	db.AutoMigrate((&models.Person{}))

	// seed in data if not yet in the db
	///db.Create(&models.PersonData)
	///db.Create(&book)

	return db
}

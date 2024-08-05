package models

import (
	"github.com/divan009/Multiple-Golang-Projects/pkg/config"
	//_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"column:name" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// CreateBook is a method on type Book - encapsulation
func (b *Book) CreateBook() *Book {
	result := db.Create(b)
	//we get values of type book in b
	if result.Error != nil {
		panic(result.Error) // or handle the error as you deem appropriate
	}
	// we return type book
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book

	// Use a pointer to `book` in the Delete method call
	db.Where("ID=?", ID).Delete(&book)

	return book
}

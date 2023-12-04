package models

import (
	"github.com/jinzhu/gorm"
	"github.com/rafsnil/CRUD-WITH-MySQL/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// Init function is executed as soon as the package is imported
func init() {
	config.Connect()
	db = config.GetDB()
	/* Address is passed here to allow the automigrate func to examine
	the struct's fields and create the corresponding table columns in the database.*/
	db.AutoMigrate(&Book{})
}

// Create a Book
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// Get all books
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

// Get Book by ID
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

// Delete a book
func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	utils "github.com/rafsnil/CRUD-WITH-MySQL/pkg/Utils"
	"github.com/rafsnil/CRUD-WITH-MySQL/pkg/models"
)

var NewBook models.Book

// GET BOOK HANDLER
func GetBook(w http.ResponseWriter, r *http.Request) {
	//Getting all books from the db
	newBooks := models.GetAllBooks()
	//Converting the data to json
	res, _ := json.Marshal(newBooks)
	//Letting the client side know about the data format (JSON in this case)
	w.Header().Set("Content-Type", "application/json")
	//Showing green light for execution
	w.WriteHeader(http.StatusOK)
	//Writing response to the client
	w.Write(res)
}

// GET BOOK BY ID HANDLER
func GetBookById(w http.ResponseWriter, r *http.Request) {
	//extracting parameters from URL in a map
	vars := mux.Vars(r)
	//Getting my desired parameter
	bookId := vars["bookId"]
	//Parsing the string to int
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Parsing Error while getting book by id!!!")
	}
	//Looking for the book in the DB
	requiredBook, _ := models.GetBookById(Id)
	//Converting the data found in DB to Json
	res, _ := json.Marshal(requiredBook)
	// w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// CREATE BOOK HANDLER
func CreateNewBook(w http.ResponseWriter, r *http.Request) {
	//Giving the structure of the Book in models to newBook
	newBook := &models.Book{}
	//Converting the given data (json) to GO data (struct, map, slice)
	utils.ParseBody(r, newBook)
	//Creating a book using the CreateBook func in the models
	b := newBook.CreateBook()
	//Converting the data back to json
	//This also automatically sets "Content-Type" to "application/json"
	res, _ := json.Marshal(b)
	// w.Header().Set("Content-Type", "application/json")
	//Green Flag ✔
	w.WriteHeader(http.StatusOK)
	//Writing back the response to the client
	w.Write(res)
}

// DELETE BOOK HANDLER
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Parsing Error while deleting!!!")
	}
	toBeDeletedBook := models.DeleteBook(id)
	res, _ := json.Marshal(toBeDeletedBook)
	// w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// UPDATE BOOK HANDLER
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var bookTobeUpdated = &models.Book{}
	utils.ParseBody(r, bookTobeUpdated)

	params := mux.Vars(r)
	bookId := params["bookId"]

	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Parsing Error while updating!!!")
	}

	bookBeforeUpdate, db := models.GetBookById(id)
	if bookTobeUpdated.Name != "" {
		bookBeforeUpdate.Name = bookTobeUpdated.Name
	}
	if bookTobeUpdated.Author != "" {
		bookBeforeUpdate.Author = bookTobeUpdated.Author
	}
	if bookTobeUpdated.Publication != "" {
		bookBeforeUpdate.Publication = bookTobeUpdated.Publication
	}

	db.Save(&bookBeforeUpdate)

	res, _ := json.Marshal(bookBeforeUpdate)

	// w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

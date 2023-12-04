package routes

import (
	"github.com/gorilla/mux"
	"github.com/rafsnil/CRUD-WITH-MySQL/pkg/Controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", Controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", Controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", Controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", Controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", Controllers.DeleteBook).Methods("DELETE")
}

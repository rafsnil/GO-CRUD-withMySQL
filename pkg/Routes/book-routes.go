package routes

import (
	"github.com/gorilla/mux"
	controllers "github.com/rafsnil/CRUD-WITH-MySQL/pkg/Controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.CreateNewBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}

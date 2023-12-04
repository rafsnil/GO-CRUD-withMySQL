package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	// "github.com/jinzhu/gorm/dialects/mysql"
	routes "github.com/rafsnil/CRUD-WITH-MySQL/pkg/Routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}

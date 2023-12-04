package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
The whole point of this config file
is to return a variable called DB which
will allow other files to interact with the database
*/
var (
	db *gorm.DB
)

// The Connect functions helps to connect with our Database
// Described properly in the Notes at no. 3
func Connect() {

	//TO DO: mysql credentials
	d, err := gorm.Open("mysql", "")
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}

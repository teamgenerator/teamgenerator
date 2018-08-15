////
// author: Nico Alimin (nicoalimin@hotmail.com)
// date: Tuesday, 14th August 2018 1:12:38 am
// lastModifiedBy: Nico Alimin (nicoalimin@hotmail.com)
// lastModifiedTime: Tuesday, 14th August 2018 1:12:38 am
//
// DESCRIPTION
//
// copyright (c) 2018 Nico Alimin
////

package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// DB is the database we want to export
var DB *gorm.DB

// Open opens the database connection
func Open(pgUser string, pgPassword string, pgDatabase string, pgHost string, pgPort string) {

	// Initializes the connection string.
	connectionString :=
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", pgHost, pgPort, pgUser, pgPassword, pgDatabase)

	// Initiates connection to the postgres database
	var err error
	DB, err = gorm.Open("postgres", connectionString)
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Printf("Opened connection to %s\n", connectionString)
}

// Close closes the database connection
func Close() error {
	return DB.Close()
}

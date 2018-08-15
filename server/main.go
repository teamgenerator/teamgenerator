////
// author: Nico Alimin (nicoalimin@gmail.com)
// date: Sunday, 5th August 2018 12:48:35 am
// lastModifiedBy: Nico Alimin (nicoalimin@gmail.com)
// lastModifiedTime: Sunday, 5th August 2018 12:52:42 am
//
// DESCRIPTION
//
// copyright (c) 2018 Nico Alimin
////

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/teamgenerator/teamgenerator/server/db"
	"github.com/teamgenerator/teamgenerator/server/models"
)

var err error

var (
	pgUser     = os.Getenv("PG_USER")
	pgPassword = os.Getenv("PG_PASSWORD")
	pgDatabase = os.Getenv("PG_DATABASE")
	pgHost     = os.Getenv("PG_HOST")
	pgPort     = os.Getenv("PG_PORT")
	port       = ":3030"
)

func main() {
	main := mux.NewRouter()
	routerAPI := main.PathPrefix("/api").Subrouter()
	router := routerAPI.PathPrefix("/v1").Subrouter()

	// Initiate connection to postgres database
	db.Open(pgUser, pgPassword, pgDatabase, pgHost, pgPort)
	defer db.Close()

	// Migrates the models
	db.DB.AutoMigrate(&models.Player{}, &models.Community{})
	db.DB.Model(&models.Player{}).AddForeignKey("community_id", "communities(id)", "CASCADE", "CASCADE")

	// Route for community-related endpoints
	router.HandleFunc("/Communities", models.GetCommunities).Methods("GET")
	router.HandleFunc("/Communities/{id}", models.GetCommunity).Methods("GET")
	router.HandleFunc("/Communities", models.CreateCommunity).Methods("POST")
	router.HandleFunc("/Communities/{id}", models.UpdateCommunity).Methods("PATCH")
	router.HandleFunc("/Communities/{id}", models.DeleteCommunity).Methods("DELETE")

	// Route for player-related endpoints
	router.HandleFunc("/Players", models.GetPlayers).Methods("GET")
	router.HandleFunc("/Players/{id}", models.GetPlayer).Methods("GET")
	router.HandleFunc("/Players", models.CreatePlayer).Methods("POST")
	router.HandleFunc("/Players/{id}", models.UpdatePlayer).Methods("PATCH")
	router.HandleFunc("/Players/{id}", models.DeletePlayer).Methods("DELETE")

	fmt.Printf("Go Backend Service initialized at port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

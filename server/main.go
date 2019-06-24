package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/teamgenerator/teamgenerator/server/db"
	"github.com/teamgenerator/teamgenerator/server/models"
	"github.com/teamgenerator/teamgenerator/server/seed"
)

var err error

var (
	pgUser     = "postgres"
	pgPassword = "password"
	pgDatabase = "postgres"
	pgHost     = "localhost"
	pgPort     = "5432"
	port       = ":3030"
)

func main() {

	// Go Routers. Defaults to /api/v1
	main := mux.NewRouter()
	routerAPI := main.PathPrefix("/api").Subrouter()
	router := routerAPI.PathPrefix("/v1").Subrouter()

	// Initiate connection to postgres database
	db.Open(pgUser, pgPassword, pgDatabase, pgHost, pgPort)
	defer db.Close()

	migrateModels()
	addRelations()
	seed.Seed()
	// Route for community-related endpoints
	router.HandleFunc("/communities", models.GetCommunities).Methods("GET")
	router.HandleFunc("/communities/{id}", models.GetCommunity).Methods("GET")
	router.HandleFunc("/communities", models.CreateCommunity).Methods("POST")
	router.HandleFunc("/communities/{id}", models.UpdateCommunity).Methods("PATCH")
	router.HandleFunc("/communities/{id}", models.DeleteCommunity).Methods("DELETE")

	// Route for player-related endpoints
	router.HandleFunc("/players", models.GetPlayers).Methods("GET")
	router.HandleFunc("/players/{id}", models.GetPlayer).Methods("GET")
	router.HandleFunc("/players", models.CreatePlayer).Methods("POST")
	router.HandleFunc("/players/{id}", models.UpdatePlayer).Methods("PATCH")
	router.HandleFunc("/players/{id}", models.DeletePlayer).Methods("DELETE")

	// Route for user-related endpoints
	router.HandleFunc("/users", models.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", models.GetUser).Methods("GET")
	router.HandleFunc("/users", models.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", models.UpdateUser).Methods("PATCH")
	router.HandleFunc("/users/{id}", models.DeleteUser).Methods("DELETE")

	// Route for session-related endpoints
	router.HandleFunc("/session", models.GetSessions).Methods("GET")
	router.HandleFunc("/session/{id}", models.GetSession).Methods("GET")
	router.HandleFunc("/session", models.CreateSession).Methods("POST")
	router.HandleFunc("/session/{id}", models.UpdateSession).Methods("PATCH")
	router.HandleFunc("/session/{id}", models.DeleteSession).Methods("DELETE")

	// Route for sessionPlayer-related endpoints
	router.HandleFunc("/session-player", models.GetSessionPlayers).Methods("GET")
	router.HandleFunc("/session-player/{id}", models.GetSessionPlayer).Methods("GET")
	router.HandleFunc("/session-player", models.CreateSessionPlayer).Methods("POST")
	router.HandleFunc("/session-player/{id}", models.UpdateSessionPlayer).Methods("PATCH")
	router.HandleFunc("/session-player/{id}", models.DeleteSessionPlayer).Methods("DELETE")

	// Route for rating-related endpoints
	router.HandleFunc("/rating", models.GetRatings).Methods("GET")
	router.HandleFunc("/rating/{id}", models.GetRating).Methods("GET")
	router.HandleFunc("/rating", models.CreateRating).Methods("POST")
	router.HandleFunc("/rating/{id}", models.UpdateRating).Methods("PATCH")
	router.HandleFunc("/rating/{id}", models.DeleteRating).Methods("DELETE")

	fmt.Printf("Go Backend Service initialized at port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

// Migrate the models
func migrateModels() {
	db.DB.AutoMigrate(&models.Player{}, &models.Community{}, &models.SessionPlayer{}, &models.Session{}, &models.User{}, &models.Rating{})
	fmt.Printf("Successfully migrated models\n")
}

// Add the relations for the tables. E.g. Foreign Keys
func addRelations() {
	db.DB.Model(&models.Player{}).AddForeignKey("community_id", "communities(id)", "CASCADE", "CASCADE")
	db.DB.Model(&models.Session{}).AddForeignKey("community_id", "communities(id)", "CASCADE", "CASCADE")
	db.DB.Model(&models.SessionPlayer{}).AddForeignKey("community_id", "communities(id)", "CASCADE", "CASCADE")
	db.DB.Model(&models.SessionPlayer{}).AddForeignKey("player_id", "players(id)", "CASCADE", "CASCADE")
	db.DB.Model(&models.Rating{}).AddForeignKey("player_id", "players(id)", "CASCADE", "CASCADE")
	fmt.Printf("Successfully added foreign keys\n")
}

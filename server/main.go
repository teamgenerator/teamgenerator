package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/teamgenerator/teamgenerator/server/db"
	"github.com/teamgenerator/teamgenerator/server/pkg/core"
	"github.com/teamgenerator/teamgenerator/server/pkg/database"
	"github.com/teamgenerator/teamgenerator/server/pkg/handler"
	"github.com/teamgenerator/teamgenerator/server/pkg/models"
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

	communityRepo := database.CommunityRepo{}
	communityCore := core.CommunityCore{
		CommunityRepo: &communityRepo,
	}
	communityHandler := handler.CommunityHandler{
		CommunityCore: communityCore,
	}

	playerRepo := database.PlayerRepo{}
	playerCore := core.PlayerCore{
		PlayerRepo:    &playerRepo,
		CommunityRepo: &communityRepo,
	}
	playerHandler := handler.PlayerHandler{
		PlayerCore: playerCore,
	}

	sessionRepo := database.SessionRepo{}
	sessionCore := core.SessionCore{
		SessionRepo:   &sessionRepo,
		CommunityRepo: &communityRepo,
	}
	sessionHandler := handler.SessionHandler{
		SessionCore: sessionCore,
	}

	userRepo := database.UserRepo{}
	userCore := core.UserCore{
		UserRepo: &userRepo,
	}
	userHandler := handler.UserHandler{
		UserCore: userCore,
	}

	// Route for community-related endpoints
	router.HandleFunc("/communities", communityHandler.GetCommunities).Methods("GET")
	router.HandleFunc("/communities/{id}", communityHandler.GetCommunity).Methods("GET")
	router.HandleFunc("/communities", communityHandler.CreateCommunity).Methods("POST")
	router.HandleFunc("/communities/{id}", communityHandler.DeleteCommunity).Methods("DELETE")

	// Route for player-related endpoints
	router.HandleFunc("/players", playerHandler.GetPlayers).Methods("GET")
	router.HandleFunc("/players/{id}", playerHandler.GetPlayer).Methods("GET")
	router.HandleFunc("/players", playerHandler.CreatePlayer).Methods("POST")
	router.HandleFunc("/players/{id}", playerHandler.DeletePlayer).Methods("DELETE")

	// Route for user-related endpoints
	router.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	// Route for session-related endpoints
	router.HandleFunc("/sessions", sessionHandler.GetSessions).Methods("GET")
	router.HandleFunc("/sessions/{id}", sessionHandler.GetSession).Methods("GET")
	router.HandleFunc("/sessions", sessionHandler.CreateSession).Methods("POST")
	router.HandleFunc("/sessions/{id}", sessionHandler.UpdateSession).Methods("PATCH")
	router.HandleFunc("/sessions/{id}", sessionHandler.DeleteSession).Methods("DELETE")

	// Route for sessionPlayer-related endpoints
	router.HandleFunc("/session-players", models.GetSessionPlayers).Methods("GET")
	router.HandleFunc("/session-players/{id}", models.GetSessionPlayer).Methods("GET")
	router.HandleFunc("/session-players", models.CreateSessionPlayer).Methods("POST")
	router.HandleFunc("/session-players/{id}", models.UpdateSessionPlayer).Methods("PATCH")
	router.HandleFunc("/session-players/{id}", models.DeleteSessionPlayer).Methods("DELETE")

	// Route for rating-related endpoints
	router.HandleFunc("/ratings", models.GetRatings).Methods("GET")
	router.HandleFunc("/ratings/{id}", models.GetRating).Methods("GET")
	router.HandleFunc("/ratings", models.CreateRating).Methods("POST")
	router.HandleFunc("/ratings/{id}", models.UpdateRating).Methods("PATCH")
	router.HandleFunc("/ratings/{id}", models.DeleteRating).Methods("DELETE")

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

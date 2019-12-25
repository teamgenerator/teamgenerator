package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/teamgenerator/teamgenerator/server/conf"
	"github.com/teamgenerator/teamgenerator/server/db"
	"github.com/teamgenerator/teamgenerator/server/pkg/core"
	"github.com/teamgenerator/teamgenerator/server/pkg/database"
	"github.com/teamgenerator/teamgenerator/server/pkg/handler"
	"github.com/teamgenerator/teamgenerator/server/pkg/models"
	"github.com/teamgenerator/teamgenerator/server/seed"
)

var (
	err           error
	config        conf.Conf
	globalConfigs = conf.NewConf()
)

func main() {

	// Go Routers. Defaults to /api/v1
	main := mux.NewRouter()
	routerAPI := main.PathPrefix("/api").Subrouter()
	router := routerAPI.PathPrefix("/v1").Subrouter()

	// Initiate connection to postgres database
	db.Open(globalConfigs.PgUser, globalConfigs.PgPassword, globalConfigs.PgDatabase, globalConfigs.PgHost, globalConfigs.PgPort)
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

	ratingRepo := database.RatingRepo{}
	ratingCore := core.RatingCore{
		RatingRepo: &ratingRepo,
		PlayerRepo: &playerRepo,
	}
	ratingHandler := handler.RatingHandler{
		RatingCore: ratingCore,
	}

	sessionPlayerRepo := database.SessionPlayerRepo{}
	sessionPlayerCore := core.SessionPlayerCore{
		SessionPlayerRepo: &sessionPlayerRepo,
		CommunityRepo:     &communityRepo,
		PlayerRepo:        &playerRepo,
	}
	sessionPlayerHandler := handler.SessionPlayerHandler{
		SessionPlayerCore: sessionPlayerCore,
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
	router.HandleFunc("/session-players", sessionPlayerHandler.GetSessionPlayers).Methods("GET")
	router.HandleFunc("/session-players/{id}", sessionPlayerHandler.GetSessionPlayer).Methods("GET")
	router.HandleFunc("/session-players", sessionPlayerHandler.CreateSessionPlayer).Methods("POST")
	router.HandleFunc("/session-players/{id}", sessionPlayerHandler.DeleteSessionPlayer).Methods("DELETE")

	// Route for rating-related endpoints
	router.HandleFunc("/ratings", ratingHandler.GetRatings).Methods("GET")
	router.HandleFunc("/ratings/{id}", ratingHandler.GetRating).Methods("GET")
	router.HandleFunc("/ratings", ratingHandler.CreateRating).Methods("POST")
	router.HandleFunc("/ratings/{id}", ratingHandler.DeleteRating).Methods("DELETE")

	fmt.Printf("Go Backend Service initialized at port %s\n", globalConfigs.ApiPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", globalConfigs.ApiPort), router))
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

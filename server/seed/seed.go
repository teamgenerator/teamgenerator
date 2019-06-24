package seed

import (
	"fmt"

	"github.com/teamgenerator/teamgenerator/server/db"

	"github.com/teamgenerator/teamgenerator/server/models"
)

// Seed is to seed data to the database on startup
func Seed() {
	seedCommunityTSSSaturdays()
	seedPlayerNico()
	seedPlayerEnglo()
}

func seedCommunityTSSSaturdays() {
	community1 := models.Community{Name: "TSSSaturdays", Location: "Richmond"}

	var community models.Community
	if db.DB.Where("name = ?", community1.Name).First(&community).RecordNotFound() {
		db.DB.Create(&community1)
		fmt.Printf("Seeded TSSSaturdays to Community table\n")
	}
}

func seedPlayerNico() {
	var player models.Player

	var playerNico = models.Player{Name: "Nico Alimin", Ratings: 9, Form: 0, CommunityID: 1}
	if db.DB.Where("name = ?", playerNico.Name).First(&player).RecordNotFound() {
		db.DB.Create(&playerNico)
		fmt.Printf("Seeded Nico to Players table\n")
	}
}

func seedPlayerEnglo() {
	var player models.Player

	var playerEnglo = models.Player{Name: "Michael Englo", Ratings: 2, Form: 0, CommunityID: 1}
	if db.DB.Where("name = ?", playerEnglo.Name).First(&player).RecordNotFound() {
		db.DB.Create(&playerEnglo)
		fmt.Printf("Seeded Englo to Players table\n")
	}
}

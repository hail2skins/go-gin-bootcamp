package helpers

import (
	"gin_notes/database"
	"gin_notes/models"
)

func LoadDatabase() {
	database.Connect()
	//database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&models.Note{})
}

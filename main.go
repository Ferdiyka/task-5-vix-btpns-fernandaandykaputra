package main

import (
	"github.com/Ferdiyka/task-5-vix-btpns-fernandaandykaputra/database"
	"github.com/Ferdiyka/task-5-vix-btpns-fernandaandykaputra/router"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = database.DatabaseConnection()
)

func main() {
	defer database.CloseDatabaseConnection(db)

	router.AuthRouter()
	router.PhotoRouter()
	router.UserRouter()
}
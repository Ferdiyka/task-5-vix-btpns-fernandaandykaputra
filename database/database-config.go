package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/Ferdiyka/task-5-vix-btpns-fernandaandykaputra/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// setup database connection
func DatabaseConnection() *gorm.DB {
	if godotenv.Load() != nil {
		panic("Gagal meload env file")
	}

	//cofiguration db on user pass host and name
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dataConn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&local=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dataConn), &gorm.Config{})

	if err != nil {
		panic("Gagal menghubungkan ke database")
	}
	db.AutoMigrate(&models.Photo{}, &models.User{})
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dataSQL, err := db.DB()
	if err != nil {
		panic("Gagal memutuskan koneksi ke database")
	}
	dataSQL.Close()
}
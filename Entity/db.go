package Entity

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbpassword := os.Getenv("DB_PASSWORD")
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")

	config := fmt.Sprintf("", user, dbName, dbpassword, dbhost, dbport)
	db, err := gorm.Open(mysql.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connecting to database")
	}
	return db

}

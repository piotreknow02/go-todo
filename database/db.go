package database

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDatabase() (*gorm.DB, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		// If .env file is not found in cwd, load from the directory above (needed for tests)
		godotenv.Load("../.env")
		// If .env file is not found also in the directory above, ignore it
	}

	// Data needed for connection string
	dbhost := os.Getenv("DB_HOST")
	dbname := "gotodo"
	dbuser := os.Getenv("DB_USER")
	dbpassword := os.Getenv("DB_PASSWORD")
	dbport := "3306"

	// Connect to the database
	connstring := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	println(connstring)
	db, err := gorm.Open(mysql.Open(connstring), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

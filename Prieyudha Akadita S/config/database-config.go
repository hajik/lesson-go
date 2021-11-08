package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
* Method Name
* * Important information is highlighted
* ! Deprecated method, do not use
* ? Should this method be exposed in the public API
* TODO: refactor this method so that is conforms to the API
 */

/**
* SetupDatabaseConnection
* * Important information is highlighted
* ? @param
* ? @return *gorm.DB
* ! -
* TODO: -
 */

func SetupDatabaseConnection() *gorm.DB {

	err := godotenv.Load()

	if err != nil {
		panic("Failed load a file .env")
	}

	dbHost := os.Getenv("localhost")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {

		panic("Failed to create a connection to your database")

	}

	// db.AutoMigrate()

	return db

}

/**
* CloseDatabaseConnection
* * Closing database conection
* ? @param *gorm.DB
* ? @return
* ! -
* TODO: -
 */

func CloseDatabaseConnection(db *gorm.DB) {

	dbSQL, err := db.DB()

	if err != nil {
		panic("Failed to close database connection")
	}

	dbSQL.Close()

}

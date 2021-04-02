package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SetupDatabaseConnection is creating a new connection to database
func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()

	if errEnv != nil {
		panic("Failed to load env file")
	} 

	dbUser := os.Getenv(("DB_USER"))
	dbPass := os.Getenv(("DB_PASS"))
	dbHost := os.Getenv(("DB_HOST"))
	dbName := os.Getenv(("DB_NAME"))
	dbPort := os.Getenv(("DB_PORT"))

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	db, errDb := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if errDb != nil {
		panic("Failed to create a connection to database")
	}

	return db
}

// CloseDatabaseConnection method is closing a connection to DB
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, errDbSQL := db.DB()

	if errDbSQL != nil {
		panic("Failed to close DB connection")
	}

	dbSQL.Close()
}
package database

import (
	"fmt"
	mylog "go-kpl/internal/pkg/logger"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() *gorm.DB {
	status := os.Getenv("STATUS")

	var dbDSN string

	if status == "production" {
		dbDSN = os.Getenv("DATABASE_URL")
		if dbDSN == "" {
			mylog.Panicf("DATABASE_URL is not set for production")
		}
	} else {
		DBHost := os.Getenv("DB_HOST")
		DBUser := os.Getenv("DB_USER")
		DBPassword := os.Getenv("DB_PASS")
		DBName := os.Getenv("DB_NAME")
		DBPort := os.Getenv("DB_PORT")

		if DBHost == "" || DBUser == "" || DBPassword == "" || DBName == "" || DBPort == "" {
			mylog.Panicf("Database environment variables are not properly set for development")
		}

		dbDSN = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			DBHost, DBUser, DBPassword, DBName, DBPort,
		)
	}

	fmt.Println(mylog.ColorizeInfo("\n=========== Setup Database ==========="))
	mylog.Infof("Connecting to database...")

	db, err := gorm.Open(postgres.Open(dbDSN), &gorm.Config{})
	if err != nil {
		mylog.Errorf("Failed to connect to database")
		mylog.Panicf("Failed to connect to database: %v", err)
	}

	mylog.Infof("Success connect to database\n")
	return db
}

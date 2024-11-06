package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/retry19/challenge-hacktiv8/08-assignment/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var PgClient *gorm.DB

func init() {
	var err error

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	PgClient, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Warn,
				IgnoreRecordNotFoundError: true,
				// ParameterizedQueries:      true,
				Colorful: true,
			},
		),
	})
	if err != nil {
		panic(err)
	}

	PgClient.AutoMigrate(&models.Order{}, &models.Item{})
}

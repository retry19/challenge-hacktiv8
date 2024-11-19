package database

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/retry19/challenge-hacktiv8/12-final-project/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type BaseModel struct {
	Id        uint64    `json:"id,omitempty" gorm:"column:id;primary_key;auto_increment"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"column:createdAt;index"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"column:updatedAt;index"`
}

var db *gorm.DB

func Init() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{}, &Photo{}, &SocialMedia{}, &Comment{})
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Error("failed to get sql db", err.Error())
		return
	}

	err = sqlDB.Close()
	if err != nil {
		log.Error("failed to close sql db", err.Error())
		return
	}

	log.Info("sql db closed")
}

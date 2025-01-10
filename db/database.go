package db

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/Anuragdubeyy/ticket-booking-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


func Init(config *config.EnvConfig, DBMigrate func(db *gorm.DB) error) *gorm.DB {
	uri  := fmt.Sprint(`host=%s user=%s password=%s dbname=%s sslmode=%s port=5432`, config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBSslMode)
	db , err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}

	log.Info("connected to database")

	if err := DBMigrate(db); err != nil {
		log.Fatalf("Unable to migrate database: %v", err)
	}
	return db
}
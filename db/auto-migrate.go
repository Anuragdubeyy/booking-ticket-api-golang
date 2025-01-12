package db

import (
	"github.com/Anuragdubeyy/ticket-booking-api/models"
	"gorm.io/gorm"
)

func DBMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.Event{}, &models.Ticket{}, &models.User{})
}
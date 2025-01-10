package repository

import (
	"context"
	"github.com/Anuragdubeyy/ticket-booking-api/models"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	// events := []*models.Event{}

	// events = append(events, &models.Event{
	// 	ID:        "012345678923456789",
	// 	Name:      "Event 1",
	// 	Location:  "Location 1",
	// 	Date:      time.Now(),
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// })

	// return events, nil
	events := []*models.Event{}

	res := r.db.Model(&models.Event{}).Find(&events)

	if res.Error != nil {
		return nil, res.Error
	}

	return events, nil
}

func (r *EventRepository) GetOne(ctx context.Context, eventId string) (*models.Event, error) {
	return nil, nil
}

func (r *EventRepository) CreateOne(ctx context.Context, event models.Event) error {
	return nil
}

func NewEventRepository(db *gorm.DB) models.EventRepository {
	return &EventRepository{
		db: db,
	}
}

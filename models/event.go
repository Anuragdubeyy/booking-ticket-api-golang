
package models

type Event struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	Title     string    `json:"title"`
	Date      time.Time `json:"date"`
	createdAt time.Time `json:"created_at"`
	updatedAt time.Time `json:"updated_at"`
	StartTime string    `json:"start_time"`
	EndTime   string    `json:"end_time"`
}

type EventRepository interface {
	GetMany( ctx context.Context) ([]*Event, error)
	GetOne(ctx context.Context, eventId string) (*Event, error)
	CreateOne(ctx context.Context, event Event) error
	UpdateEvent(event Event) error
	DeleteEvent(id string) error
}
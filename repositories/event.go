package repository
import "github.com/Anuragdubeyy/ticket-booking-api/models"

type EventRepository interface {
	db any
	GetMany( ctx context.Context) ([]*models.Event, error)
	GetOne(ctx context.Context, eventId string) (*models.Event, error)
	CreateOne(ctx context.Context, event models.Event) error
	UpdateEvent(event models.Event) error
	DeleteEvent(id string) error
}

func (r *EventRepository) GetMany( ctx context.Context) ([]*models.Event, error) {
	return nil, nil
}

func (r *EventRepository) GetOne(ctx context.Context, eventId string) (*models.Event, error) {
	return nil, nil
}

func (r *EventRepository) CreateOne(ctx context.Context, event models.Event) error {
	return nil, nil
}
func NewEventRepository(db any) models.EventRepository {
	return &EventRepository{db:db}
}
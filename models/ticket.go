package models

import (
	"context"
	"time"
)

type Ticket struct {
	ID        string    `json:"id"`
	EventID   string    `json:"event_id"`
	Event     Event     `json:"event" gorm:"foreignKey:EventID;contraint:onUpdate:CASCADE,onDelete:CASCADE,OnDelete:CASCADE"`
	Entered   bool      `json:"entered" defualt:"false"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type TicketRepository interface {
	GetMany(ctx context.Context) ([]*Ticket, error)
	GetOne(ctx context.Context, ticketId uint) (*Ticket, error)
	CreateOne(ctx context.Context, ticket *Ticket	) (*Ticket, error)
	UpdateOne(ctx context.Context, ticketId uint, updateData map[string]interface{}) (*Ticket, error)
	DeleteOne(ctx context.Context, ticketId uint) error
}

type ValidateTicket struct {
	TicketId uint `json:"ticketId"`
}
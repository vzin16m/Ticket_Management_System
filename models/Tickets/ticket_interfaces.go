package models

import (
	"github.com/google/uuid"
)

type TicketWrite interface {
	CreateTicket(title string, description string) error
	EditTicket(id uuid.UUID, title string, description string) error
	Updateticket(id uuid.UUID) error
	DeleteTicket(id uuid.UUID) error
}

type TicketRead interface {
	GetTicket(id uuid.UUID) (*Tickets, error)
	GetAllTicket() (*Tickets, error)
}

type ITiecket interface {
	TicketWrite
	TicketRead
}

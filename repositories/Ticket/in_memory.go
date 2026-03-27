package models

import (
	"time"

	"github.com/google/uuid"
)

type Tickets struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"name"`
	Description string    `json:"email"`
	DateCreate  time.Time `json:"datecreate`
	DateClose   time.Time `json:"dateclose`
	Status      bool      `json:"status"`
}

type TicketRepository struct {
	tickets map[uuid.UUID]*Tickets
}

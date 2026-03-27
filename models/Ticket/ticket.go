package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

var ErrNotFound = fmt.Errorf("not found")

func (t *TicketRepository) CreateTicket(title string, desc string) error {
	tck := &Tickets{
		ID:          uuid.New(),
		Title:       title,
		Description: desc,
		DateCreate:  time.Now(),
		Status:      true,
	}
	t.tickets[tck.ID] = tck
	return nil
}

func (t *TicketRepository) EditTicket(id uuid.UUID, title string, description string) error {
	if _, ok := t.tickets[id]; !ok {
		return ErrNotFound
	}
	t.tickets[id].Title = title
	t.tickets[id].Title = title
	return nil
}

func (t *TicketRepository) UpdateTicket(id uuid.UUID) error {
	if _, ok := t.tickets[id]; !ok {
		return ErrNotFound
	}
	t.tickets[id].Status = false
	return nil
}

func (t *TicketRepository) DeleteTicket(id uuid.UUID) error {
	if _, ok := t.tickets[id]; !ok {
		return ErrNotFound
	}
	delete(t.tickets, id)
	return nil
}

func (t *TicketRepository) GetTicket(id uuid.UUID) (*Tickets, error) {
	if tck, ok := t.tickets[id]; ok {
		return tck, nil
	}
	return nil, ErrNotFound
}

func (t *TicketRepository) GetAllTicket() ([]Tickets, error) {
	var tck []Tickets
	for _, v := range t.tickets {
		tck = append(tck, *v)
	}

	return tck, nil
}

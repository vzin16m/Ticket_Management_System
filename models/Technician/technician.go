package technicians

import (
	"github.com/google/uuid"
)

type Technician struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Email  string    `json:"email"`
	Status bool      `json:"status"`
}

// mapper
func (t *Technician) ToUpdate() *TechnicianToUpdate {
	return &TechnicianToUpdate{
		ID:     t.ID,
		Name:   &t.Name,
		Email:  &t.Email,
		Status: &t.Status,
	}
}

// DTO
type TechnicianToUpdate struct {
	ID     uuid.UUID `json:"id"`
	Name   *string   `json:"name"`
	Email  *string   `json:"email"`
	Status *bool     `json:"status"`
}

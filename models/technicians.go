package models

import (
	"github.com/google/uuid"
)

type Technician struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Email  string    `json:"email"`
	Status bool      `json:"status"`
}

type TechnicianStore struct {
	technicians map[uuid.UUID]*Technician
}

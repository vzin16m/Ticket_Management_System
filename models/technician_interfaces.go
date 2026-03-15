package models

import (
	"github.com/google/uuid"
)

type ITechnicianModify interface {
	CreateTechnician(name string, email string) error
	UpdateTechnician(id uuid.UUID, name string, email string) error
	DeleteTechnician(id uuid.UUID) error
}

type ITechnicianView interface {
	GetTechnician(id uuid.UUID) (*Technician, error)
	GetAllTechnicians() ([]*Technician, error)
}

type ITechnician interface {
	ITechnicianModify
	ITechnicianView
}

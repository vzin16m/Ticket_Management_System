package technicians

import (
	"github.com/google/uuid"
)

type ITechnicianWrite interface {
	SaveTechnician(*Technician) error
	UpdateTechnician(*TechnicianToUpdate) error
	DeleteTechnician(id uuid.UUID) error
}

type ITechnicianRead interface {
	GetTechnician(id uuid.UUID) (*Technician, error)
	GetAllTechnicians() ([]*Technician, error)
}

type ITechnicianRepositoty interface {
	ITechnicianWrite
	ITechnicianRead
}

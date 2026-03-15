package models

import (
	"github.com/google/uuid"
)

func (t *TechnicianStore) CreateTechnician(name string, email string) error {
	tec := &Technician{
		ID:     uuid.New(),
		Name:   name,
		Email:  email,
		Status: true,
	}
	t.technicians[tec.ID] = tec
	return nil
}

func (t *TechnicianStore) UpdateTechnician(id uuid.UUID, name string, email string) error {
	if _, ok := t.technicians[id]; !ok {
		return ErrNotFound
	}
	t.technicians[id].Name = name
	t.technicians[id].Email = email
	return nil
}

func (t *TechnicianStore) DeleteTechnician(id uuid.UUID) error {
	if _, ok := t.technicians[id]; !ok {
		return ErrNotFound
	}
	delete(t.technicians, id)
	return nil
}

func (t *TechnicianStore) GetTechnician(id uuid.UUID) (*Technician, error) {
	if tec, ok := t.technicians[id]; ok {
		return tec, nil
	}
	return nil, ErrNotFound
}

func (t *TechnicianStore) GetAllTechnicians() (*[]Technician, error) {
	var tec []Technician
	for _, v := range t.technicians {
		tec = append(tec, *v)
	}
	return &tec, nil
}

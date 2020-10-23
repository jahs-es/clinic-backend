package _interface

import (
	"github.com/jahs/clinic-backend/src/domain"
	"github.com/jahs/clinic-backend/src/shared/entity"
)

type IPatientRepository interface {
	Get(id entity.ID) (*domain.Patient, error)
	GetByEmail(email string) (*domain.Patient, error)
	Search(e *domain.Patient) ([]*domain.Patient, error)
	Create(e *domain.Patient) (entity.ID, error)
	Update(e *domain.Patient) error
	Delete(id entity.ID) error
}


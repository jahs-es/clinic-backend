package repository

import (
	"github.com/jahs/clinic-backend/adapter/entity"
	"github.com/jahs/clinic-backend/domain/model"
)

type PatientRepository interface {
	Get(id entity.ID) (*model.Patient, error)
	GetByEmail(email string) (*model.Patient, error)
	Search(e *model.Patient) ([]*model.Patient, error)
	Create(e *model.Patient) (entity.ID, error)
	Update(e *model.Patient) error
	Delete(id entity.ID) error
}


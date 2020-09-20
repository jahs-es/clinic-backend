package repository

import (
	"github.com/jahs/clinic-backend/adapter/entity"
	presenterDTO "github.com/jahs/clinic-backend/adapter/presenter"
	"github.com/jahs/clinic-backend/domain/model"
)

type PatientTreatmentRepository interface {
	Get(id entity.ID) (*model.PatientTreatment, error)
	Search(e *presenterDTO.PatientTreatmentDTO) ([]*presenterDTO.PatientTreatmentDTO, error)
	Create(e *model.PatientTreatment) (entity.ID, error)
	Update(e *model.PatientTreatment) error
	Delete(id entity.ID) error
}


package _interface

import (
	"github.com/jahs/clinic-backend/src/domain"
	"github.com/jahs/clinic-backend/src/infrastructure/presenter/dto"
	"github.com/jahs/clinic-backend/src/shared/entity"
)

type IPatientTreatmentRepository interface {
	Get(id entity.ID) (*domain.PatientTreatment, error)
	Search(e *dto.PatientTreatmentDTO) ([]*dto.PatientTreatmentDTO, error)
	Create(e *domain.PatientTreatment) (entity.ID, error)
	Update(e *domain.PatientTreatment) error
	Delete(id entity.ID) error
}


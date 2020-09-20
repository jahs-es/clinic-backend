package validator

import (
	"github.com/jahs/clinic-backend/usecase/repository"
)

type patientTreatmentValidator struct {
	PatientTreatmentRepository repository.PatientTreatmentRepository
}

type PatientTreatmentValidator interface {
	PatientTreatmentNotExist(email string) error
}

func NewPatientTreatmentValidator(r repository.PatientTreatmentRepository) PatientTreatmentValidator {
	return &patientTreatmentValidator{r}
}

func (uv *patientTreatmentValidator) PatientTreatmentNotExist(id string) error {
	return nil
}

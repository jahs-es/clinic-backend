package service

import (
	"github.com/jahs/clinic-backend/src/infrastructure/persistence/interface"
)

type patientTreatmentValidator struct {
	PatientTreatmentRepository _interface.IPatientTreatmentRepository
}

type PatientTreatmentValidator interface {
	PatientTreatmentNotExist(email string) error
}

func NewPatientTreatmentValidator(r _interface.IPatientTreatmentRepository) PatientTreatmentValidator {
	return &patientTreatmentValidator{r}
}

func (uv *patientTreatmentValidator) PatientTreatmentNotExist(id string) error {
	return nil
}

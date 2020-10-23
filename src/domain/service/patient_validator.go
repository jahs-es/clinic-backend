package service

import (
	"fmt"
	"github.com/jahs/clinic-backend/src/domain/exception"
	"github.com/jahs/clinic-backend/src/infrastructure/persistence/interface"
)

type patientValidator struct {
	PatientRepository _interface.IPatientRepository
}

type PatientValidator interface {
	PatientNotExist(email string) error
}

func NewPatientValidator(r _interface.IPatientRepository) PatientValidator {
	return &patientValidator{r}
}

func (uv *patientValidator) PatientNotExist(email string) error {
	data, err := uv.PatientRepository.GetByEmail(email)

	if err != nil {
		return err
	}

	fmt.Println(data)

	if data.Email != "" {
		return exception.ErrAlreadyExist
	}

	return nil
}

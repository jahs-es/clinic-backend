package validator

import (
	"fmt"
	"github.com/jahs/clinic-backend/usecase/exception"
	"github.com/jahs/clinic-backend/usecase/repository"
)

type patientValidator struct {
	PatientRepository repository.PatientRepository
}

type PatientValidator interface {
	PatientNotExist(email string) error
}

func NewPatientValidator(r repository.PatientRepository) PatientValidator {
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

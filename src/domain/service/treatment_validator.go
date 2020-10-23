package service

import (
	"fmt"
	"github.com/jahs/clinic-backend/src/domain/exception"
	"github.com/jahs/clinic-backend/src/infrastructure/persistence/interface"
)

type treatmentValidator struct {
	TreatmentRepository _interface.ITreatmentRepository
}

type TreatmentValidator interface {
	TreatmentNotExist(email string) error
}

func NewTreatmentValidator(r _interface.ITreatmentRepository) TreatmentValidator {
	return &treatmentValidator{r}
}

func (uv *treatmentValidator) TreatmentNotExist(name string) error {
	data, err := uv.TreatmentRepository.GetByName(name)

	if err != nil {
		return err
	}

	fmt.Println(data)

	if data.Name != "" {
		return exception.ErrAlreadyExist
	}

	return nil
}

package validator

import (
	"fmt"
	"github.com/jahs/clinic-backend/usecase/exception"
	"github.com/jahs/clinic-backend/usecase/repository"
)

type treatmentValidator struct {
	TreatmentRepository repository.TreatmentRepository
}

type TreatmentValidator interface {
	TreatmentNotExist(email string) error
}

func NewTreatmentValidator(r repository.TreatmentRepository) TreatmentValidator {
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

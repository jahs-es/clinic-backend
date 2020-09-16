package interactor

import (
	"fmt"
	"github.com/jahs/clinic-backend/adapter/entity"
	"github.com/jahs/clinic-backend/adapter/password"
	presenterDTO "github.com/jahs/clinic-backend/adapter/presenter"
	"github.com/jahs/clinic-backend/domain/model"
	"github.com/jahs/clinic-backend/usecase/presenter"
	"github.com/jahs/clinic-backend/usecase/repository"
	"github.com/jahs/clinic-backend/usecase/validator"
	"time"
)

type patientInteractor struct {
	PatientRepository repository.PatientRepository
	PatientPresenter  presenter.PatientPresenter
	Service           password.Service
	Validator         validator.PatientValidator
}

type PatientInteractor interface {
	Create(e *model.Patient) (entity.ID, error)
	Get(id entity.ID) (*presenterDTO.PatientDTO, error)
	Delete(id entity.ID) error
	Find(e *model.Patient) ([]*presenterDTO.PatientDTO, error)
}

func NewPatientInteractor(r repository.PatientRepository, p presenter.PatientPresenter, s password.Service, v validator.PatientValidator) *patientInteractor {
	return &patientInteractor{r, p, s, v}
}

func (us *patientInteractor) Create(e *model.Patient) (entity.ID, error) {
	err := us.Validator.PatientNotExist(e.Email)
	if err != nil {
		return e.ID, err
	}

	e.ID = entity.NewID()
	e.CreatedAt = time.Now()

	return us.PatientRepository.Create(e)
}

func (us *patientInteractor) Get(id entity.ID) (*presenterDTO.PatientDTO, error) {
	u, err := us.PatientRepository.Get(id)

	if err != nil {
		return nil, err
	}

	fmt.Println(u)

	return &presenterDTO.PatientDTO{
		ID:      u.ID,
		Name:    u.Name,
		Email:   u.Email,
		Address: u.Address,
		Phone:   u.Phone,
		Active:  u.Active,
	}, nil
}

func (us *patientInteractor) Delete(id entity.ID) error {
	return us.PatientRepository.Delete(id)
}

func (us *patientInteractor) Find(e *model.Patient) ([]*presenterDTO.PatientDTO, error) {
	u, err := us.PatientRepository.Search(e)

	if err != nil {
		return nil, err
	}
	return us.PatientPresenter.Response(u), nil
}

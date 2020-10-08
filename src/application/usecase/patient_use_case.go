package usecase

import (
	"fmt"
	"github.com/jahs/clinic-backend/src/domain"
	"github.com/jahs/clinic-backend/src/domain/service"
	"github.com/jahs/clinic-backend/src/infrastructure/persistence/interface"
	presenter2 "github.com/jahs/clinic-backend/src/infrastructure/presenter"
	"github.com/jahs/clinic-backend/src/infrastructure/presenter/dto"
	"github.com/jahs/clinic-backend/src/shared/entity"
	"github.com/jahs/clinic-backend/src/shared/password"
	"time"
)

type patientUseCase struct {
	PatientRepository _interface.IPatientRepository
	PatientPresenter  presenter2.IPatientPresenter
	Service           password.Service
	Validator         service.PatientValidator
}

type IPatientUseCase interface {
	Create(e *domain.Patient) (entity.ID, error)
	Update(e *domain.Patient) error
	Get(id entity.ID) (*dto.PatientDTO, error)
	Delete(id entity.ID) error
	Find(e *domain.Patient) ([]*dto.PatientDTO, error)
}

func NewPatientInteractor(r _interface.IPatientRepository, p presenter2.IPatientPresenter, s password.Service, v service.PatientValidator) *patientUseCase {
	return &patientUseCase{r, p, s, v}
}

func (us *patientUseCase) Create(e *domain.Patient) (entity.ID, error) {
	err := us.Validator.PatientNotExist(e.Email)
	if err != nil {
		return e.ID, err
	}

	//e.ID = entity.NewID()
	e.CreatedAt = time.Now()

	return us.PatientRepository.Create(e)
}

func (us *patientUseCase) Update(e *domain.Patient) error {
	return us.PatientRepository.Update(e)
}

func (us *patientUseCase) Get(id entity.ID) (*dto.PatientDTO, error) {
	u, err := us.PatientRepository.Get(id)

	if err != nil {
		return nil, err
	}

	fmt.Println(u)

	return &dto.PatientDTO{
		ID:      u.ID,
		Name:    u.Name,
		Email:   u.Email,
		Address: u.Address,
		Phone:   u.Phone,
		Active:  u.Active,
	}, nil
}

func (us *patientUseCase) Delete(id entity.ID) error {
	return us.PatientRepository.Delete(id)
}

func (us *patientUseCase) Find(e *domain.Patient) ([]*dto.PatientDTO, error) {
	u, err := us.PatientRepository.Search(e)

	if err != nil {
		return nil, err
	}
	return us.PatientPresenter.Response(u), nil
}

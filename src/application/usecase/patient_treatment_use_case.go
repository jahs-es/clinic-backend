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

type patientTreatmentUseCase struct {
	PatientTreatmentRepository _interface.IPatientTreatmentRepository
	PatientTreatmentPresenter  presenter2.IPatientTreatmentPresenter
	Service                    password.Service
	Validator                  service.PatientTreatmentValidator
}

type IPatientTreatmentUseCase interface {
	Create(e *domain.PatientTreatment) (entity.ID, error)
	Update(e *domain.PatientTreatment) error
	Get(id entity.ID) (*dto.PatientTreatmentDTO, error)
	Delete(id entity.ID) error
	Find(e *dto.PatientTreatmentDTO) ([]*dto.PatientTreatmentDTO, error)
}

func NewPatientTreatmentInteractor(r _interface.IPatientTreatmentRepository, p presenter2.IPatientTreatmentPresenter, s password.Service, v service.PatientTreatmentValidator) *patientTreatmentUseCase {
	return &patientTreatmentUseCase{r, p, s, v}
}

func (us *patientTreatmentUseCase) Create(e *domain.PatientTreatment) (entity.ID, error) {
	err := us.Validator.PatientTreatmentNotExist("e.ID")
	if err != nil {
		return e.ID, err
	}

	//e.ID = entity.NewID()
	e.CreatedAt = time.Now()

	return us.PatientTreatmentRepository.Create(e)
}

func (us *patientTreatmentUseCase) Update(e *domain.PatientTreatment) error {
	return us.PatientTreatmentRepository.Update(e)
}

func (us *patientTreatmentUseCase) Get(id entity.ID) (*dto.PatientTreatmentDTO, error) {
	u, err := us.PatientTreatmentRepository.Get(id)

	if err != nil {
		return nil, err
	}

	fmt.Println(u)

	return &dto.PatientTreatmentDTO{
		ID:          u.ID,
		PatientId:   u.PatientId,
		TreatmentId: u.TreatmentId,
		Detail:      u.Detail,
		Active:      u.Active,
	}, nil
}

func (us *patientTreatmentUseCase) Delete(id entity.ID) error {
	return us.PatientTreatmentRepository.Delete(id)
}

func (us *patientTreatmentUseCase) Find(e *dto.PatientTreatmentDTO) ([]*dto.PatientTreatmentDTO, error) {
	u, err := us.PatientTreatmentRepository.Search(e)

	if err != nil {
		return nil, err
	}
	return u, nil
}

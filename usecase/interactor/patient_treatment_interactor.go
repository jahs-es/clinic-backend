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

type patientTreatmentInteractor struct {
	PatientTreatmentRepository repository.PatientTreatmentRepository
	PatientTreatmentPresenter  presenter.PatientTreatmentPresenter
	Service                    password.Service
	Validator                  validator.PatientTreatmentValidator
}

type PatientTreatmentInteractor interface {
	Create(e *model.PatientTreatment) (entity.ID, error)
	Update(e *model.PatientTreatment) error
	Get(id entity.ID) (*presenterDTO.PatientTreatmentDTO, error)
	Delete(id entity.ID) error
	Find(e *presenterDTO.PatientTreatmentDTO) ([]*presenterDTO.PatientTreatmentDTO, error)
}

func NewPatientTreatmentInteractor(r repository.PatientTreatmentRepository, p presenter.PatientTreatmentPresenter, s password.Service, v validator.PatientTreatmentValidator) *patientTreatmentInteractor {
	return &patientTreatmentInteractor{r, p, s, v}
}

func (us *patientTreatmentInteractor) Create(e *model.PatientTreatment) (entity.ID, error) {
	err := us.Validator.PatientTreatmentNotExist("e.ID")
	if err != nil {
		return e.ID, err
	}

	e.ID = entity.NewID()
	e.CreatedAt = time.Now()

	return us.PatientTreatmentRepository.Create(e)
}

func (us *patientTreatmentInteractor) Update(e *model.PatientTreatment) error {
	return us.PatientTreatmentRepository.Update(e)
}

func (us *patientTreatmentInteractor) Get(id entity.ID) (*presenterDTO.PatientTreatmentDTO, error) {
	u, err := us.PatientTreatmentRepository.Get(id)

	if err != nil {
		return nil, err
	}

	fmt.Println(u)

	return &presenterDTO.PatientTreatmentDTO{
		ID:          u.ID,
		PatientId:   u.PatientId,
		TreatmentId: u.TreatmentId,
		Detail:      u.Detail,
		Active:      u.Active,
	}, nil
}

func (us *patientTreatmentInteractor) Delete(id entity.ID) error {
	return us.PatientTreatmentRepository.Delete(id)
}

func (us *patientTreatmentInteractor) Find(e *presenterDTO.PatientTreatmentDTO) ([]*presenterDTO.PatientTreatmentDTO, error) {
	u, err := us.PatientTreatmentRepository.Search(e)

	if err != nil {
		return nil, err
	}
	return u, nil
}

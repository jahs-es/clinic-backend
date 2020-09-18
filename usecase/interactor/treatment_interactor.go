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

type treatmentInteractor struct {
	TreatmentRepository repository.TreatmentRepository
	TreatmentPresenter  presenter.TreatmentPresenter
	Service           password.Service
	Validator         validator.TreatmentValidator
}

type TreatmentInteractor interface {
	Create(e *model.Treatment) (entity.ID, error)
	Update(e *model.Treatment) error
	Get(id entity.ID) (*presenterDTO.TreatmentDTO, error)
	Delete(id entity.ID) error
	Find(e *model.Treatment) ([]*presenterDTO.TreatmentDTO, error)
}

func NewTreatmentInteractor(r repository.TreatmentRepository, p presenter.TreatmentPresenter, s password.Service, v validator.TreatmentValidator) *treatmentInteractor {
	return &treatmentInteractor{r, p, s, v}
}

func (us *treatmentInteractor) Create(e *model.Treatment) (entity.ID, error) {
	err := us.Validator.TreatmentNotExist(e.Name)
	if err != nil {
		return e.ID, err
	}

	e.ID = entity.NewID()
	e.CreatedAt = time.Now()

	return us.TreatmentRepository.Create(e)
}

func (us *treatmentInteractor) Update(e *model.Treatment) error {
	return us.TreatmentRepository.Update(e)
}

func (us *treatmentInteractor) Get(id entity.ID) (*presenterDTO.TreatmentDTO, error) {
	u, err := us.TreatmentRepository.Get(id)

	if err != nil {
		return nil, err
	}

	fmt.Println(u)

	return &presenterDTO.TreatmentDTO{
		ID:      u.ID,
		Name:    u.Name,
		Active:  u.Active,
	}, nil
}

func (us *treatmentInteractor) Delete(id entity.ID) error {
	return us.TreatmentRepository.Delete(id)
}

func (us *treatmentInteractor) Find(e *model.Treatment) ([]*presenterDTO.TreatmentDTO, error) {
	u, err := us.TreatmentRepository.Search(e)

	if err != nil {
		return nil, err
	}
	return us.TreatmentPresenter.Response(u), nil
}

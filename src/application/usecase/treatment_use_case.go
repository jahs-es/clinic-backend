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

type treatmentUseCase struct {
	TreatmentRepository _interface.ITreatmentRepository
	TreatmentPresenter  presenter2.ITreatmentPresenter
	Service             password.Service
	Validator           service.TreatmentValidator
}

type ITreatmentUseCase interface {
	Create(e *domain.Treatment) (entity.ID, error)
	Update(e *domain.Treatment) error
	Get(id entity.ID) (*dto.TreatmentDTO, error)
	Delete(id entity.ID) error
	Find(e *domain.Treatment) ([]*dto.TreatmentDTO, error)
}

func NewTreatmentUseCase(r _interface.ITreatmentRepository, p presenter2.ITreatmentPresenter, s password.Service, v service.TreatmentValidator) *treatmentUseCase {
	return &treatmentUseCase{r, p, s, v}
}

func (us *treatmentUseCase) Create(e *domain.Treatment) (entity.ID, error) {
	err := us.Validator.TreatmentNotExist(e.Name)
	if err != nil {
		return e.ID, err
	}

	//e.ID = entity.NewID()
	e.CreatedAt = time.Now()

	return us.TreatmentRepository.Create(e)
}

func (us *treatmentUseCase) Update(e *domain.Treatment) error {
	return us.TreatmentRepository.Update(e)
}

func (us *treatmentUseCase) Get(id entity.ID) (*dto.TreatmentDTO, error) {
	u, err := us.TreatmentRepository.Get(id)

	if err != nil {
		return nil, err
	}

	fmt.Println(u)

	return &dto.TreatmentDTO{
		ID:      u.ID,
		Name:    u.Name,
		Active:  u.Active,
	}, nil
}

func (us *treatmentUseCase) Delete(id entity.ID) error {
	return us.TreatmentRepository.Delete(id)
}

func (us *treatmentUseCase) Find(e *domain.Treatment) ([]*dto.TreatmentDTO, error) {
	u, err := us.TreatmentRepository.Search(e)

	if err != nil {
		return nil, err
	}
	return us.TreatmentPresenter.Response(u), nil
}

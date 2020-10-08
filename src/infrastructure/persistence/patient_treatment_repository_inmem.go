package persistence

import (
	"github.com/jahs/clinic-backend/src/domain"
	"github.com/jahs/clinic-backend/src/domain/exception"
	user "github.com/jahs/clinic-backend/src/infrastructure/persistence/test"
	"github.com/jahs/clinic-backend/src/infrastructure/presenter/dto"
	entity2 "github.com/jahs/clinic-backend/src/shared/entity"
)

type IPatientTreatmentRepo struct {
	m map[entity2.ID]*domain.PatientTreatment
}

func NewInmemPatientTreatmentRepository() *IPatientTreatmentRepo {
	var m = map[entity2.ID]*domain.PatientTreatment{}
	return &IPatientTreatmentRepo{
		m: m,
	}
}

func (r *IPatientTreatmentRepo) Create(e *domain.PatientTreatment) (entity2.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

func (r *IPatientTreatmentRepo) Get(id entity2.ID) (*domain.PatientTreatment, error) {
	if r.m[id] == nil {
		return nil, exception.ErrNotFound
	}
	return r.m[id], nil
}

func (r *IPatientTreatmentRepo) Update(e *domain.PatientTreatment) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

func (r *IPatientTreatmentRepo) Search(e *dto.PatientTreatmentDTO) ([]*dto.PatientTreatmentDTO, error) {
	var d []*dto.PatientTreatmentDTO

	d = append(d, user.NewFixturePatientTreatment())

	return d, nil
}

func (r *IPatientTreatmentRepo) Delete(id entity2.ID) error {
	if r.m[id] == nil {
		return exception.ErrNotFound
	}
	r.m[id] = nil
	return nil
}

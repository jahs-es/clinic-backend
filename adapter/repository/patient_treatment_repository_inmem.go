package repository

import (
	entity2 "github.com/jahs/clinic-backend/adapter/entity"
	"github.com/jahs/clinic-backend/adapter/presenter"
	user "github.com/jahs/clinic-backend/adapter/repository/test"
	"github.com/jahs/clinic-backend/domain/model"
	"github.com/jahs/clinic-backend/usecase/exception"
)

type IPatientTreatmentRepo struct {
	m map[entity2.ID]*model.PatientTreatment
}

func NewInmemPatientTreatmentRepository() *IPatientTreatmentRepo {
	var m = map[entity2.ID]*model.PatientTreatment{}
	return &IPatientTreatmentRepo{
		m: m,
	}
}

func (r *IPatientTreatmentRepo) Create(e *model.PatientTreatment) (entity2.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

func (r *IPatientTreatmentRepo) Get(id entity2.ID) (*model.PatientTreatment, error) {
	if r.m[id] == nil {
		return nil, exception.ErrNotFound
	}
	return r.m[id], nil
}

func (r *IPatientTreatmentRepo) Update(e *model.PatientTreatment) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

func (r *IPatientTreatmentRepo) Search(e *presenter.PatientTreatmentDTO) ([]*presenter.PatientTreatmentDTO, error) {
	var d []*presenter.PatientTreatmentDTO

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

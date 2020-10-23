package persistence

import (
	"github.com/jahs/clinic-backend/src/domain"
	"github.com/jahs/clinic-backend/src/domain/exception"
	user "github.com/jahs/clinic-backend/src/infrastructure/persistence/test"
	entity2 "github.com/jahs/clinic-backend/src/shared/entity"
)

type IPatientRepo struct {
	m map[entity2.ID]*domain.Patient
}

func NewInmemPatientRepository() *IPatientRepo {
	var m = map[entity2.ID]*domain.Patient{}
	return &IPatientRepo{
		m: m,
	}
}

func (r *IPatientRepo) Create(e *domain.Patient) (entity2.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

func (r *IPatientRepo) Get(id entity2.ID) (*domain.Patient, error) {
	if r.m[id] == nil {
		return nil, exception.ErrNotFound
	}
	return r.m[id], nil
}

func (r *IPatientRepo) GetByEmail(email string) (*domain.Patient, error) {
	var d *domain.Patient

	d = &domain.Patient{Email: ""}

	if email == "" {
		return nil, exception.ErrNotFound
	}

	return d, nil
}

func (r *IPatientRepo) Update(e *domain.Patient) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

func (r *IPatientRepo) Search(e *domain.Patient) ([]*domain.Patient, error) {
	var d []*domain.Patient

	d = append(d, user.NewFixturePatient())

	return d, nil
}

func (r *IPatientRepo) Delete(id entity2.ID) error {
	if r.m[id] == nil {
		return exception.ErrNotFound
	}
	r.m[id] = nil
	return nil
}

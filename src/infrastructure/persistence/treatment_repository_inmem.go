package persistence

import (
	"github.com/jahs/clinic-backend/src/domain"
	"github.com/jahs/clinic-backend/src/domain/exception"
	user "github.com/jahs/clinic-backend/src/infrastructure/persistence/test"
	entity2 "github.com/jahs/clinic-backend/src/shared/entity"
)

type ITreatmentRepo struct {
	m map[entity2.ID]*domain.Treatment
}

func NewInmemTreatmentRepository() *ITreatmentRepo {
	var m = map[entity2.ID]*domain.Treatment{}
	return &ITreatmentRepo{
		m: m,
	}
}

func (r *ITreatmentRepo) Create(e *domain.Treatment) (entity2.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

func (r *ITreatmentRepo) Get(id entity2.ID) (*domain.Treatment, error) {
	if r.m[id] == nil {
		return nil, exception.ErrNotFound
	}
	return r.m[id], nil
}

func (r *ITreatmentRepo) GetByName(name string) (*domain.Treatment, error) {
	var d *domain.Treatment

	d = &domain.Treatment{Name: ""}

	if name == "" {
		return nil, exception.ErrNotFound
	}

	return d, nil
}

func (r *ITreatmentRepo) Update(e *domain.Treatment) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

func (r *ITreatmentRepo) Search(e *domain.Treatment) ([]*domain.Treatment, error) {
	var d []*domain.Treatment

	d = append(d, user.NewFixtureTreatment())

	return d, nil
}

func (r *ITreatmentRepo) Delete(id entity2.ID) error {
	if r.m[id] == nil {
		return exception.ErrNotFound
	}
	r.m[id] = nil
	return nil
}

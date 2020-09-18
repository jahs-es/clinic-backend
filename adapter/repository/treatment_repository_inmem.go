package repository

import (
	entity2 "github.com/jahs/clinic-backend/adapter/entity"
	user "github.com/jahs/clinic-backend/adapter/repository/test"
	"github.com/jahs/clinic-backend/domain/model"
	"github.com/jahs/clinic-backend/usecase/exception"
)

//ITreatmentRepo in memory repo
type ITreatmentRepo struct {
	m map[entity2.ID]*model.Treatment
}

func NewInmemTreatmentRepository() *ITreatmentRepo {
	var m = map[entity2.ID]*model.Treatment{}
	return &ITreatmentRepo{
		m: m,
	}
}

//Create an user
func (r *ITreatmentRepo) Create(e *model.Treatment) (entity2.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

//Get an user
func (r *ITreatmentRepo) Get(id entity2.ID) (*model.Treatment, error) {
	if r.m[id] == nil {
		return nil, exception.ErrNotFound
	}
	return r.m[id], nil
}

//Get an user by email
func (r *ITreatmentRepo) GetByName(name string) (*model.Treatment, error) {
	var d *model.Treatment

	d = &model.Treatment{Name: ""}

	if name == "" {
		return nil, exception.ErrNotFound
	}

	return d, nil
}

//Update an user
func (r *ITreatmentRepo) Update(e *model.Treatment) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

//Search users
func (r *ITreatmentRepo) Search(e *model.Treatment) ([]*model.Treatment, error) {
	var d []*model.Treatment

	d = append(d, user.NewFixtureTreatment())

	return d, nil
}

//Delete an user
func (r *ITreatmentRepo) Delete(id entity2.ID) error {
	if r.m[id] == nil {
		return exception.ErrNotFound
	}
	r.m[id] = nil
	return nil
}

package repository

import (
	entity2 "github.com/jahs/clinic-backend/adapter/entity"
	user "github.com/jahs/clinic-backend/adapter/repository/test"
	"github.com/jahs/clinic-backend/domain/model"
	"github.com/jahs/clinic-backend/usecase/exception"
)

//IPatientRepo in memory repo
type IPatientRepo struct {
	m map[entity2.ID]*model.Patient
}

func NewInmemPatientRepository() *IPatientRepo {
	var m = map[entity2.ID]*model.Patient{}
	return &IPatientRepo{
		m: m,
	}
}

//Create an user
func (r *IPatientRepo) Create(e *model.Patient) (entity2.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

//Get an user
func (r *IPatientRepo) Get(id entity2.ID) (*model.Patient, error) {
	if r.m[id] == nil {
		return nil, exception.ErrNotFound
	}
	return r.m[id], nil
}

//Get an user by email
func (r *IPatientRepo) GetByEmail(email string) (*model.Patient, error) {
	var d *model.Patient

	d = &model.Patient{Email: ""}

	if email == "" {
		return nil, exception.ErrNotFound
	}

	return d, nil
}

//Update an user
func (r *IPatientRepo) Update(e *model.Patient) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

//Search users
func (r *IPatientRepo) Search(e *model.Patient) ([]*model.Patient, error) {
	var d []*model.Patient

	d = append(d, user.NewFixturePatient())

	return d, nil
}

//Delete an user
func (r *IPatientRepo) Delete(id entity2.ID) error {
	if r.m[id] == nil {
		return exception.ErrNotFound
	}
	r.m[id] = nil
	return nil
}
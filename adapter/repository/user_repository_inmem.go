package repository

import (
	entity2 "github.com/jahs/clinic-backend/adapter/entity"
	user "github.com/jahs/clinic-backend/adapter/repository/test"
	"github.com/jahs/clinic-backend/domain/model"
	"github.com/jahs/clinic-backend/usecase/exception"
)

//IRepo in memory repo
type IUserRepo struct {
	m map[entity2.ID]*model.User
}

//NewInmemUserRepository create new repository
func NewInmemUserRepository() *IUserRepo {
	var m = map[entity2.ID]*model.User{}
	return &IUserRepo{
		m: m,
	}
}

//Create an user
func (r *IUserRepo) Create(e *model.User) (entity2.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

//Get an user
func (r *IUserRepo) Get(id entity2.ID) (*model.User, error) {
	if r.m[id] == nil {
		return nil, exception.ErrNotFound
	}
	return r.m[id], nil
}

//Get an user by email
func (r *IUserRepo) GetByEmail(email string) (*model.User, error) {
	var d *model.User

	d = &model.User{Email: ""}

	if email == "" {
		return nil, exception.ErrNotFound
	}

	return d, nil
}

//Update an user
func (r *IUserRepo) Update(e *model.User) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

//Search users
func (r *IUserRepo) Search(e *model.User) ([]*model.User, error) {
	var d []*model.User

	d = append(d, user.NewFixtureUser())

	return d, nil
}

//Delete an user
func (r *IUserRepo) Delete(id entity2.ID) error {
	if r.m[id] == nil {
		return exception.ErrNotFound
	}
	r.m[id] = nil
	return nil
}

package persistence

import (
	"github.com/jahs/clinic-backend/src/domain"
	"github.com/jahs/clinic-backend/src/domain/exception"
	user "github.com/jahs/clinic-backend/src/infrastructure/persistence/test"
	entity2 "github.com/jahs/clinic-backend/src/shared/entity"
)

type IUserRepo struct {
	m map[entity2.ID]*domain.User
}

func NewInmemUserRepository() *IUserRepo {
	var m = map[entity2.ID]*domain.User{}
	return &IUserRepo{
		m: m,
	}
}

func (r *IUserRepo) Create(e *domain.User) (entity2.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

func (r *IUserRepo) Get(id entity2.ID) (*domain.User, error) {
	if r.m[id] == nil {
		return nil, exception.ErrNotFound
	}
	return r.m[id], nil
}

func (r *IUserRepo) GetByEmail(email string) (*domain.User, error) {
	var d *domain.User

	d = &domain.User{Email: ""}

	if email == "" {
		return nil, exception.ErrNotFound
	}

	return d, nil
}

func (r *IUserRepo) Update(e *domain.User) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

func (r *IUserRepo) Search(e *domain.User) ([]*domain.User, error) {
	var d []*domain.User

	d = append(d, user.NewFixtureUser())

	return d, nil
}

func (r *IUserRepo) Delete(id entity2.ID) error {
	if r.m[id] == nil {
		return exception.ErrNotFound
	}
	r.m[id] = nil
	return nil
}

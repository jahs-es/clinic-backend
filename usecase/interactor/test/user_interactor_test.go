package test

import (
	"github.com/jahs/clinic-backend/adapter/password/test"
	"github.com/jahs/clinic-backend/adapter/presenter"
	"github.com/jahs/clinic-backend/adapter/repository"
	"github.com/jahs/clinic-backend/adapter/repository/test"
	"github.com/jahs/clinic-backend/usecase/exception"
	"github.com/jahs/clinic-backend/usecase/interactor"
	"github.com/jahs/clinic-backend/usecase/validator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
	repo := repository.NewInmemRepository()
	presenter := presenter.NewUserPresenter()
	validator := validator.NewUserValidator(repo)

	m := interactor.NewUserInteractor(repo, presenter, test.NewFakeService(), validator)
	u := user.NewFixtureUser()
	id, err := m.Create(u)

	assert.Nil(t, err)
	assert.Equal(t, u.ID, id)
	assert.False(t, u.CreatedAt.IsZero())
	assert.True(t, u.UpdatedAt.IsZero())
}

func Test_SearchAndFind(t *testing.T) {
	repo := repository.NewInmemRepository()
	presenter := presenter.NewUserPresenter()
	validator := validator.NewUserValidator(repo)

	m := interactor.NewUserInteractor(repo, presenter, test.NewFakeService(), validator)
	u := user.NewFixtureUser()
	_, _ = m.Create(u)

	searchUser := user.NewFixtureSearchUser()

	t.Run("findUser", func(t *testing.T) {
		c, err := m.Find(searchUser)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))
		assert.Equal(t, "Jahs", c[0].Name)
	})
	t.Run("get", func(t *testing.T) {
		saved, err := m.Get(u.ID)
		assert.Nil(t, err)
		assert.Equal(t, u.Name, saved.Name)
	})
}

func TestDelete(t *testing.T) {
	repo := repository.NewInmemRepository()
	presenter := presenter.NewUserPresenter()
	validator := validator.NewUserValidator(repo)

	m := interactor.NewUserInteractor(repo, presenter, test.NewFakeService(), validator)
	u := user.NewFixtureUser()
	u1, err := m.Create(u)

	err = m.Delete(u1)
	assert.Equal(t, nil, err)
	err = m.Delete(u1)
	assert.Equal(t, exception.ErrNotFound, err)
}

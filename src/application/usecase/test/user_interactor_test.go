package test

import (
	"github.com/jahs/clinic-backend/src/application/usecase"
	"github.com/jahs/clinic-backend/src/domain/exception"
	"github.com/jahs/clinic-backend/src/domain/service"
	"github.com/jahs/clinic-backend/src/infrastructure/persistence"
	"github.com/jahs/clinic-backend/src/infrastructure/persistence/test"
	presenter2 "github.com/jahs/clinic-backend/src/infrastructure/presenter"
	"github.com/jahs/clinic-backend/src/shared/password/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
	repo := persistence.NewInmemUserRepository()
	presenter := presenter2.NewUserPresenter()
	validator := service.NewUserValidator(repo)

	m := usecase.NewUserUseCase(repo, presenter, test.NewFakeService(), validator)
	u := user.NewFixtureUser()
	id, err := m.Create(u)

	assert.Nil(t, err)
	assert.Equal(t, u.ID, id)
	assert.False(t, u.CreatedAt.IsZero())
	assert.True(t, u.UpdatedAt.IsZero())
}

func Test_SearchAndFind(t *testing.T) {
	repo := persistence.NewInmemUserRepository()
	presenter := presenter2.NewUserPresenter()
	validator := service.NewUserValidator(repo)

	m := usecase.NewUserUseCase(repo, presenter, test.NewFakeService(), validator)
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
	repo := persistence.NewInmemUserRepository()
	presenter := presenter2.NewUserPresenter()
	validator := service.NewUserValidator(repo)

	m := usecase.NewUserUseCase(repo, presenter, test.NewFakeService(), validator)
	u := user.NewFixtureUser()
	u1, err := m.Create(u)

	err = m.Delete(u1)
	assert.Equal(t, nil, err)
	err = m.Delete(u1)
	assert.Equal(t, exception.ErrNotFound, err)
}

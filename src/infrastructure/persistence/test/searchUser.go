package user

import (
	"github.com/jahs/clinic-backend/src/domain"
)

func NewFixtureSearchUser() *domain.User {
	return &domain.User{
		Email:     "jahs",
		Name:      "ahs",
		Rol:       "dm",
	}
}

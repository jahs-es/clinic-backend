package user

import (
	"github.com/jahs/clinic-backend/domain/model"
)

func NewFixtureSearchUser() *model.User {
	return &model.User{
		Email:     "jahs",
		Name:      "ahs",
		Rol:       "dm",
	}
}

package user

import (
	"github.com/jahs/clinic-backend/adapter/entity"
	"time"

	"github.com/jahs/clinic-backend/domain/model"
)

func NewFixtureUser() *model.User {
	return &model.User{
		ID:        entity.NewID(),
		Email:     "jahs.es@gmail.com",
		Password:  "123456",
		Name:      "Jahs",
		Rol:       "admin",
		CreatedAt: time.Now(),
	}
}

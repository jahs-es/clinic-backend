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

func NewFixturePatient() *model.Patient {
	return &model.Patient{
		ID:        entity.NewID(),
		Name:      "Jahs",
		Address:   "Avda jahs",
		Email:     "jahs.es@gmail.com",
		Phone:     "965845487",
		CreatedAt: time.Now(),
	}
}

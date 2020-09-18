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
		Active:    true,
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
		Active:    true,
		CreatedAt: time.Now(),
	}
}

func NewFixtureTreatment() *model.Treatment {
	return &model.Treatment{
		ID:        entity.NewID(),
		Name:      "Treatment 1",
		Active:    true,
		CreatedAt: time.Now(),
	}
}

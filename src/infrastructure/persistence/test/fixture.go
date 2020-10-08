package user

import (
	"github.com/jahs/clinic-backend/src/domain"
	"github.com/jahs/clinic-backend/src/infrastructure/presenter/dto"
	"github.com/jahs/clinic-backend/src/shared/entity"
	"time"
)

func NewFixtureUser() *domain.User {
	return &domain.User{
		ID:        entity.NewID(),
		Email:     "jahs.es@gmail.com",
		Password:  "123456",
		Name:      "Jahs",
		Rol:       "admin",
		Active:    true,
		CreatedAt: time.Now(),
	}
}

func NewFixturePatient() *domain.Patient {
	return &domain.Patient{
		ID:        entity.NewID(),
		Name:      "Jahs",
		Address:   "Avda jahs",
		Email:     "jahs.es@gmail.com",
		Phone:     "965845487",
		Active:    true,
		CreatedAt: time.Now(),
	}
}

func NewFixtureTreatment() *domain.Treatment {
	return &domain.Treatment{
		ID:        entity.NewID(),
		Name:      "Treatment 1",
		Active:    true,
		CreatedAt: time.Now(),
	}
}

func NewFixturePatientTreatment() *dto.PatientTreatmentDTO {
	return &dto.PatientTreatmentDTO{
		ID:          entity.NewID(),
		PatientId:   entity.NewID(),
		Patient:     "Paciente",
		TreatmentId: entity.NewID(),
		Treatment:   "Tratamiento",
		Active:      true,
	}
}

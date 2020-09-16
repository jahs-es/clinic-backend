package presenter

import (
	"github.com/jahs/clinic-backend/adapter/entity"
)

type PatientDTO struct {
	ID        entity.ID
	Name      string
	Address   string
	Email     string
	Phone     string
	Active    bool
}

package presenter

import (
	"github.com/jahs/clinic-backend/adapter/entity"
)

type TreatmentDTO struct {
	ID        entity.ID
	Name      string
	Active    bool
}

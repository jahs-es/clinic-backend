package presenter

import (
	"github.com/jahs/clinic-backend/adapter/entity"
)

type PatientTreatmentDTO struct {
	ID          entity.ID
	PatientId   entity.ID
	Patient     string
	TreatmentId entity.ID
	Treatment   string
	Detail      string
	Active      bool
}

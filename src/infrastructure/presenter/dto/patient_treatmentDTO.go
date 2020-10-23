package dto

import (
	"github.com/jahs/clinic-backend/src/shared/entity"
)

type PatientTreatmentDTO struct {
	ID          entity.ID `json:"id"`
	PatientId   entity.ID `json:"patient_id"`
	Patient     string    `json:"patient"`
	TreatmentId entity.ID `json:"treatment_id"`
	Treatment   string    `json:"treatment"`
	Detail      string    `json:"detail"`
	Active      bool      `json:"active"`
}

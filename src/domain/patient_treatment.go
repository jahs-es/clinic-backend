package domain

import (
	"github.com/jahs/clinic-backend/src/shared/entity"
	"time"
)

type PatientTreatment struct {
	ID          entity.ID
	PatientId   entity.ID
	TreatmentId entity.ID
	Detail      string
	Active      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedBy   string
	UpdatedBy   string
}

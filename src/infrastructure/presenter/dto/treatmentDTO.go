package dto

import (
	"github.com/jahs/clinic-backend/src/shared/entity"
)

type TreatmentDTO struct {
	ID     entity.ID `json:"id"`
	Name   string    `json:"name"`
	Active bool      `json:"active"`
}

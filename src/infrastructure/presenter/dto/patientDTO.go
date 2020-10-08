package dto

import (
	"github.com/jahs/clinic-backend/src/shared/entity"
)

type PatientDTO struct {
	ID      entity.ID `json:"id"`
	AvatarPath    string    `json:"avatar_path"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
	Email   string    `json:"email"`
	Phone   string    `json:"phone"`
	Active  bool      `json:"active"`
}

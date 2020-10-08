package dto

import "github.com/jahs/clinic-backend/src/shared/entity"

type UserDTO struct {
	ID     entity.ID `json:"id"`
	Email  string    `json:"email"`
	Name   string    `json:"name"`
	Rol    string    `json:"rol"`
	Active bool      `json:"active"`
}

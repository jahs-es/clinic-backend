package presenter

import "github.com/jahs/clinic-backend/adapter/entity"

type UserDTO struct {
	ID        entity.ID
	Email     string
	Name      string
	Rol       string
	Active    bool
}

package model

import (
	"github.com/jahs/clinic-backend/adapter/entity"
	"time"
)

type User struct {
	ID        entity.ID
	Email     string
	Password  string
	Name      string
	Rol       string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

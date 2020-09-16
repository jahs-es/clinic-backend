package model

import (
	"github.com/jahs/clinic-backend/adapter/entity"
	"time"
)

type Patient struct {
	ID        entity.ID
	Name      string
	Address   string
	Email     string
	Phone     string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

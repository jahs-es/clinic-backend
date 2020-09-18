package model

import (
	"github.com/jahs/clinic-backend/adapter/entity"
	"time"
)

type Treatment struct {
	ID        entity.ID
	Name      string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy string
	UpdatedBy string
}

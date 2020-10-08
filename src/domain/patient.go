package domain

import (
	"github.com/jahs/clinic-backend/src/shared/entity"
	"time"
)

type Patient struct {
	ID        entity.ID
	AvatarPath string
	Name      string
	Address   string
	Email     string
	Phone     string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy string
	UpdatedBy string
}

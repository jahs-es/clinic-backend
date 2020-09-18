package repository

import (
	"github.com/jahs/clinic-backend/adapter/entity"
	"github.com/jahs/clinic-backend/domain/model"
)

type TreatmentRepository interface {
	Get(id entity.ID) (*model.Treatment, error)
	GetByName(name string) (*model.Treatment, error)
	Search(e *model.Treatment) ([]*model.Treatment, error)
	Create(e *model.Treatment) (entity.ID, error)
	Update(e *model.Treatment) error
	Delete(id entity.ID) error
}


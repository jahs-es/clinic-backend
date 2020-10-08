package _interface

import (
	"github.com/jahs/clinic-backend/src/domain"
	"github.com/jahs/clinic-backend/src/shared/entity"
)

type ITreatmentRepository interface {
	Get(id entity.ID) (*domain.Treatment, error)
	GetByName(name string) (*domain.Treatment, error)
	Search(e *domain.Treatment) ([]*domain.Treatment, error)
	Create(e *domain.Treatment) (entity.ID, error)
	Update(e *domain.Treatment) error
	Delete(id entity.ID) error
}


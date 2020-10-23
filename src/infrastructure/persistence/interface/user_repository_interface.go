package _interface

import (
	"github.com/jahs/clinic-backend/src/domain"
	"github.com/jahs/clinic-backend/src/shared/entity"
)

type IUserRepository interface {
	Get(id entity.ID) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	Search(e *domain.User) ([]*domain.User, error)
	Create(e *domain.User) (entity.ID, error)
	Update(e *domain.User) error
	Delete(id entity.ID) error
}


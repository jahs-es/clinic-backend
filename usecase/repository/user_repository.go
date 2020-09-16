package repository

import (
	"github.com/jahs/clinic-backend/adapter/entity"
	"github.com/jahs/clinic-backend/domain/model"
)

type UserRepository interface {
	Get(id entity.ID) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Search(e *model.User) ([]*model.User, error)
	Create(e *model.User) (entity.ID, error)
	Update(e *model.User) error
	Delete(id entity.ID) error
}


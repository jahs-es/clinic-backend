package service

import (
	"fmt"
	"github.com/jahs/clinic-backend/src/domain/exception"
	"github.com/jahs/clinic-backend/src/infrastructure/persistence/interface"
)

type userValidator struct {
	UserRepository _interface.IUserRepository
}

type UserValidator interface {
	UserNotExist(email string) error
}

func NewUserValidator(r _interface.IUserRepository) UserValidator {
	return &userValidator{r}
}

func (uv *userValidator) UserNotExist(email string) error {
	data, err := uv.UserRepository.GetByEmail(email)

	if err != nil {
		return err
	}

	fmt.Println(data)

	if data.Email != "" {
		return exception.ErrAlreadyExist
	}

	return nil
}

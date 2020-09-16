package validator

import (
	"fmt"
	"github.com/jahs/clinic-backend/usecase/exception"
	"github.com/jahs/clinic-backend/usecase/repository"
)

type userValidator struct {
	UserRepository repository.UserRepository
}

type UserValidator interface {
	UserNotExist(email string) error
}

func NewUserValidator(r repository.UserRepository) UserValidator {
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

package interactor

import (
	"fmt"
	"github.com/jahs/clinic-backend/adapter/auth"
	"github.com/jahs/clinic-backend/adapter/entity"
	"github.com/jahs/clinic-backend/adapter/password"
	presenterDTO "github.com/jahs/clinic-backend/adapter/presenter"
	"github.com/jahs/clinic-backend/domain/model"
	"github.com/jahs/clinic-backend/usecase/presenter"
	"github.com/jahs/clinic-backend/usecase/repository"
	"github.com/jahs/clinic-backend/usecase/validator"
	"time"
)

type userInteractor struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
	Service        password.Service
	Validator        validator.UserValidator
}

type UserInteractor interface {
	Login(email string, password string) (*presenterDTO.Token, error)
	Create(e *model.User) (entity.ID, error)
	Get(id entity.ID) (*presenterDTO.UserDTO, error)
	Delete(id entity.ID) error
	Find(e *model.User) ([]*presenterDTO.UserDTO, error)
}

func NewUserInteractor(r repository.UserRepository, p presenter.UserPresenter, s password.Service, v validator.UserValidator) *userInteractor {
	return &userInteractor{r, p, s, v}
}

func (us *userInteractor) Login(email string, password string) (*presenterDTO.Token, error) {
	data, err := us.UserRepository.GetByEmail(email)

	if data == nil {
		return nil, err
	}

	toJ := &model.User{
		ID:       data.ID,
		Password: data.Password,
	}

	err = us.Service.Compare(toJ.Password, password)

	if err != nil {
		return nil, err
	}

	token, err := auth.CreateToken(toJ.ID)

	tokenDTO := &presenterDTO.Token{
		AccessToken: token,
	}

	return tokenDTO, err
}

func (us *userInteractor) Create(e *model.User) (entity.ID, error) {
	err := us.Validator.UserNotExist(e.Email)
	if err != nil {
		return e.ID, err
	}

	e.ID = entity.NewID()
	e.CreatedAt = time.Now()
	pwd, err := us.Service.Generate(e.Password)
	if err != nil {
		return e.ID, err
	}

	e.Password = pwd

	return us.UserRepository.Create(e)
}

func (us *userInteractor) Get(id entity.ID) (*presenterDTO.UserDTO, error) {
	u, err := us.UserRepository.Get(id)

	if err != nil {
		return nil, err
	}

	fmt.Println(u)

	return &presenterDTO.UserDTO{
		ID: u.ID,
		Email:  u.Email,
		Name:   u.Name,
		Rol:    u.Rol,
		Active: u.Active,
	}, nil
}

func (us *userInteractor) Delete(id entity.ID) error {
	return us.UserRepository.Delete(id)
}

func (us *userInteractor) Find(e *model.User) ([]*presenterDTO.UserDTO, error) {
	u, err := us.UserRepository.Search(e)

	if err != nil {
		return nil, err
	}
	return us.UserPresenter.Response(u), nil
}

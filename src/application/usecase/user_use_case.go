package usecase

import (
	"fmt"
	"github.com/jahs/clinic-backend/src/domain"
	"github.com/jahs/clinic-backend/src/domain/service"
	"github.com/jahs/clinic-backend/src/infrastructure/persistence/interface"
	presenter2 "github.com/jahs/clinic-backend/src/infrastructure/presenter"
	"github.com/jahs/clinic-backend/src/infrastructure/presenter/dto"
	"github.com/jahs/clinic-backend/src/shared/auth"
	"github.com/jahs/clinic-backend/src/shared/entity"
	"github.com/jahs/clinic-backend/src/shared/password"
	"time"
)

type userUseCase struct {
	UserRepository _interface.IUserRepository
	UserPresenter  presenter2.IUserPresenter
	Service        password.Service
	Validator      service.UserValidator
}

type IUserUseCase interface {
	Login(email string, password string) (*dto.TokenDTO, error)
	Create(e *domain.User) (entity.ID, error)
	Get(id entity.ID) (*dto.UserDTO, error)
	Delete(id entity.ID) error
	Find(e *domain.User) ([]*dto.UserDTO, error)
}

func NewUserUseCase(r _interface.IUserRepository, p presenter2.IUserPresenter, s password.Service, v service.UserValidator) *userUseCase {
	return &userUseCase{r, p, s, v}
}

func (us *userUseCase) Login(email string, password string) (*dto.TokenDTO, error) {
	user, err := us.UserRepository.GetByEmail(email)

	if user == nil {
		return nil, err
	}

	err = us.Service.Compare(user.Password, password)

	if err != nil {
		return nil, err
	}

	createdToken, err := auth.CreateToken(user.ID)

	tokenDTO := &dto.TokenDTO{
		Id: user.ID,
		Email: user.Email,
		AvatarPath: "https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50.jpg",
		Token: createdToken,
	}

	return tokenDTO, err
}

func (us *userUseCase) Create(e *domain.User) (entity.ID, error) {
	err := us.Validator.UserNotExist(e.Email)
	if err != nil {
		return e.ID, err
	}

	//e.ID = entity.NewID()
	e.CreatedAt = time.Now()
	pwd, err := us.Service.Generate(e.Password)
	if err != nil {
		return e.ID, err
	}

	e.Password = pwd

	return us.UserRepository.Create(e)
}

func (us *userUseCase) Get(id entity.ID) (*dto.UserDTO, error) {
	u, err := us.UserRepository.Get(id)

	if err != nil {
		return nil, err
	}

	fmt.Println(u)

	return &dto.UserDTO{
		ID: u.ID,
		Email:  u.Email,
		Name:   u.Name,
		Rol:    u.Rol,
		Active: u.Active,
	}, nil
}

func (us *userUseCase) Delete(id entity.ID) error {
	return us.UserRepository.Delete(id)
}

func (us *userUseCase) Find(e *domain.User) ([]*dto.UserDTO, error) {
	u, err := us.UserRepository.Search(e)

	if err != nil {
		return nil, err
	}
	return us.UserPresenter.Response(u), nil
}

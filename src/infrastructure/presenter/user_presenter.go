package presenter

import (
	"github.com/jahs/clinic-backend/src/domain"
	"github.com/jahs/clinic-backend/src/infrastructure/presenter/dto"
)

type userPresenter struct {
}

type IUserPresenter interface {
	Response(us []*domain.User) []*dto.UserDTO
}

func NewUserPresenter() IUserPresenter {
	return &userPresenter{}
}

func (up *userPresenter) Response(us []*domain.User) []*dto.UserDTO {
	userDTOList := make([]*dto.UserDTO, 0)

	for _, d := range us {
		userDTOList = append(userDTOList, &dto.UserDTO{
			ID:     d.ID,
			Email:  d.Email,
			Name:   d.Name,
			Rol:    d.Rol,
			Active: d.Active,
		})
	}
	return userDTOList
}

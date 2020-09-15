package presenter

import (
	"github.com/jahs/clinic-backend/domain/model"
)

type userPresenter struct {
}

type UserPresenter interface {
	Response(us []*model.User) []*UserDTO
}

func NewUserPresenter() UserPresenter {
	return &userPresenter{}
}

func (up *userPresenter) Response(us []*model.User) []*UserDTO {
	var userDTOList []*UserDTO
	for _, d := range us {
		userDTOList = append(userDTOList, &UserDTO{
			ID:     d.ID,
			Email:  d.Email,
			Name:   d.Name,
			Rol:    d.Rol,
			Active: d.Active,
		})
	}
	return userDTOList
}

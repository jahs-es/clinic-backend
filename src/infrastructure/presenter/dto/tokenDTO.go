package dto

import "github.com/jahs/clinic-backend/src/shared/entity"

type TokenDTO struct {
	Id         entity.ID `json:"id"`
	AvatarPath string    `json:"avatar_path"`
	Email      string    `json:"email"`
	Token      string    `json:"token"`
}

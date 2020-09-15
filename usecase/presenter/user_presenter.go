package presenter

import (
	"github.com/jahs/clinic-backend/adapter/presenter"
	"github.com/jahs/clinic-backend/domain/model"
)

type UserPresenter interface {
	Response(u []*model.User) []*presenter.UserDTO
}

package registry

import (
	"github.com/jahs/clinic-backend/adapter/controller"
	ip "github.com/jahs/clinic-backend/adapter/presenter"
	ir "github.com/jahs/clinic-backend/adapter/repository"
	"github.com/jahs/clinic-backend/usecase/interactor"
	up "github.com/jahs/clinic-backend/usecase/presenter"
	ur "github.com/jahs/clinic-backend/usecase/repository"
	s "github.com/jahs/clinic-backend/adapter/password"
	"github.com/jahs/clinic-backend/usecase/validator"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter(), s.NewService(), r.NewUserValidator())
}

func (r *registry) NewUserValidator() validator.UserValidator {
	return validator.NewUserValidator(r.NewUserRepository())
}

func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewMySQLUserRepository(r.db)
}

func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}

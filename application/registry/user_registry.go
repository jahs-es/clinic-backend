package registry

import (
	"database/sql"
	"github.com/jahs/clinic-backend/adapter/controller"
	ip "github.com/jahs/clinic-backend/adapter/presenter"
	ir "github.com/jahs/clinic-backend/adapter/repository"
	"github.com/jahs/clinic-backend/usecase/interactor"
	up "github.com/jahs/clinic-backend/usecase/presenter"
	ur "github.com/jahs/clinic-backend/usecase/repository"
	s "github.com/jahs/clinic-backend/adapter/password"
	"github.com/jahs/clinic-backend/usecase/validator"
)

type userRegistry struct {
	db *sql.DB
}

func NewUserRegistry(db *sql.DB)  *userRegistry {
	return &userRegistry{db}
}

func (r *userRegistry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *userRegistry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter(), s.NewService(), r.NewUserValidator())
}

func (r *userRegistry) NewUserValidator() validator.UserValidator {
	return validator.NewUserValidator(r.NewUserRepository())
}

func (r *userRegistry) NewUserRepository() ur.UserRepository {
	return ir.NewMySQLUserRepository(r.db)
}

func (r *userRegistry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}

package registry

import (
	"database/sql"
	"github.com/jahs/clinic-backend/src/application/usecase"
	"github.com/jahs/clinic-backend/src/domain/service"
	controller2 "github.com/jahs/clinic-backend/src/infrastructure/controller"
	"github.com/jahs/clinic-backend/src/infrastructure/persistence"
	"github.com/jahs/clinic-backend/src/infrastructure/persistence/interface"
	"github.com/jahs/clinic-backend/src/infrastructure/presenter"
	s "github.com/jahs/clinic-backend/src/shared/password"
)

type userRegistry struct {
	db *sql.DB
}

func NewUserRegistry(db *sql.DB)  *userRegistry {
	return &userRegistry{db}
}

func (r *userRegistry) NewUserController() controller2.UserController {
	return controller2.NewUserController(r.NewUserInteractor())
}

func (r *userRegistry) NewUserInteractor() usecase.IUserUseCase {
	return usecase.NewUserUseCase(r.NewUserRepository(), r.NewUserPresenter(), s.NewService(), r.NewUserValidator())
}

func (r *userRegistry) NewUserValidator() service.UserValidator {
	return service.NewUserValidator(r.NewUserRepository())
}

func (r *userRegistry) NewUserRepository() _interface.IUserRepository {
	return persistence.NewMySQLUserRepository(r.db)
}

func (r *userRegistry) NewUserPresenter() presenter.IUserPresenter {
	return presenter.NewUserPresenter()
}

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

func (r *registry) NewPatientController() controller.PatientController {
	return controller.NewPatientController(r.NewPatientInteractor())
}

func (r *registry) NewPatientInteractor() interactor.PatientInteractor {
	return interactor.NewPatientInteractor(r.NewPatientRepository(), r.NewPatientPresenter(), s.NewService(), r.NewPatientValidator())
}

func (r *registry) NewPatientValidator() validator.PatientValidator {
	return validator.NewPatientValidator(r.NewPatientRepository())
}

func (r *registry) NewPatientRepository() ur.PatientRepository {
	return ir.NewMySQLPatientRepository(r.db)
}

func (r *registry) NewPatientPresenter() up.PatientPresenter {
	return ip.NewPatientPresenter()
}

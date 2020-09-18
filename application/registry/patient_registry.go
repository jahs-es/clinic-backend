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

type patientRegistry struct {
	db *sql.DB
}

func NewPatientRegistry(db *sql.DB) *patientRegistry {
	return &patientRegistry{db}
}

func (r *patientRegistry) NewPatientController() controller.PatientController {
	return controller.NewPatientController(r.NewPatientInteractor())
}

func (r *patientRegistry) NewPatientInteractor() interactor.PatientInteractor {
	return interactor.NewPatientInteractor(r.NewPatientRepository(), r.NewPatientPresenter(), s.NewService(), r.NewPatientValidator())
}

func (r *patientRegistry) NewPatientValidator() validator.PatientValidator {
	return validator.NewPatientValidator(r.NewPatientRepository())
}

func (r *patientRegistry) NewPatientRepository() ur.PatientRepository {
	return ir.NewMySQLPatientRepository(r.db)
}

func (r *patientRegistry) NewPatientPresenter() up.PatientPresenter {
	return ip.NewPatientPresenter()
}

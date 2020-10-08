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

type patientRegistry struct {
	db *sql.DB
}

func NewPatientRegistry(db *sql.DB) *patientRegistry {
	return &patientRegistry{db}
}

func (r *patientRegistry) NewPatientController() controller2.PatientController {
	return controller2.NewPatientController(r.NewPatientInteractor())
}

func (r *patientRegistry) NewPatientInteractor() usecase.IPatientUseCase {
	return usecase.NewPatientInteractor(r.NewPatientRepository(), r.NewPatientPresenter(), s.NewService(), r.NewPatientValidator())
}

func (r *patientRegistry) NewPatientValidator() service.PatientValidator {
	return service.NewPatientValidator(r.NewPatientRepository())
}

func (r *patientRegistry) NewPatientRepository() _interface.IPatientRepository {
	return persistence.NewMySQLPatientRepository(r.db)
}

func (r *patientRegistry) NewPatientPresenter() presenter.IPatientPresenter {
	return presenter.NewPatientPresenter()
}

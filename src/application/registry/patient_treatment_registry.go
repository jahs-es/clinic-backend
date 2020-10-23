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

type patientTreatmentRegistry struct {
	db *sql.DB
}

func NewPatientTreatmentRegistry(db *sql.DB) *patientTreatmentRegistry {
	return &patientTreatmentRegistry{db}
}

func (r *patientTreatmentRegistry) NewPatientTreatmentController() controller2.PatientTreatmentController {
	return controller2.NewPatientTreatmentController(r.NewPatientTreatmentInteractor())
}

func (r *patientTreatmentRegistry) NewPatientTreatmentInteractor() usecase.IPatientTreatmentUseCase {
	return usecase.NewPatientTreatmentInteractor(r.NewPatientTreatmentRepository(), r.NewPatientTreatmentPresenter(), s.NewService(), r.NewPatientTreatmentValidator())
}

func (r *patientTreatmentRegistry) NewPatientTreatmentValidator() service.PatientTreatmentValidator {
	return service.NewPatientTreatmentValidator(r.NewPatientTreatmentRepository())
}

func (r *patientTreatmentRegistry) NewPatientTreatmentRepository() _interface.IPatientTreatmentRepository {
	return persistence.NewMySQLPatientTreatmentRepository(r.db)
}

func (r *patientTreatmentRegistry) NewPatientTreatmentPresenter() presenter.IPatientTreatmentPresenter {
	return presenter.NewPatientTreatmentPresenter()
}

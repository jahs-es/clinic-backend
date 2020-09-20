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

type patientTreatmentRegistry struct {
	db *sql.DB
}

func NewPatientTreatmentRegistry(db *sql.DB) *patientTreatmentRegistry {
	return &patientTreatmentRegistry{db}
}

func (r *patientTreatmentRegistry) NewPatientTreatmentController() controller.PatientTreatmentController {
	return controller.NewPatientTreatmentController(r.NewPatientTreatmentInteractor())
}

func (r *patientTreatmentRegistry) NewPatientTreatmentInteractor() interactor.PatientTreatmentInteractor {
	return interactor.NewPatientTreatmentInteractor(r.NewPatientTreatmentRepository(), r.NewPatientTreatmentPresenter(), s.NewService(), r.NewPatientTreatmentValidator())
}

func (r *patientTreatmentRegistry) NewPatientTreatmentValidator() validator.PatientTreatmentValidator {
	return validator.NewPatientTreatmentValidator(r.NewPatientTreatmentRepository())
}

func (r *patientTreatmentRegistry) NewPatientTreatmentRepository() ur.PatientTreatmentRepository {
	return ir.NewMySQLPatientTreatmentRepository(r.db)
}

func (r *patientTreatmentRegistry) NewPatientTreatmentPresenter() up.PatientTreatmentPresenter {
	return ip.NewPatientTreatmentPresenter()
}

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

type treatmentRegistry struct {
	db *sql.DB
}

func NewTreatmentRegistry(db *sql.DB) *treatmentRegistry {
	return &treatmentRegistry{db}
}

func (r *treatmentRegistry) NewTreatmentController() controller.TreatmentController {
	return controller.NewTreatmentController(r.NewTreatmentInteractor())
}

func (r *treatmentRegistry) NewTreatmentInteractor() interactor.TreatmentInteractor {
	return interactor.NewTreatmentInteractor(r.NewTreatmentRepository(), r.NewTreatmentPresenter(), s.NewService(), r.NewTreatmentValidator())
}

func (r *treatmentRegistry) NewTreatmentValidator() validator.TreatmentValidator {
	return validator.NewTreatmentValidator(r.NewTreatmentRepository())
}

func (r *treatmentRegistry) NewTreatmentRepository() ur.TreatmentRepository {
	return ir.NewMySQLTreatmentRepository(r.db)
}

func (r *treatmentRegistry) NewTreatmentPresenter() up.TreatmentPresenter {
	return ip.NewTreatmentPresenter()
}

package registry

import (
	"database/sql"
	"github.com/jahs/clinic-backend/src/application/usecase"
	"github.com/jahs/clinic-backend/src/domain/service"
	"github.com/jahs/clinic-backend/src/infrastructure/controller"
	"github.com/jahs/clinic-backend/src/infrastructure/persistence"
	"github.com/jahs/clinic-backend/src/infrastructure/persistence/interface"
	"github.com/jahs/clinic-backend/src/infrastructure/presenter"
	s "github.com/jahs/clinic-backend/src/shared/password"
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

func (r *treatmentRegistry) NewTreatmentInteractor() usecase.ITreatmentUseCase {
	return usecase.NewTreatmentUseCase(r.NewTreatmentRepository(), r.NewTreatmentPresenter(), s.NewService(), r.NewTreatmentValidator())
}

func (r *treatmentRegistry) NewTreatmentValidator() service.TreatmentValidator {
	return service.NewTreatmentValidator(r.NewTreatmentRepository())
}

func (r *treatmentRegistry) NewTreatmentRepository() _interface.ITreatmentRepository {
	return persistence.NewMySQLTreatmentRepository(r.db)
}

func (r *treatmentRegistry) NewTreatmentPresenter() presenter.ITreatmentPresenter {
	return presenter.NewTreatmentPresenter()
}

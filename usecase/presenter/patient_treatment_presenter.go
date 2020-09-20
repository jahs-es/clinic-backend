package presenter

import (
	"github.com/jahs/clinic-backend/adapter/presenter"
	"github.com/jahs/clinic-backend/domain/model"
)

type PatientTreatmentPresenter interface {
	Response(u []*model.PatientTreatment) []*presenter.PatientTreatmentDTO
}

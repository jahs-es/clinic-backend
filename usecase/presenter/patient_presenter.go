package presenter

import (
	"github.com/jahs/clinic-backend/adapter/presenter"
	"github.com/jahs/clinic-backend/domain/model"
)

type PatientPresenter interface {
	Response(u []*model.Patient) []*presenter.PatientDTO
}

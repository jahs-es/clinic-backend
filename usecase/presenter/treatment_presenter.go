package presenter

import (
	"github.com/jahs/clinic-backend/adapter/presenter"
	"github.com/jahs/clinic-backend/domain/model"
)

type TreatmentPresenter interface {
	Response(u []*model.Treatment) []*presenter.TreatmentDTO
}

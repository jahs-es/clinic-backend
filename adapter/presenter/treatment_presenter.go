package presenter

import (
	"github.com/jahs/clinic-backend/domain/model"
)

type treatmentPresenter struct {
}

type TreatmentPresenter interface {
	Response(us []*model.Treatment) []*TreatmentDTO
}

func NewTreatmentPresenter() TreatmentPresenter {
	return &treatmentPresenter{}
}

func (up *treatmentPresenter) Response(us []*model.Treatment) []*TreatmentDTO {
	var treatmentDTOList []*TreatmentDTO
	for _, d := range us {
		treatmentDTOList = append(treatmentDTOList, &TreatmentDTO{
			ID:      d.ID,
			Name:    d.Name,
			Active:  d.Active,
		})
	}
	return treatmentDTOList
}

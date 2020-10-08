package presenter

import (
	"github.com/jahs/clinic-backend/src/domain"
	"github.com/jahs/clinic-backend/src/infrastructure/presenter/dto"
)

type treatmentPresenter struct {
}

type ITreatmentPresenter interface {
	Response(us []*domain.Treatment) []*dto.TreatmentDTO
}

func NewTreatmentPresenter() ITreatmentPresenter {
	return &treatmentPresenter{}
}

func (up *treatmentPresenter) Response(us []*domain.Treatment) []*dto.TreatmentDTO {
	treatmentDTOList := make([]*dto.TreatmentDTO,0)

	for _, d := range us {
		treatmentDTOList = append(treatmentDTOList, &dto.TreatmentDTO{
			ID:      d.ID,
			Name:    d.Name,
			Active:  d.Active,
		})
	}
	return treatmentDTOList
}

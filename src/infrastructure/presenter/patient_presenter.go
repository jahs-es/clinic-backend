package presenter

import (
	"github.com/jahs/clinic-backend/src/domain"
	"github.com/jahs/clinic-backend/src/infrastructure/presenter/dto"
)

type patientPresenter struct {
}

type IPatientPresenter interface {
	Response(us []*domain.Patient) []*dto.PatientDTO
}

func NewPatientPresenter() IPatientPresenter {
	return &patientPresenter{}
}

func (up *patientPresenter) Response(us []*domain.Patient) []*dto.PatientDTO {
	patientDTOList := make([]*dto.PatientDTO,0)

	for _, d := range us {
		patientDTOList = append(patientDTOList, &dto.PatientDTO{
			ID:      d.ID,
			AvatarPath:    d.AvatarPath,
			Name:    d.Name,
			Address: d.Address,
			Email:   d.Email,
			Phone:   d.Phone,
			Active:  d.Active,
		})
	}
	return patientDTOList
}

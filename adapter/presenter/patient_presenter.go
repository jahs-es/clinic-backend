package presenter

import (
	"github.com/jahs/clinic-backend/domain/model"
)

type patientPresenter struct {
}

type PatientPresenter interface {
	Response(us []*model.Patient) []*PatientDTO
}

func NewPatientPresenter() PatientPresenter {
	return &patientPresenter{}
}

func (up *patientPresenter) Response(us []*model.Patient) []*PatientDTO {
	var patientDTOList []*PatientDTO
	for _, d := range us {
		patientDTOList = append(patientDTOList, &PatientDTO{
			ID:      d.ID,
			Name:    d.Name,
			Address: d.Address,
			Email:   d.Email,
			Phone:   d.Phone,
			Active:  d.Active,
		})
	}
	return patientDTOList
}

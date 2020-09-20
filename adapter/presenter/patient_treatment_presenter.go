package presenter

import (
	"github.com/jahs/clinic-backend/domain/model"
)

type patientTreatmentTreatmentPresenter struct {
}

type PatientTreatmentPresenter interface {
	Response(us []*model.PatientTreatment) []*PatientTreatmentDTO
}

func NewPatientTreatmentPresenter() PatientTreatmentPresenter {
	return &patientTreatmentTreatmentPresenter{}
}

func (up *patientTreatmentTreatmentPresenter) Response(us []*model.PatientTreatment) []*PatientTreatmentDTO {
	patientTreatmentDTOList := make([]*PatientTreatmentDTO,0)

	for _, d := range us {
		patientTreatmentDTOList = append(patientTreatmentDTOList, &PatientTreatmentDTO{
			ID:          d.ID,
			PatientId:   d.PatientId,
			TreatmentId: d.TreatmentId,
			Detail:      d.Detail,
			Active:      d.Active,
		})
	}
	return patientTreatmentDTOList
}

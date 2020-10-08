package presenter

import (
	"github.com/jahs/clinic-backend/src/domain"
	"github.com/jahs/clinic-backend/src/infrastructure/presenter/dto"
)

type patientTreatmentTreatmentPresenter struct {
}

type IPatientTreatmentPresenter interface {
	Response(us []*domain.PatientTreatment) []*dto.PatientTreatmentDTO
}

func NewPatientTreatmentPresenter() IPatientTreatmentPresenter {
	return &patientTreatmentTreatmentPresenter{}
}

func (up *patientTreatmentTreatmentPresenter) Response(us []*domain.PatientTreatment) []*dto.PatientTreatmentDTO {
	patientTreatmentDTOList := make([]*dto.PatientTreatmentDTO,0)

	for _, d := range us {
		patientTreatmentDTOList = append(patientTreatmentDTOList, &dto.PatientTreatmentDTO{
			ID:          d.ID,
			PatientId:   d.PatientId,
			TreatmentId: d.TreatmentId,
			Detail:      d.Detail,
			Active:      d.Active,
		})
	}
	return patientTreatmentDTOList
}

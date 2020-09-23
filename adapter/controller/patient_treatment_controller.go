package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jahs/clinic-backend/adapter/auth"
	"github.com/jahs/clinic-backend/adapter/entity"
	presenterDTO "github.com/jahs/clinic-backend/adapter/presenter"
	"github.com/jahs/clinic-backend/domain/model"
	"github.com/jahs/clinic-backend/usecase/exception"
	"github.com/jahs/clinic-backend/usecase/interactor"
	"log"
	"net/http"
	"time"
)

type patientTreatmentController struct {
	patientTreatmentInteractor interactor.PatientTreatmentInteractor
}

type PatientTreatmentController interface {
	Find() http.Handler
	Get() http.Handler
	Create() http.Handler
	Update() http.Handler
	Delete() http.Handler
}

func NewPatientTreatmentController(us interactor.PatientTreatmentInteractor) *patientTreatmentController {
	return &patientTreatmentController{us}
}

func (uc *patientTreatmentController) Find() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading patientTreatments"

		detail := r.URL.Query().Get("detail")
		patient := r.URL.Query().Get("name")
		treatment := r.URL.Query().Get("rol")

		u := &presenterDTO.PatientTreatmentDTO{
			Detail:    detail,
			Patient:   patient,
			Treatment: treatment,
		}

		data, err := uc.patientTreatmentInteractor.Find(u)

		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != exception.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if err := json.NewEncoder(w).Encode(data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func (uc *patientTreatmentController) Create() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding patientTreatment"
		var input struct {
			PatientId   string `json:"patient_id"`
			TreatmentId string `json:"treatment_id"`
			Detail      string `json:"detail"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		userId, error := auth.ExtractTokenID(r)
		if error != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(errorMessage))
			return
		}

		patientId, _ := entity.StringToID(input.PatientId)
		treatmentId, _ := entity.StringToID(input.TreatmentId)

		//TODO: validate data ;)
		u := &model.PatientTreatment{
			ID:          entity.NewID(),
			PatientId:   patientId,
			TreatmentId: treatmentId,
			Detail:      input.Detail,
			CreatedAt:   time.Now(),
			CreatedBy:   userId.(string),
		}
		u.ID, err = uc.patientTreatmentInteractor.Create(u)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		toJ := &presenterDTO.PatientTreatmentDTO{
			ID:          u.ID,
			PatientId:   u.PatientId,
			TreatmentId: u.TreatmentId,
			Detail:      u.Detail,
			Active:      true,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func (uc *patientTreatmentController) Update() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding patientTreatment"
		var input struct {
			Detail string `json:"detail"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		userId, error := auth.ExtractTokenID(r)
		if error != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(errorMessage))
			return
		}

		//TODO: validate data ;)
		u := &model.PatientTreatment{
			ID:        entity.NewID(),
			Detail:    input.Detail,
			UpdatedAt: time.Now(),
			UpdatedBy: userId.(string),
		}
		err = uc.patientTreatmentInteractor.Update(u)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusNoContent)
	})
}

func (uc *patientTreatmentController) Get() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading patientTreatment"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		data, err := uc.patientTreatmentInteractor.Get(id)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if err := json.NewEncoder(w).Encode(data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func (uc *patientTreatmentController) Delete() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing patientTreatment"

		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		err = uc.patientTreatmentInteractor.Delete(id)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

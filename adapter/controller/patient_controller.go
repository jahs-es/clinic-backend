package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jahs/clinic-backend/adapter/entity"
	"github.com/jahs/clinic-backend/domain/model"
	"github.com/jahs/clinic-backend/usecase/exception"
	"github.com/jahs/clinic-backend/usecase/interactor"
	"log"
	"net/http"
	"time"
)

type patientController struct {
	patientInteractor interactor.PatientInteractor
}

type PatientController interface {
	Find() http.Handler
	Get() http.Handler
	Create() http.Handler
	Delete() http.Handler
}

func NewPatientController(us interactor.PatientInteractor) *patientController {
	return &patientController{us}
}

func (uc *patientController) Find() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading patients"
		var input struct {
			Email   string `json:"email"`
			Name    string `json:"name"`
			Address string `json:"rol"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		//TODO: validate data ;)
		u := &model.Patient{
			Email:   input.Email,
			Name:    input.Name,
			Address: input.Address,
		}

		data, err := uc.patientInteractor.Find(u)

		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != exception.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		if err := json.NewEncoder(w).Encode(data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func (uc *patientController) Create() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding patient"
		var input struct {
			Name    string `json:"name"`
			Address string `json:"address"`
			Email   string `json:"email"`
			Phone   string `json:"phone"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		//TODO: validate data ;)
		u := &model.Patient{
			ID:        entity.NewID(),
			Name:      input.Name,
			Address:   input.Address,
			Email:     input.Email,
			Phone:     input.Phone,
			CreatedAt: time.Now(),
		}
		u.ID, err = uc.patientInteractor.Create(u)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		toJ := &model.Patient{
			ID:    u.ID,
			Name:  u.Name,
			Address: u.Address,
			Email: u.Email,
			Phone: u.Phone,
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

func (uc *patientController) Get() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading patient"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		data, err := uc.patientInteractor.Get(id)
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

func (uc *patientController) Delete() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing patient"

		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		err = uc.patientInteractor.Delete(id)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

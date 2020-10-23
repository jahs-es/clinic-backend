package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jahs/clinic-backend/src/application/usecase"
	"github.com/jahs/clinic-backend/src/infrastructure/presenter/dto"
	"github.com/jahs/clinic-backend/src/shared/auth"
	"github.com/jahs/clinic-backend/src/shared/entity"

	"github.com/jahs/clinic-backend/src/domain"
	"github.com/jahs/clinic-backend/src/domain/exception"
	"log"
	"net/http"
	"time"
)

type treatmentController struct {
	treatmentInteractor usecase.ITreatmentUseCase
}

type TreatmentController interface {
	Find() http.Handler
	Get() http.Handler
	Create() http.Handler
	Update() http.Handler
	Delete() http.Handler
}

func NewTreatmentController(us usecase.ITreatmentUseCase) *treatmentController {
	return &treatmentController{us}
}

func (uc *treatmentController) Find() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading treatments"

		name := r.URL.Query().Get("name")

		u := &domain.Treatment{
			Name: name,
		}

		data, err := uc.treatmentInteractor.Find(u)

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

func (uc *treatmentController) Create() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding treatment"
		var input struct {
			ID   string `json:"id"`
			Name string `json:"name"`
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

		id, error := entity.StringToID(input.ID)
		if error != nil {
			log.Println(error)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errorMessage))
			return
		}

		//TODO: validate data ;)
		u := &domain.Treatment{
			ID:        id,
			Name:      input.Name,
			CreatedAt: time.Now(),
			CreatedBy: userId.(string),
		}
		u.ID, err = uc.treatmentInteractor.Create(u)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		toJ := &dto.TreatmentDTO{
			ID:     u.ID,
			Name:   u.Name,
			Active: true,
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

func (uc *treatmentController) Update() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding treatment"
		var input struct {
			ID      string `json:"id"`
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

		userId, error := auth.ExtractTokenID(r)
		if error != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(errorMessage))
			return
		}

		id, error := entity.StringToID(input.ID)
		if error != nil {
			log.Println(error)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errorMessage))
			return
		}

		//TODO: validate data ;)
		u := &domain.Treatment{
			ID:        id,
			Name:      input.Name,
			UpdatedAt: time.Now(),
			UpdatedBy: userId.(string),
		}
		err = uc.treatmentInteractor.Update(u)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusNoContent)
	})
}

func (uc *treatmentController) Get() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading treatment"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		data, err := uc.treatmentInteractor.Get(id)
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

func (uc *treatmentController) Delete() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing treatment"

		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		err = uc.treatmentInteractor.Delete(id)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

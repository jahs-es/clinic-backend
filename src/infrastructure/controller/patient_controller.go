package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jahs/clinic-backend/src/application/usecase"
	"github.com/jahs/clinic-backend/src/domain"
	"github.com/jahs/clinic-backend/src/domain/exception"
	"github.com/jahs/clinic-backend/src/infrastructure/presenter/dto"
	"github.com/jahs/clinic-backend/src/shared/auth"
	"github.com/jahs/clinic-backend/src/shared/entity"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type patientController struct {
	patientInteractor usecase.IPatientUseCase
}

type PatientController interface {
	Find() http.Handler
	Get() http.Handler
	Create() http.Handler
	Update() http.Handler
	Delete() http.Handler
}

func NewPatientController(us usecase.IPatientUseCase) *patientController {
	return &patientController{us}
}

func (uc *patientController) Find() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//errorMessage := "Error reading patients"

		email := r.URL.Query().Get("email")
		name := r.URL.Query().Get("name")
		address := r.URL.Query().Get("address")

		u := &domain.Patient{
			Email:   email,
			Name:    name,
			Address: address,
		}

		data, err := uc.patientInteractor.Find(u)

		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != exception.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if err := json.NewEncoder(w).Encode(data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	})
}

func (uc *patientController) Create() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding patient"
		var input struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			AvatarPath    string `json:"avatar_path"`
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

		if input.AvatarPath == "" {
			input.AvatarPath = getAvatarPath()
		}

		u := &domain.Patient{
			ID:        id,
			AvatarPath: input.AvatarPath,
			Name:      input.Name,
			Address:   input.Address,
			Email:     input.Email,
			Phone:     input.Phone,
			CreatedAt: time.Now(),
			CreatedBy: userId.(string),
		}
		u.ID, err = uc.patientInteractor.Create(u)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		toJ := &dto.PatientDTO{
			ID:      u.ID,
			AvatarPath:    u.AvatarPath,
			Name:    u.Name,
			Address: u.Address,
			Email:   u.Email,
			Phone:   u.Phone,
			Active:  true,
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

func (uc *patientController) Update() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error updating patient"
		var input struct {
			ID      string `json:"id"`
			AvatarPath    string `json:"avatar_path"`
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
		u := &domain.Patient{
			ID:        id,
			AvatarPath:    input.AvatarPath,
			Name:      input.Name,
			Address:   input.Address,
			Email:     input.Email,
			Phone:     input.Phone,
			UpdatedAt: time.Now(),
			UpdatedBy: userId.(string),
		}
		err = uc.patientInteractor.Update(u)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusNoContent)
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


func getAvatarPath() string {
	type Avatars struct {
		Height      int `json:"height"`
		Size    string    `json:"size"`
		Url    string    `json:"url"`
		Width string    `json:"width"`
	}

	type AvatarMain struct {
		Avatars      []Avatars `json:"avatars"`
	}

	url := "https://tinyfac.es/api/users"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)

	var newsList []AvatarMain
	json.Unmarshal(responseData, &newsList)

	return newsList[0].Avatars[0].Url
}


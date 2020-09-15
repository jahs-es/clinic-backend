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


type userController struct {
	userInteractor interactor.UserInteractor
}

type UserController interface {
	Login() http.Handler
	Find() http.Handler
	Get() http.Handler
	Create() http.Handler
	Delete() http.Handler
}

func NewUserController(us interactor.UserInteractor) *userController {
	return &userController{us}
}

func (uc *userController) Login() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error logging user"

		var input struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		token, err := uc.userInteractor.Login(input.Email, input.Password)

		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(errorMessage))
			return
		}

		if err := json.NewEncoder(w).Encode(token); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(errorMessage))
		}
	})
}

func (uc *userController) Find() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading users"
		var input struct {
			Email    string `json:"email"`
			Name     string `json:"name"`
			Rol      string `json:"rol"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		//TODO: validate data ;)
		u := &model.User{
			Email:     input.Email,
			Name:      input.Name,
			Rol:       input.Rol,
		}

		data, err := uc.userInteractor.Find(u)

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

func (uc *userController) Create() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding user"
		var input struct {
			Email    string `json:"email"`
			Password string `json:"password"`
			Name     string `json:"name"`
			Rol      string `json:"rol"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		//TODO: validate data ;)
		u := &model.User{
			ID:        entity.NewID(),
			Email:     input.Email,
			Password:  input.Password,
			Name:      input.Name,
			Rol:       input.Rol,
			CreatedAt: time.Now(),
		}
		u.ID, err = uc.userInteractor.Create(u)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		toJ := &model.User{
			ID:    u.ID,
			Email: u.Email,
			Name:  u.Name,
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

func (uc *userController) Get() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading user"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		data, err := uc.userInteractor.Get(id)
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

func (uc *userController) Delete() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing user"

		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		err = uc.userInteractor.Delete(id)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}


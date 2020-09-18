package controller

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jahs/clinic-backend/adapter/controller"
)

func MakeUserHandlers(r *mux.Router, n negroni.Negroni, appController controller.UserController) {
	r.Handle("/v1/user", n.With(
		negroni.Wrap(appController.Find()),
	)).Methods("GET", "OPTIONS").Name("Find")

	r.Handle("/v1/user", n.With(
		negroni.Wrap(appController.Create()),
	)).Methods("POST", "OPTIONS").Name("createUser")

	r.Handle("/v1/user/{id}", n.With(
		negroni.Wrap(appController.Get()),
	)).Methods("GET", "OPTIONS").Name("getUser")

	r.Handle("/v1/user/{id}", n.With(
		negroni.Wrap(appController.Delete()),
	)).Methods("DELETE", "OPTIONS").Name("deleteUser")
}

func MakeLoginHandlers(r *mux.Router, n negroni.Negroni, appController controller.UserController) {
	r.Handle("/v1/login", n.With(
		negroni.Wrap(appController.Login()),
	)).Methods("POST", "OPTIONS").Name("login")
}

package rest_endpoints

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	controller2 "github.com/jahs/clinic-backend/src/infrastructure/controller"
)

func MakeUserHandlers(r *mux.Router, n negroni.Negroni, appController controller2.UserController) {
	r.Handle("/api/v1/user", n.With(
		negroni.Wrap(appController.Find()),
	)).Methods("GET", "OPTIONS").Name("Find")

	r.Handle("/api/v1/user", n.With(
		negroni.Wrap(appController.Create()),
	)).Methods("POST", "OPTIONS").Name("createUser")

	r.Handle("/api/v1/user/{id}", n.With(
		negroni.Wrap(appController.Get()),
	)).Methods("GET", "OPTIONS").Name("getUser")

	r.Handle("/api/v1/user/{id}", n.With(
		negroni.Wrap(appController.Delete()),
	)).Methods("DELETE", "OPTIONS").Name("deleteUser")
}

func MakeLoginHandlers(r *mux.Router, n negroni.Negroni, appController controller2.UserController) {
	r.Handle("/api/v1/login", n.With(
		negroni.Wrap(appController.Login()),
	)).Methods("POST", "OPTIONS").Name("login")
}

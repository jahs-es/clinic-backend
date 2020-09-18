package controller

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jahs/clinic-backend/adapter/controller"
)

func MakePatientHandlers(r *mux.Router, n negroni.Negroni, appController controller.PatientController) {
	r.Handle("/v1/patient", n.With(
		negroni.Wrap(appController.Find()),
	)).Methods("GET", "OPTIONS").Name("Find")

	r.Handle("/v1/patient", n.With(
		negroni.Wrap(appController.Create()),
	)).Methods("POST", "OPTIONS").Name("createPatient")

	r.Handle("/v1/patient", n.With(
		negroni.Wrap(appController.Update()),
	)).Methods("PUT", "OPTIONS").Name("updatePatient")

	r.Handle("/v1/patient/{id}", n.With(
		negroni.Wrap(appController.Get()),
	)).Methods("GET", "OPTIONS").Name("getPatient")

	r.Handle("/v1/patient/{id}", n.With(
		negroni.Wrap(appController.Delete()),
	)).Methods("DELETE", "OPTIONS").Name("deletePatient")
}

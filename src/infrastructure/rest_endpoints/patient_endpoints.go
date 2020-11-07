package rest_endpoints

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	controller2 "github.com/jahs/clinic-backend/src/infrastructure/controller"
)

func MakePatientHandlers(r *mux.Router, n negroni.Negroni, appController controller2.PatientController) {
	r.Handle("/api/v1/patient", n.With(
		negroni.Wrap(appController.Find()),
	)).Methods("GET", "OPTIONS").Name("Find")

	r.Handle("/api/v1/patient", n.With(
		negroni.Wrap(appController.Create()),
	)).Methods("POST", "OPTIONS").Name("createPatient")

	r.Handle("/api/v1/patient", n.With(
		negroni.Wrap(appController.Update()),
	)).Methods("PUT", "OPTIONS").Name("updatePatient")

	r.Handle("/api/v1/patient/{id}", n.With(
		negroni.Wrap(appController.Get()),
	)).Methods("GET", "OPTIONS").Name("getPatient")

	r.Handle("/api/v1/patient/{id}", n.With(
		negroni.Wrap(appController.Delete()),
	)).Methods("DELETE", "OPTIONS").Name("deletePatient")
}

package controller

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jahs/clinic-backend/adapter/controller"
)

func MakePatientTreatmentHandlers(r *mux.Router, n negroni.Negroni, appController controller.PatientTreatmentController) {
	r.Handle("/v1/patient_treatment", n.With(
		negroni.Wrap(appController.Find()),
	)).Methods("GET", "OPTIONS").Name("Find")

	r.Handle("/v1/patient_treatment", n.With(
		negroni.Wrap(appController.Create()),
	)).Methods("POST", "OPTIONS").Name("createPatientTreatment")

	r.Handle("/v1/patient_treatment", n.With(
		negroni.Wrap(appController.Update()),
	)).Methods("PUT", "OPTIONS").Name("updatePatientTreatment")

	r.Handle("/v1/patient_treatment/{id}", n.With(
		negroni.Wrap(appController.Get()),
	)).Methods("GET", "OPTIONS").Name("getPatientTreatment")

	r.Handle("/v1/patient_treatment/{id}", n.With(
		negroni.Wrap(appController.Delete()),
	)).Methods("DELETE", "OPTIONS").Name("deletePatientTreatment")
}

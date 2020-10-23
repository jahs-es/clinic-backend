package rest_endpoints

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	controller2 "github.com/jahs/clinic-backend/src/infrastructure/controller"
)

func MakePatientTreatmentHandlers(r *mux.Router, n negroni.Negroni, appController controller2.PatientTreatmentController) {
	r.Handle("/v1/patient_treatment", n.With(
		negroni.Wrap(appController.Find()),
	)).Methods("GET", "OPTIONS").Name("Find")

	r.Handle("/v1/patient_treatment/{patient-id}", n.With(
		negroni.Wrap(appController.FindByPatientID()),
	)).Methods("GET", "OPTIONS").Name("getPatientTreatmentByPatientID")

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

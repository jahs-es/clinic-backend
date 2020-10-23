package rest_endpoints

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jahs/clinic-backend/src/infrastructure/controller"
)

func MakeTreatmentHandlers(r *mux.Router, n negroni.Negroni, appController controller.TreatmentController) {
	r.Handle("/v1/treatment", n.With(
		negroni.Wrap(appController.Find()),
	)).Methods("GET", "OPTIONS").Name("Find")

	r.Handle("/v1/treatment", n.With(
		negroni.Wrap(appController.Create()),
	)).Methods("POST", "OPTIONS").Name("createTreatment")

	r.Handle("/v1/treatment", n.With(
		negroni.Wrap(appController.Update()),
	)).Methods("PUT", "OPTIONS").Name("updateTreatment")

	r.Handle("/v1/treatment/{id}", n.With(
		negroni.Wrap(appController.Get()),
	)).Methods("GET", "OPTIONS").Name("getTreatment")

	r.Handle("/v1/treatment/{id}", n.With(
		negroni.Wrap(appController.Delete()),
	)).Methods("DELETE", "OPTIONS").Name("deleteTreatment")
}


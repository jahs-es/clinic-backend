package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/jahs/clinic-backend/src/application/registry"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/codegangsta/negroni"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/jahs/clinic-backend/src/application/config"
	"github.com/jahs/clinic-backend/src/infrastructure/rest_endpoints"
	"github.com/jahs/clinic-backend/src/shared/metric"
	"github.com/jahs/clinic-backend/src/shared/middleware"
)

func main() {
	db, err := openDatabase()

	userRegistry := registry.NewUserRegistry(db)
	patientRegistry := registry.NewPatientRegistry(db)
	treatmentRegistry := registry.NewTreatmentRegistry(db)
	patientTreatmentRegistry := registry.NewPatientTreatmentRegistry(db)

	metricService, err := metric.NewPrometheusService()
	if err != nil {
		log.Fatal(err.Error())
	}

	r := mux.NewRouter()

	//handlers with security
	securedHandler := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.HandlerFunc(middleware.Authentication),
		negroni.HandlerFunc(middleware.Metrics(metricService)),
		negroni.NewLogger(),
	)

	//handlers not secured
	notSecuredHandler := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.HandlerFunc(middleware.Metrics(metricService)),
		negroni.NewLogger(),
	)

	//login
	rest_endpoints.MakeLoginHandlers(r, *notSecuredHandler, userRegistry.NewUserController())

	//user
	rest_endpoints.MakeUserHandlers(r, *securedHandler, userRegistry.NewUserController())

	//patient
	rest_endpoints.MakePatientHandlers(r, *securedHandler, patientRegistry.NewPatientController())

	//treatment
	rest_endpoints.MakeTreatmentHandlers(r, *securedHandler, treatmentRegistry.NewTreatmentController())

	//patient
	rest_endpoints.MakePatientTreatmentHandlers(r, *securedHandler, patientTreatmentRegistry.NewPatientTreatmentController())

	http.Handle("/", r)
	http.Handle("/metrics", promhttp.Handler())

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(config.API_PORT),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func openDatabase() (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_DATABASE)
	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		log.Fatalf("could not connect to the MySQL database... %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("could not ping DB... %v", err)
	}

	//defer db.Close()

	//Migration
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("could not start sql migration... %v", err)
	}

	var migrationDir = flag.String("migration.files", "../migrations", "Directory where the migration files are located ?")

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", *migrationDir),
		"mysql", driver)

	if err != nil {
		log.Fatalf("migration failed... %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("An error occurred while syncing the database.. %v", err)
	}

	log.Println("Database migrated ...")

	return db, err
}

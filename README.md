# Clinic backend in Go

An clinic api created using Go following a clean code architecture approach.

- Uses MySql as database, but due to design followed is so easy to change to a different one (just implement repository interfaces and inyect new implementation). 

- Security implemented using JWT. Endpoints securized through Negroni middleware.

- Perfomance test done using k6's. Result report can be seen from `Graphana` after `InfluxDB acts` (see README in perfomance folder).
 
- Prometheus is also included in order to catch some metrics.

- All can be launched from included docker compose. 

- Initial migration adds a user with email `admin@gmail.com` and pass `admin`.  

- To test endpoints firstly call login endpoint with valid user and grab token in order to include it as bearer token in securized endpoints calls.

## Build for development

  make

## Build for production (check config files)

  make build-api-prod

## Run tests

  make test

## Creates docker

  make generate-docker

## Launch docker compose with database, graphana and prometheus

  make launch-docker

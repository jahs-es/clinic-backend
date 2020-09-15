// +build prod

package config

const (
	DB_USER                = "root"
	DB_PASSWORD            = "root001"
	DB_DATABASE            = "clinic"
	DB_HOST                = "127.0.0.1"
	API_PORT               = 8080
	PROMETHEUS_PUSHGATEWAY = "http://localhost:9091/"
)

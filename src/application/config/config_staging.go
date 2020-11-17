// +build staging

package config

const (
	DB_USER                = "root"
	DB_PASSWORD            = "root001"
	DB_DATABASE            = "clinic"
	DB_HOST                = "127.0.0.1"
	API_PORT               = 3001
	PROMETHEUS_PUSHGATEWAY = "http://localhost:9091/"
	SENDGRID_KEY		   = "your api key"
	SENDGRID_FROM		   = "your sender mail"
)

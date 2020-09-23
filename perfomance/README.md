# Clinic backend in Go - Perfomance test

Launch using K6

## Launch without Grafana 

`npm install .`

`npm run-script webpack`

`k6 run build/app.bundle.js`

Take in count entry in webpack.config.js

## Launch with InfluxDB and Grafana 

Launch `run-influx-grafana` from Makefile   

From console `k6 run scripts/script.js --out influxdb=http://localhost:8086/k6`

Open Grafana console in `http://localhost:3000` and search for `k6 perfomance test` dashboard

## More info 

https://github.com/k6io

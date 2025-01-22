# Sensor ingress


## Backend
In progress app to aggregate sensor readings into mongodb

Backend is a golang application with stepped docker build which can be done like`docker build .`

## Sensor reader
A simpleest possible datafeeder app created with python is also included [sensor-data-feeder/](here). It reads some raspberry pi CPU (package) temperatures and sends those to the backend.

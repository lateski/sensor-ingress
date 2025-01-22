# syntax=docker/dockerfile:1

FROM golang:1.18 as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

COPY configs configs

RUN go build -o /sensoringress

#deploy
FROM gcr.io/distroless/base-debian11

COPY --from=build /sensoringress /

COPY .env .env

EXPOSE 9100

USER nonroot:nonroot

ENTRYPOINT ["/sensoringress"]


package main

import (
	"context"
	"net/http"
	"sensor-ingress/configs"
	"time"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type SensorEntryRequest struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type SensorEntry struct {
	Name  string
	Value float64
	At    time.Time
}

func main() {
	e := echo.New()
	e.Pre(middleware.AddTrailingSlash())
	e.Pre(middleware.Logger())

	e.POST("/sensor/", saveReading, middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == "api-secret", nil
	}))
	e.Logger.Fatal(e.Start(":9100"))
}

func saveReading(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var se SensorEntry
	defer cancel()

	sensorCollection := configs.GetCollection(configs.DB, "sensorTS")

	if err := c.Bind(&se); err != nil {
		return c.String(http.StatusBadRequest, "Binding request to model failed")
	}
	newEntry := SensorEntry{
		Name:  se.Name,
		Value: se.Value,
		At:    time.Now().UTC(),
	}
	_, err := sensorCollection.InsertOne(ctx, newEntry)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to insert into collection")
	}
	return c.JSON(http.StatusCreated, newEntry)
}

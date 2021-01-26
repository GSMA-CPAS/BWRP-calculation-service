package main

import (
	"net/http"

	"github.com/GSMA-CPAS/BWRP-calculation-service/cmd/calculation-api/models"

	engine "github.com/GSMA-CPAS/BWRP-calculation-service/pkg/engine/charging"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Calculation API
// @version 1.0
// @description Calculation API

// @contact.name BWRP
// @contact.email developers@horizon.red

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http https
func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	//Handlers
	e.GET("/calculate", Calculate)

	//Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

// Calculate calculate the deal value of a contract.  godoc
// @Summary Calculate the dealvalue
// @Description Calculate the deal value by getting the contract and usage data
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} models.Result
// @Router / [post]
func Calculate(c echo.Context) error {
	var e engine.CalculationEngine
	var usage engine.AggregatedUsage
	var contract engine.Contract

	result := e.Calculate(usage, contract)
	return c.JSON(http.StatusOK, models.ConvertFromEngineResult(result))
}

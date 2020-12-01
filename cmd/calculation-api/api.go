package main

import (
	"net/http"

	engine "github.com/GSMA-CPAS/BWRP-calculation-service/pkg/engine/charging"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	//Handlers
	e.GET("/calculate", calculate)

	e.Logger.Fatal(e.Start(":8080"))
}

func calculate(c echo.Context) error {
	var e engine.CalculationEngine
	var usage engine.AggregatedUsage
	var contract engine.Contract

	result := e.Calculate(usage, contract)
	return c.JSON(http.StatusOK, result)
}

type Engine struct{}

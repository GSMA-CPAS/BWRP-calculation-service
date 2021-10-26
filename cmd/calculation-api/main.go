package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/GSMA-CPAS/BWRP-calculation-service/cmd/calculation-api/handlers"
	"github.com/GSMA-CPAS/BWRP-calculation-service/cmd/calculation-api/models"

	_ "github.com/GSMA-CPAS/BWRP-calculation-service/cmd/calculation-api/docs"
	engine "github.com/GSMA-CPAS/BWRP-calculation-service/pkg/engine/charging"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var executableHash string

const version = "0.0.1"

// @title Calculation API
// @version 1.0
// @description Calculation API

// @contact.name BWRP
// @contact.email tilak.vardhan@mobileum.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http https
func main() {

	var err error
	executableHash, err = hashFileMd5()
	if err != nil {
		os.Exit(1)
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	//Handlers
	e.POST("/calculate", Calculate)
	e.GET("/status", Status)
	e.POST("/calculatecsv", CalculateWithCSV)

	//Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

// Status provides the version and hash of the executable code of the calculation engine.  godoc
// @Summary Provide the status of the calculation engine
// @Description Provides the version and hash of the executable code of the calculation engine
// @Tags root
// @Accept json
// @Produce json
// @Success 200 {object} models.Header
// @Router /status [get]
func Status(c echo.Context) error {
	return c.JSON(http.StatusOK, models.Header{Version: version, Hash: executableHash})
}

// CalculateWithCSV calculates the dealValue from uploaded files
// @Summary Calculate the dealvalue
// @Description Calculate the deal value by processing uploaded contract and usage file
// @Tags root
// @Accept File
// @Produce json
// @Success 200 {object} models.Result
// @Router /calculatecsv [post]

func CalculateWithCSV(c echo.Context) error {
	if c.Request().Body != nil {
		contractdata, _ := c.FormFile("contractdata")
		cont, err := contractdata.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		defer cont.Close()
		contBuffer, err := io.ReadAll(cont)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		usagedata, _ := c.FormFile("usagedata")
		usage, err := usagedata.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		defer usage.Close()
		usageBuffer, err := io.ReadAll(usage)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		parsedFile := handlers.ParseFiles(contBuffer, usageBuffer)
		fmt.Print(parsedFile)
	}
	result := "ok"
	return c.JSON(http.StatusOK, result)
}

// Calculate calculate the deal value of a contract.  godoc
// @Summary Calculate the dealvalue
// @Description Calculate the deal value by getting the contract and usage data
// @Tags root
// @Accept json
// @Produce json
// @Param request body models.CalculateRequest true "Discount agreements and usage data"
// @Success 200 {object} models.Result
// @Router /calculate [post]
func Calculate(c echo.Context) error {
	var e engine.CalculationEngine
	var usage engine.AggregatedUsage
	var contract engine.Contract

	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}

	var request models.CalculateRequest

	if err := json.Unmarshal(bodyBytes, &request); err != nil {
		fmt.Printf("Error json unmarshal request %s to QueryRequest: %s\n", string(bodyBytes), err.Error())
		return c.JSON(http.StatusBadRequest, err)
	}

	usage = models.ConvertToEngineAggregatedUsage(request.Usage)
	contract = models.ConvertToEngineContract(request.DiscountModels)

	engineResult := e.Calculate(usage, contract)
	result := models.ConvertFromEngineResult(engineResult)
	result.Header = models.Header{Version: version, Hash: executableHash}
	return c.JSON(http.StatusOK, result)
}

// calculate the md5 hash of the executable
func hashFileMd5() (string, error) {
	var returnMD5String string

	self, err := os.Executable()
	if err != nil {
		return returnMD5String, err
	}

	file, err := os.Open(self)
	if err != nil {
		return returnMD5String, err
	}

	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}

	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)

	return returnMD5String, nil

}

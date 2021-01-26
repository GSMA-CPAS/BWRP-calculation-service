# BWRP-calculation-service

The calculation service applies rates on usage data.
Rates are trusted, as they are agreed on and signed by both partner organizations.
Usage data is retrieved from an external service.
As the calculation service is implemented using Golang, it is executable on a wide range of execution environments.

## Generate swagger documentation

The api in this project uses swaggo to generate the documentation.

1. Install swaggo `go get -u github.com/swaggo/swag/cmd/swag`
2. `swag init` inside the api directory to generate the documentation.
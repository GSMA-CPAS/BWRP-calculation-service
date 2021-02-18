# BWRP-calculation-service

The calculation service applies rates on usage data.
Rates are trusted, as they are agreed on and signed by both partner organizations.
Usage data is retrieved from an external service.
As the calculation service is implemented using Golang, it is executable on a wide range of execution environments.

## Generate swagger documentation

The api in this project uses swaggo to generate the documentation.

1. Install swaggo `go get -u github.com/swaggo/swag/cmd/swag`
2. `swag init` inside the api directory to generate the documentation.


Start the server `go run main.go` then visit `localhost:8080/swagger/index.html`


## docker-conpose instruction
Default `docker-compose.yaml` is configure to be run on port 8080. If this port is conflicting with other localservice, please edit
```
    ports:
      - XXXX:8080
```

and replace 'XXXX' with the desired port.

once completed to "start" run
`docker-componse up`

to "stop and remove" run
`docker-componse down --rmi all`

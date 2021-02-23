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


## docker-compose instruction
Default `docker-compose.yaml` is configure to be run on port 8080. If this port is conflicting with other localservice, please edit
```
    ports:
      - XXXX:8080
```

and replace 'XXXX' with the desired port.

once completed to "start" run
`docker-compose up`

to "stop and remove" run
`docker-compose down --rmi all`

example request (please replace the localhost and XXXX port number accordingly)
`curl -X POST "http://localhost:XXXX/calculate" -H  "accept: application/json" -d '{{"discountModels":{"DTAG":{"condition":{"selectedConditionName":"Unconditional"},"serviceGroups":[{"homeTadigs":{"codes":["AAZOR","GRCPF"]},"visitorTadigs":{"codes":["AAZVF","AAZTD"]},"chosenServices":[{"name":"SMSMO","includedInCommitment":true,"rate":[{"threshold":0,"linearPrice":5},{"threshold":1500,"linearPrice":3}]},{"name":"MOC","rate":[{"threshold":0,"fixedPrice":1500}]},{"name":"VOLTE","balancedRate":[{"threshold":0,"linearPrice":15}],"unbalancedRate":[{"threshold":0,"linearPrice":20,"fixedPrice":1000}]},{"name":"MOCEU","rate":[{"threshold":0,"linearPrice":12}]}]}]},"TMUS":{"condition":{"selectedConditionName":"Unconditional"},"serviceGroups":[{"homeTadigs":{"codes":["AAZVF","AAZTD"]},"visitorTadigs":{"codes":["AAZOR","GRCPF"]},"services":[{"name":"SMSMO","rate":[{"threshold":0,"fixedPrice":5000}]},{"name":"MOC","rate":[{"threshold":0,"fixedPrice":3000}]},{"name":"VOLTE","balancedRate":[{"threshold":0,"linearPrice":15}],"unbalancedRate":[{"threshold":0,"linearPrice":20,"fixedPrice":1000}]},{"name":"MOCEU","rate":[{"threshold":0,"linearPrice":15}]}]}]}},"usage":[{"service":"SMSMO","volume":3000.5,"charge":1007,"homeTadig":"AAZVF","visitorTadig":"AAZOR"},{"service":"SMSMO","volume":3000,"charge":1007,"homeTadig":"GRCPF","visitorTadig":"AAZTD"},{"service":"MOC","volume":2000,"charge":1005,"homeTadig":"AAZTD","visitorTadig":"AAZOR"},{"service":"MOC","volume":5000,"charge":3003,"homeTadig":"AAZOR","visitorTadig":"AAZVF"},{"service":"VOLTE","volume":3000,"charge":1000,"homeTadig":"GRCPF","visitorTadig":"HOR2"},{"service":"VOLTE","volume":2000,"charge":800,"homeTadig":"AAZVF","visitorTadig":"GRCPF"},{"service":"MOCEU","volume":6000,"charge":800,"homeTadig":"GRCPF","visitorTadig":"AAZVF"},{"service":"MOCEU","volume":4000,"charge":800,"homeTadig":"AAZTD","visitorTadig":"AAZOR"}]}}'`

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

example request (please replace the localhost and XXXX port number accordingly)
`curl -X POST "http://localhost:XXXX/calculate" -H  "accept: application/json" -d '{"discounts":{"HOME":{"serviceGroups":[{"homeTadigs":["HOR1"],"visitorTadigs":["HOR2"],"services":[{"service":"SMSMO","usagePricing":{"ratingPlan":{"rate":{"thresholds":[{"start":0,"linearPrice":5},{"start":1500,"linearPrice":3}]}}}},{"service":"MOC","usagePricing":{"ratingPlan":{"rate":{"thresholds":[{"start":0,"fixedPrice":1500}]}}}},{"service":"VOLTE","usagePricing":{"ratingPlan":{"balancedRate":{"value":10},"unbalancedRate":{"value":20}}}},{"service":"MOCEU","usagePricing":{"ratingPlan":{"rate":{"linearPrice":12}}}}]}]},"VISITOR":{"serviceGroups":[{"homeTadigs":["HOR2"],"visitorTadigs":["HOR1"],"services":[{"service":"SMSMO","usagePricing":{"ratingPlan":{"rate":{"thresholds":[{"start":0,"fixedPrice":5000}]}}}},{"service":"MOC","usagePricing":{"ratingPlan":{"rate":{"thresholds":[{"start":0,"fixedPrice":3000}]}}}},{"service":"VOLTE","usagePricing":{"ratingPlan":{"balancedRate":{"value":10},"unbalancedRate":{"value":20}}}},{"service":"MOCEU","usagePricing":{"ratingPlan":{"rate":{"linearPrice":15}}}}]}]}},"usage":[{"service":"SMSMO","volume":3000,"charge":1007,"homeTadig":"HOR1","visitorTadig":"HOR2"},{"service":"MOC","volume":2000,"charge":1005,"homeTadig":"HOR2","visitorTadig":"HOR1"},{"service":"MOC","volume":5000,"charge":3003,"homeTadig":"HOR1","visitorTadig":"HOR2"},{"service":"VOLTE","volume":3000,"charge":1000,"homeTadig":"HOR1","visitorTadig":"HOR2"},{"service":"VOLTE","volume":2000,"charge":800,"homeTadig":"HOR2","visitorTadig":"HOR1"},{"service":"MOCEU","volume":6000,"charge":800,"homeTadig":"HOR1","visitorTadig":"HOR2"},{"service":"MOCEU","volume":4000,"charge":800,"homeTadig":"HOR2","visitorTadig":"HOR1"}]}'`

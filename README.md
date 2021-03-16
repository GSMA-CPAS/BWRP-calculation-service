# BWRP-calculation-service

## General
The calculation service applies rates as specified in the discount agreement on usage data.  
Rates are trusted, as they are agreed on and signed by both partner organizations.  
Usage data is retrieved from an external service.  

The calculation services consists of an API and a calculation package.  
As the calculation service is implemented using Golang, it is executable on a wide range of execution environments.  

## Functional
The calculate API expects the discount agreement and usage data. The discount agreement is then converted to the model used by the calculation package.    
The usage data in the API request is converted to an aggregated map of usage data as according to the model used by the calculation package.  

The calculation package executes the different charging models as specified in the discount agreement with the aggregated usage data.  
The result consists of a dealValue per service and homeTadig / visitorTadig combinations.  

The calculation packgae result is converted to an API result.

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
`curl -X POST "http://localhost:XXXX/calculate" -H  "accept: application/json" -d '{"discounts":{"DTAG":{"condition":{"kind":"Unconditional"},"serviceGroups":[{"homeTadigs":["AAZOR","GRCPF"],"visitorTadigs":["AAZVF","AAZTD"],"services":[{"service":"SMSMO","includedInCommitment":true,"usagePricing":{"ratingPlan":{"rate":{"thresholds":[{"start":"0","linearPrice":"5"},{"start":"1500","linearPrice":"3"}]}}}},{"service":"MOCInternational","usagePricing":{"ratingPlan":{"rate":{"thresholds":[{"start":"0","fixedPrice":"1500"}]}}}},{"service":"VOLTE","usagePricing":{"ratingPlan":{"balancedRate":{"thresholds":[{"start":"0","linearPrice":"15"}]},"unbalancedRate":{"thresholds":[{"start":"0","linearPrice":"20","fixedPrice":"1000"}]}}}},{"service":"MOCEU","usagePricing":{"ratingPlan":{"rate":{"thresholds":[{"start":"0","linearPrice":"12"}]}}}},{"service":"NB-IOT","accessPricing":{"ratingPlan":{"rate":{"thresholds":[{"start":"0","linearPrice":"12"},{"start":"1000","linearPrice":"5"}]}}}}]}]},"TMUS":{"condition":{"kind":"Unconditional"},"serviceGroups":[{"homeTadigs":["AAZVF","AAZTD"],"visitorTadigs":["AAZOR","GRCPF"],"services":[{"service":"SMSMO","usagePricing":{"ratingPlan":{"rate":{"thresholds":[{"start":"0","fixedPrice":"5000"}]}}}},{"service":"MOCInternational","usagePricing":{"ratingPlan":{"rate":{"thresholds":[{"start":"0","fixedPrice":"3000"}]}}}},{"service":"VOLTE","usagePricing":{"ratingPlan":{"balancedRate":{"thresholds":[{"start":"0","linearPrice":"15"}]},"unbalancedRate":{"thresholds":[{"start":"0","linearPrice":"20.5","fixedPrice":"1000"}]}}}},{"service":"MOCEU","usagePricing":{"ratingPlan":{"rate":{"thresholds":[{"start":"0","linearPrice":"15"}]}}}},{"service":"NB-IOT","accessPricing":{"ratingPlan":{"rate":{"thresholds":[{"start":"0","linearPrice":"12"},{"start":"500","linearPrice":"8"}]}}}}]}]}},"usage":{"inbound":[{"service":"SMSMO","usage":"3000","charges":"1007","taxes":"15","homeTadig":"GRCPF","visitorTadig":"AAZTD"},{"service":"MOCInternational","usage":"5000","charges":"3003","taxes":"15","homeTadig":"AAZOR","visitorTadig":"AAZVF"},{"service":"VOLTE","usage":"3000","charges":"1000","taxes":"10","homeTadig":"GRCPF","visitorTadig":"AAZVF"},{"service":"MOCEU","usage":"6000","charges":"900","taxes":"9","homeTadig":"GRCPF","visitorTadig":"AAZVF"},{"service":"NB-IOT","usage":"2000","charges":"600","taxes":"6","homeTadig":"GRCPF","visitorTadig":"AAZVF"}],"outbound":[{"service":"SMSMO","usage":"3000.5","charges":"1007","taxes":"15","homeTadig":"AAZVF","visitorTadig":"AAZOR"},{"service":"MOCInternational","usage":"2000","charges":"1005","taxes":"15","homeTadig":"AAZTD","visitorTadig":"AAZOR"},{"service":"VOLTE","usage":"2000","charges":"800.5","taxes":"8","homeTadig":"AAZVF","visitorTadig":"GRCPF"},{"service":"MOCEU","usage":"4000","charges":"1200","taxes":"12","homeTadig":"AAZTD","visitorTadig":"AAZOR"},{"service":"NB-IOT","usage":"3000","charges":"800","taxes":"8","homeTadig":"AAZTD","visitorTadig":"AAZOR"}]}}'`

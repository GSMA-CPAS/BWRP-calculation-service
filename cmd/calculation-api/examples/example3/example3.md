# Description

Given discount agreements, one conditional on Deal Revenue Commitment and the other unconditional and usage data, the result will be the calculated dealValue for each (sub-)service.

## Result

For the conditional discount agreement the aggregated calculated <b>dealValue</b> for the all services that have the includedInCommitment flag set, will be compared to the commitment value.  
If commitment value > aggregated calculated dealValue for included services => no deal, no result per service  
If commitment value < aggregated calculated dealValue for included services => deal, result will include the calculated dealValue per service  
  
In example the included service is has following charging model:
```
{
    "service":"SMSMO",
    "includedInCommitment":true,
    "usagePricing":{
        "ratingPlan":{
        "rate":{
            "thresholds":[
                {
                    "start":"0",
                    "linearPrice":"5"
                },
                {
                    "start":"1500",
                    "linearPrice":"3"
                }
            ]
        }
        }
    }
},
```

And following usage data:
```
       {
          "service":"SMSMO",
          "usage":"3000",
          "charges":"1007",
          "taxes":"15",
          "homeTadig":"GRCPF",
          "visitorTadig":"AAZTD"
       },
```

Calculated dealValue = 1500 * 5 + 1500 * 3 = 12000  

Condition in discount agreement:
```
"condition":{
    "kind":"Deal Revenue Commitment",
    "commitment":{
    "value":"12500"
    }
},
```
This means no deal, so for the homeTadigs in this discount agreement there will be no dealValue per service  
If the commitment value is changed to 11000 there will be a deal and the result will include the dealValue per service for this discount agreement 
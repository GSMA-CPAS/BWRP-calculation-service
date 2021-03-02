# Description

Given discount agreements, one conditional on Contract Revenue Commitment and the other unconditional and usage data, the result will be the calculated dealValue for each (sub-)service.

## Result

For the conditional discount agreement the aggregated <b>charges</b> will be compared to the commitment value.  
If includingTaxes flag is set the taxes in the usage data will be included.  
If commitment value > aggregated charges from usage data => no deal, no result per service  
If commitment value < aggregated charges from usage data => deal, result will include the aggregated charge per service  

In example the conditions is:
```
"condition":{
    "kind":"Contract Revenue Commitment",
    "commitment":{
    "value":"6550",
    "includingTaxes":true
    }
},
```   

Where the aggregated charges including taxes for the usage of homeTadigs GRCPF or AAZOR are (charge + tax):  

(1007 + 15) + (3003 + 15) + (1000 + 10) + (900 + 9) + (600 + 6) = 6565

The result will be a deal, and included the aggregated charges per service

If includingTaxes would be false the charge would be 6510 and that means no deal.

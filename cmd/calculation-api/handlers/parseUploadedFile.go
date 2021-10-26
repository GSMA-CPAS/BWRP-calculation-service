package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/GSMA-CPAS/BWRP-calculation-service/cmd/calculation-api/models"
)

// Parse the uploaded files and generate json for calculate API
func ParseFiles(contBuffer []byte, usageBuffer []byte) string {

	var cont models.ParseFileRequest
	err := json.Unmarshal(contBuffer, &cont)
	if err != nil {
		panic(err)
	}
	contbytes, err := json.Marshal(cont.Body.Discounts)
	if err != nil {
		panic(err)
	}
	fmt.Println("{\"discounts\":" + string(contbytes) + "}")
	return ""
}

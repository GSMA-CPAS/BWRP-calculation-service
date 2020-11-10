/*
 License-Placeholder
*/

package engine

import models "github.com/GSMA-CPAS/BWRP-calculation-service/pkg/engine/models"

type FlatRate struct {
	rate int64
}

func (c *FlatRate) Charge(usage models.Usage) int64 {
	return c.rate
}

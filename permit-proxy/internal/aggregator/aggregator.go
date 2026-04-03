package aggregator

import (
	"permit-proxy/internal/models"
)

type ZipCount struct {
	Zipcode string `json:"zipcode"`
	Count int `json:"count"`
}

func ByZip(p []models.Permit) ([]ZipCount, error) {
	z := make(map[string]int)
	c := make([]ZipCount, 0)
	
	for _, i := range p{
		z[i.Zipcode]++
	}

	for k, v := range z {
		c = append(c, ZipCount{
			Zipcode: k,
			Count: v,
		})
	}

	return c, nil
}

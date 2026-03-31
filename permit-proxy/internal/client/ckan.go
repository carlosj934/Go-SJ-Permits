package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"permit-proxy/internal/models"
)

const (
	pageSize    = 1000
	CKANBaseUrl = "https://data.sanjoseca.gov/api/3/action/datastore_search?resource_id=761b7ae8-3be1-4ad6-923d-c7af6404a904&limit=%d&offset=%d"
)

func Paginator(url string) ([]models.RawPermitData, error) {
	offset := 0
	accumulator := make([]models.RawPermitData, 0)

	for {
		var result models.WrapResult
		urlFormatted := fmt.Sprintf(url, pageSize, offset)

		res, err := http.Get(urlFormatted)
		if err != nil {
			return nil, err
		}

		if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
			return nil, err
		}

		res.Body.Close()

		accumulator = append(accumulator, result.Result.Records...)
		if int64(len(accumulator)) >= result.Result.Total {
			break
		}

		offset += pageSize

	}

	return accumulator, nil
}

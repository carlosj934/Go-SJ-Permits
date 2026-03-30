package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"permit-proxy/internal/models"
)

func TestPaginator(t *testing.T) {
	tests := []struct {
		name    string
		handler func(w http.ResponseWriter, r *http.Request)
		wantLen int
		wantErr bool
	}{
		{
			name: "server returns valid JSON with records",
			handler: func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, `
					{
						"success": true,
						"result": {
							"total": 2,
							"records": [{"Status":"Active","gx_location":"     ,   ","FOLDERNUMBER":"2015-101797-IR"},{"Status":"Active","gx_location":"     ,   ","FOLDERNUMBER":"2015-001272-IR"}]
						}
					}
					`)
			},
			wantLen: 2,
			wantErr: false,
		},
		{
			name: "return error if invalid JSON is provided",
			handler: func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, `"invalid json"`)
			},
			wantLen: 0,
			wantErr: true,
		},
		{
			name: "verify Paginator correctly fetches multiple pages",
			handler: func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Query().Get("offset") == "0" {
					test := make([]models.RawPermitData, 0)
					for i := 0; i < 1000; i++ {
						test = append(test, models.RawPermitData{})
					}

					json.NewEncoder(w).Encode(&models.WrapResult{
						Success: true,
						Result: models.WrapRecords{
							Total:   1001,
							Records: test,
						},
					})
				} else if r.URL.Query().Get("offset") == "1000" {
					test := make([]models.RawPermitData, 1)

					json.NewEncoder(w).Encode(&models.WrapResult{
						Success: true,
						Result: models.WrapRecords{
							Total:   1001,
							Records: test,
						},
					})
				}
			},
			wantLen: 1001,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exec := httptest.NewServer(http.HandlerFunc(tt.handler))
			defer exec.Close()

			records, err := Paginator(exec.URL + "?limit=%d&offset=%d")

			if tt.wantErr == true {
				require.Error(t, err)
			} else if tt.wantErr == false {
				require.NoError(t, err)
			}

			assert.Equal(t, len(records), tt.wantLen)

		})
	}
}

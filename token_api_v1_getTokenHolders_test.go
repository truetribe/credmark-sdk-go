package credmark

import (
	"log"
	"os"
	"reflect"
	"testing"

	. "github.com/go-playground/assert/v2"
	"github.com/joho/godotenv"
)

func TestIntegrationGetTokenHoldersSuccess(t *testing.T) {

	cases := []struct {
		name         string
		chainId      int
		tokenAddress string
		pageSize     int
		cursor       string
		quoteAddress string
		scaled       bool
		blockNum     int
		timestamp    int
	}{
		{
			name:         "TestGetTokenHoldersSuccess",
			chainId:      CHAIN_ID_ETHEREUM,
			tokenAddress: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
			pageSize:     10,
			cursor:       "null",
			quoteAddress: "",
			scaled:       true,
			blockNum:     0,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			err := godotenv.Load()
			if err != nil {
				log.Fatal("Error loading .env file")
			}

			cfg := &Config{
				ApiKey: os.Getenv("CREDMARK_API_KEY"),
			}
			c, _ := NewClient(cfg)

			var payload GetTokenHoldersPayload
			payload.ChainId = tc.chainId
			payload.TokenAddr = tc.tokenAddress
			payload.PageSize = tc.pageSize
			if tc.blockNum > 0 {
				payload.BlockNum = tc.blockNum
			}

			if tc.timestamp > 0 {
				payload.Timestamp = tc.timestamp
			}

			respType := reflect.TypeOf(GetTokenHoldersCredmarkResponse{})
			dataField, _ := respType.FieldByName("Data")

			resp, err := c.GetTokenHolders(payload)
			Equal(t, err, nil)
			Equal(t, resp.TokenAddress, tc.tokenAddress)
			Equal(t, resp.BlockTimestamp, tc.timestamp)
			Equal(t, reflect.ValueOf(resp.Data).Type(), dataField.Type)
		})
	}
}

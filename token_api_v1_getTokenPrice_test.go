package credmark

import (
	"log"
	"os"
	"testing"

	. "github.com/go-playground/assert/v2"
	"github.com/joho/godotenv"
)

func TestIntegrationGetTokenPriceSuccess(t *testing.T) {

	cases := []struct {
		name         string
		chainId      int
		tokenAddress string
		blockNum     int
		timestamp    int
		expect       string
	}{
		{
			name:         "TestIntegrationGetTokenPriceSuccess - ETH",
			chainId:      CHAIN_ID_ETHEREUM,
			tokenAddress: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
			blockNum:     15490034,
			expect:       "",
		},
		{
			name:         "TestIntegrationGetTokenPriceSuccess - ETH",
			chainId:      CHAIN_ID_ETHEREUM, //ETH
			tokenAddress: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
			timestamp:    1682582087,
			expect:       "",
		},
		{
			name:         "TestIntegrationGetTokenPriceSuccess - BNB Chain",
			chainId:      CHAIN_ID_BNB_CHAIN, //BNB Chain
			tokenAddress: "0x2170Ed0880ac9A755fd29B2688956BD959F933F8",
			timestamp:    1682482087,
			expect:       "",
		},
		{
			name:         "TestIntegrationGetTokenPriceSuccess - POLYGON",
			chainId:      CHAIN_ID_POLYGON, //Polygon
			tokenAddress: "0x7ceB23fD6bC0adD59E62ac25578270cFf1b9f619",
			timestamp:    1680582087,
			expect:       "",
		},
		{
			name:         "TestIntegrationGetTokenPriceSuccess - ARB",
			chainId:      CHAIN_ID_ARBITRUM, //Arb
			tokenAddress: "0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9",
			timestamp:    1682512087,
			expect:       "",
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

			var payload GetTokenPricePayload
			payload.ChainId = tc.chainId
			payload.TokenAddr = tc.tokenAddress
			if tc.blockNum > 0 {
				payload.BlockNum = tc.blockNum
			}

			if tc.timestamp > 0 {
				payload.Timestamp = tc.timestamp
			}

			payload.Source = ""

			resp, err := c.GetTokenPrice(payload)
			Equal(t, err, nil)
			Equal(t, resp.TokenAddress, tc.tokenAddress)
			Equal(t, resp.BlockTimestamp, tc.timestamp)
		})
	}
}

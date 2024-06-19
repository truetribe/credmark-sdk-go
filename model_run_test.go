package credmark

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	. "github.com/go-playground/assert/v2"

	"github.com/joho/godotenv"
)

type TokenCategorizedSupply struct {
	Slug           string `json:"slug"`
	Version        string `json:"version"`
	ChainID        int    `json:"chainId"`
	BlockNumber    int    `json:"blockNumber"`
	BlockTimestamp int    `json:"blockTimestamp"`
	Output         struct {
		Categories []struct {
			Accounts struct {
				Accounts []struct {
					Address string `json:"address"`
				} `json:"accounts"`
			} `json:"accounts"`
			CategoryName string  `json:"categoryName"`
			CategoryType string  `json:"categoryType"`
			Circulating  bool    `json:"circulating"`
			AmountScaled float64 `json:"amountScaled"`
			ValueUsd     float64 `json:"valueUsd"`
		} `json:"categories"`
		Token struct {
			Address string `json:"address"`
		} `json:"token"`
		TotalSupplyScaled       float64 `json:"totalSupplyScaled"`
		TotalSupplyUsd          float64 `json:"totalSupplyUsd"`
		CirculatingSupplyScaled float64 `json:"circulatingSupplyScaled"`
		CirculatingSupplyUsd    float64 `json:"circulatingSupplyUsd"`
	} `json:"output"`
	Dependencies struct {
		ContractMetadata struct {
			One1 int `json:"1.1"`
		} `json:"contract.metadata"`
		TokenUnderlyingMaybe struct {
			One1 int `json:"1.1"`
		} `json:"token.underlying-maybe"`
		ChainlinkGetFeedRegistry struct {
			One0 int `json:"1.0"`
		} `json:"chainlink.get-feed-registry"`
		ChainlinkPriceByRegistry struct {
			One7 int `json:"1.7"`
		} `json:"chainlink.price-by-registry"`
		ChainlinkPriceFromRegistryMaybe struct {
			One4 int `json:"1.4"`
		} `json:"chainlink.price-from-registry-maybe"`
		PriceOracleChainlink struct {
			One13 int `json:"1.13"`
		} `json:"price.oracle-chainlink"`
		PriceOracleChainlinkMaybe struct {
			One3 int `json:"1.3"`
		} `json:"price.oracle-chainlink-maybe"`
		PriceCex struct {
			Zero6 int `json:"0.6"`
		} `json:"price.cex"`
		PriceCexMaybe struct {
			Zero6 int `json:"0.6"`
		} `json:"price.cex-maybe"`
		PriceQuote struct {
			One16 int `json:"1.16"`
		} `json:"price.quote"`
		TokenCategorizedSupply struct {
			One3 int `json:"1.3"`
		} `json:"token.categorized-supply"`
	} `json:"dependencies"`
	Cached  bool `json:"cached"`
	Runtime int  `json:"runtime"`
}

func TestIntegrationModelRunSuccess(t *testing.T) {

	cases := []struct {
		name        string
		requestBody string
		expect      string
	}{
		{
			name:   "TestIntegrationRunModelSuccess",
			expect: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
			requestBody: `{
    "slug": "token.categorized-supply",
    "chainId": 1,
    "blockNumber": "latest",
    "input": {
        "token": {
            "address": "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"
        },
        "categories": [
            {
                "accounts": {
                    "accounts": [
                        {
                            "address": "0x1F98431c8aD98523631AE4a59f267346ea31F984"
                        }
                    ]
                },
                "categoryType": "treasury",
                "categoryName": "Treasury",
                "circulating": false
            },
            {
                "accounts": {
                    "accounts": [
                        {
                            "address": "0x1F98431c8aD98523631AE4a59f267346ea31F984"
                        }
                    ]
                },
                "categoryType": "locked",
                "categoryName": "Locked",
                "circulating": false
            }
        ]
    }
}`,
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

			resp, _ := c.RunModel(tc.requestBody)

			var data TokenCategorizedSupply
			err = json.Unmarshal([]byte(resp), &data)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			Equal(t, data.Output.Token.Address, tc.expect)
		})
	}
}

package credmark

import (
	"log"
	"os"
	"testing"

	. "github.com/go-playground/assert/v2"
	"github.com/joho/godotenv"
)

func TestIntegrationGetTokenVolumeSuccess(t *testing.T) {

	cases := []struct {
		name    string
		address string
		window  string
		expect  string
	}{
		{
			name:    "TestIntegrationGetTokenVolumeSuccess",
			address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
			window:  "24 hours",
			expect:  "",
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

			var payload GetTokenVolumePayload
			payload.Address = tc.address
			payload.Window = tc.window

			resp, err := c.GetTokenVolume(payload)
			Equal(t, err, nil)
			Equal(t, resp.Output.Address, tc.address)
		})
	}
}

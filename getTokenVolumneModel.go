package credmark

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime/debug"

	log "github.com/sirupsen/logrus"
)

type GetTokenVolumePayload struct {
	Address      string `validate:"required"`
	Window       string `validate:"required"` //"24 hours"
	IncludePrice bool
}

type GetTokenVolumeResponse struct {
	Raw    string
	Output CredmarkTokenVolumeResponseOutput
}

type CredmarkGetTokenVolumePayload struct {
	Slug        string                             `json:"slug"`
	ChainID     int                                `json:"chainId"`
	BlockNumber string                             `json:"blockNumber"`
	Input       CredmarkGetTokenVolumePayloadInput `json:"input"`
}
type CredmarkGetTokenVolumePayloadInput struct {
	Address      string `json:"address"`
	Window       string `json:"window"`
	IncludePrice bool   `json:"include_price"`
}

type CredmarkTokenVolumeResponse struct {
	Slug         string                                  `json:"slug"`
	Version      string                                  `json:"version"`
	ChainID      int                                     `json:"chainId"`
	BlockNumber  int                                     `json:"blockNumber"`
	Output       CredmarkTokenVolumeResponseOutput       `json:"output"`
	Dependencies CredmarkTokenVolumeResponseDependencies `json:"dependencies"`
	Runtime      int                                     `json:"runtime"`
}
type CredmarkTokenVolumeResponseOutput struct {
	Address      string  `json:"address"`
	Volume       int64   `json:"volume"`
	VolumeScaled float64 `json:"volume_scaled"`
	ValueLast    float64 `json:"value_last"`
	FromBlock    int     `json:"from_block"`
	ToBlock      int     `json:"to_block"`
}
type RPCGetBlocknumber struct {
	One0 int `json:"1.0"`
}
type LedgerErc20TokenTransferData struct {
	One0 int `json:"1.0"`
}
type ContractMetadata struct {
	One0 int `json:"1.0"`
}
type TokenUnderlyingMaybe struct {
	One1 int `json:"1.1"`
}
type ChainlinkGetFeedRegistry struct {
	One0 int `json:"1.0"`
}
type ChainlinkPriceByRegistry struct {
	One3 int `json:"1.3"`
}
type ChainlinkPriceFromRegistryMaybe struct {
	One2 int `json:"1.2"`
}
type PriceOracleChainlink struct {
	One7 int `json:"1.7"`
}
type PriceOracleChainlinkMaybe struct {
	One2 int `json:"1.2"`
}
type PriceQuote struct {
	One9 int `json:"1.9"`
}
type TokenOverallVolumeBlock struct {
	One0 int `json:"1.0"`
}
type TokenOverallVolumeWindow struct {
	One0 int `json:"1.0"`
}
type CredmarkTokenVolumeResponseDependencies struct {
	RPCGetBlocknumber               RPCGetBlocknumber               `json:"rpc.get-blocknumber"`
	LedgerErc20TokenTransferData    LedgerErc20TokenTransferData    `json:"ledger.erc20_token_transfer_data"`
	ContractMetadata                ContractMetadata                `json:"contract.metadata"`
	TokenUnderlyingMaybe            TokenUnderlyingMaybe            `json:"token.underlying-maybe"`
	ChainlinkGetFeedRegistry        ChainlinkGetFeedRegistry        `json:"chainlink.get-feed-registry"`
	ChainlinkPriceByRegistry        ChainlinkPriceByRegistry        `json:"chainlink.price-by-registry"`
	ChainlinkPriceFromRegistryMaybe ChainlinkPriceFromRegistryMaybe `json:"chainlink.price-from-registry-maybe"`
	PriceOracleChainlink            PriceOracleChainlink            `json:"price.oracle-chainlink"`
	PriceOracleChainlinkMaybe       PriceOracleChainlinkMaybe       `json:"price.oracle-chainlink-maybe"`
	PriceQuote                      PriceQuote                      `json:"price.quote"`
	TokenOverallVolumeBlock         TokenOverallVolumeBlock         `json:"token.overall-volume-block"`
	TokenOverallVolumeWindow        TokenOverallVolumeWindow        `json:"token.overall-volume-window"`
}

// https://gateway.credmark.com/api/?urls.primaryName=DeFi%20API%20-%20Run%20Model%20Requests#/Models/runModel-token.overall-volume-window
func (c *Client) GetTokenVolume(payload GetTokenVolumePayload) (gtvResp GetTokenVolumeResponse, err error) {

	if err := ValidateStruct(payload); err != nil {
		log.Error(err, string(debug.Stack()))
		return gtvResp, err
	}

	var gtvPayload CredmarkGetTokenVolumePayload
	gtvPayload.Slug = "token.overall-volume-window"
	gtvPayload.ChainID = 1
	gtvPayload.BlockNumber = "latest"
	gtvPayload.Input.Address = payload.Address
	gtvPayload.Input.Window = payload.Window
	gtvPayload.Input.IncludePrice = payload.IncludePrice

	if err := ValidateStruct(gtvPayload); err != nil {
		log.Error(err, string(debug.Stack()))
		return gtvResp, err
	}

	b, err := json.Marshal(gtvPayload)
	if err != nil {
		return gtvResp, err
	}

	apiReq := ApiRequest{
		Endpoint: CREDMARK_GATEWAY_URL + CREDMARK_API_V1_XR_URI,
		Method:   "POST",
		Body:     string(b),
	}

	res, err := c.doHttpRequeset(apiReq, GetCurrentFuncName())
	if err != nil {
		return gtvResp, err
	}
	defer res.Body.Close()

	response := &CredmarkTokenVolumeResponse{}

	if res.StatusCode != http.StatusOK { //exception
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		bodyString := string(bodyBytes)
		return gtvResp, fmt.Errorf("%v: Credmark GetTokenVolume Error http response status code: %v, %v", "token.overall-volume-window", res.StatusCode, bodyString)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return gtvResp, fmt.Errorf("%v: Credmark GetTokenVolume Error: %v", "token.overall-volume-window", body)
	}
	gtvResp.Raw = string(body[:])
	_ = json.Unmarshal(body, &response)
	gtvResp.Output = response.Output

	return
}

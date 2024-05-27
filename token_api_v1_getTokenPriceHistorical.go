package credmark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"runtime/debug"
	"strconv"

	log "github.com/sirupsen/logrus"
)

const (
	CREDMARK_API_V1_URI_TOKEN_PRICE_HISTORICAL = "/v1/tokens/%s/%s/price" //chainId, tokenAddress
)

type GetTokenPriceHistoricalPayload struct {
	ChainID          int    `validate:"required"` //Chain ID. Use 1 for mainnet.
	TokenAddr        string `validate:"required"` //The address of the token requested.
	StartBlockNumber *uint64
	EndBlockNumber   *uint64
	BlockInterval    *uint64
	StartTimestamp   *uint64
	EndTimestamp     *uint64
	TimeInterval     *uint64
	QuoteAddress     string
	Src              string
}

type GetTokenPriceHistoricalCredmarkResponse struct {
	ChainID          int    `json:"chainId"`
	StartBlockNumber int    `json:"startBlockNumber"`
	EndBlockNumber   int    `json:"endBlockNumber"`
	StartTimestamp   int    `json:"startTimestamp"`
	EndTimestamp     int    `json:"endTimestamp"`
	TokenAddress     string `json:"tokenAddress"`
	QuoteAddress     string `json:"quoteAddress"`
	Data             []struct {
		BlockNumber    int     `json:"blockNumber"`
		BlockTimestamp int     `json:"blockTimestamp"`
		Price          float64 `json:"price"`
		Src            string  `json:"src"`
		SrcInternal    string  `json:"srcInternal"`
	} `json:"data"`
}

func (c *Client) GetTokenPriceHistorical(payload GetTokenPriceHistoricalPayload) (response GetTokenPriceHistoricalCredmarkResponse, err error) {

	if err := ValidateStruct(payload); err != nil {
		log.Error(err, string(debug.Stack()))
		return response, err
	}

	uri := fmt.Sprintf(CREDMARK_API_V1_URI_TOKEN_PRICE_HISTORICAL, strconv.Itoa(payload.ChainID), payload.TokenAddr)
	if err != nil {
		log.Error(err, string(debug.Stack()))
		return response, err
	}

	endpoint := CREDMARK_GATEWAY_URL + uri
	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		log.Error(err, string(debug.Stack()))
		return response, err
	}

	values := queryUrl.Query()

	if payload.StartBlockNumber != nil {
		values.Add("startBlockNumber", strconv.FormatUint(*payload.StartBlockNumber, 10))
	}
	if payload.EndBlockNumber != nil {
		values.Add("endBlockNumber", strconv.FormatUint(*payload.EndBlockNumber, 10))
	}
	if payload.BlockInterval != nil {
		values.Add("blockInterval", strconv.FormatUint(*payload.BlockInterval, 10))
	}
	if payload.StartTimestamp != nil {
		values.Add("startTimestamp", strconv.FormatUint(*payload.StartTimestamp, 10))
	}
	if payload.EndTimestamp != nil {
		values.Add("endTimestamp", strconv.FormatUint(*payload.EndTimestamp, 10))
	}
	if payload.TimeInterval != nil {
		values.Add("timeInterval", strconv.FormatUint(*payload.TimeInterval, 10))
	}
	if payload.QuoteAddress != "" {
		values.Add("quoteAddress", payload.QuoteAddress)
	}
	if payload.Src != "" {
		values.Add("src", payload.Src)
	}

	queryUrl.RawQuery = values.Encode()

	apiReq := ApiRequest{
		Endpoint: queryUrl.String(),
		Method:   "GET",
	}

	res, err := c.doHttpRequeset(apiReq, GetCurrentFuncName())
	if err != nil {
		return response, err
	}
	defer res.Body.Close()

	body := &bytes.Buffer{}
	_, err = io.Copy(body, res.Body)
	if err != nil {
		return response, fmt.Errorf("%v: Response Error: %v", CREDMARK_API_V1_URI_TOKEN_PRICE_HISTORICAL, body)
	}

	err = json.NewDecoder(body).Decode(&response)
	if err != nil {
		return response, fmt.Errorf("%v: Decode Error: %v", CREDMARK_API_V1_URI_TOKEN_PRICE_HISTORICAL, err)
	}

	return
}

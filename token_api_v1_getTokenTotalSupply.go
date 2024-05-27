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
	CREDMARK_TOKEN_API_V1_TOTAL_SUPPLY = "/v1/tokens/%s/%s/total-supply" //chainId, tokenAddress
)

type GetTokenTotalSupplyPayload struct {
	ChainID     int    `validate:"required"`
	TokenAddr   string `validate:"required"`
	BlockNumber *uint64
	Timestamp   *uint64
	Scaled      bool
}

type GetTokenTotalSupplyResponse struct {
	ChainID        int    `json:"chainId"`
	BlockNumber    int    `json:"blockNumber"`
	BlockTimestamp int    `json:"blockTimestamp"`
	TokenAddress   string `json:"tokenAddress"`
	Scaled         bool   `json:"scaled"`
	TotalSupply    int    `json:"totalSupply"`
}

func (c *Client) GetTokenTotalSupply(payload GetTokenTotalSupplyPayload) (response GetTokenTotalSupplyResponse, err error) {

	slug := CREDMARK_TOKEN_API_V1_TOTAL_SUPPLY

	if err := ValidateStruct(payload); err != nil {
		log.Error(err, string(debug.Stack()))
		return response, err
	}

	uri := fmt.Sprintf(slug, strconv.Itoa(payload.ChainID), payload.TokenAddr)
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

	values.Add("scaled", strconv.FormatBool(payload.Scaled))

	if payload.BlockNumber != nil {
		values.Add("blockNumber", strconv.FormatUint(*payload.BlockNumber, 10))
	}

	if payload.Timestamp != nil {
		values.Add("timestamp", strconv.FormatUint(*payload.Timestamp, 10))
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
		return response, fmt.Errorf("%v: Response Error: %v", slug, body)
	}

	err = json.NewDecoder(body).Decode(&response)
	if err != nil {
		return response, fmt.Errorf("%v: Decode Error: %v", slug, err)
	}

	return
}

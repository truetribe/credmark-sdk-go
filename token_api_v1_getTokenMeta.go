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
	CREDMARK_API_V1_TOKEN_META = "/v1/tokens/%s/%s" //chainId, tokenAddress
)

type GetTokenMetaPayload struct {
	ChainID      int    `validate:"required"`
	TokenAddress string `validate:"required"`
	BlockNumber  *uint64
	Timestamp    *uint64
}

type GetTokenMetaResponse struct {
	ChainID        int    `json:"chainId"`
	BlockNumber    int    `json:"blockNumber"`
	BlockTimestamp int    `json:"blockTimestamp"`
	TokenAddress   string `json:"tokenAddress"`
	Name           string `json:"name"`
	Symbol         string `json:"symbol"`
	Decimals       int    `json:"decimals"`
}

func (c *Client) GetTokenMeta(payload GetTokenMetaPayload) (response GetTokenMetaResponse, err error) {

	slug := CREDMARK_API_V1_TOKEN_META

	if err := ValidateStruct(payload); err != nil {
		log.Error(err, string(debug.Stack()))
		return response, err
	}

	endpoint := fmt.Sprintf(slug, strconv.Itoa(payload.ChainID), payload.TokenAddress)
	if err != nil {
		log.Error(err, string(debug.Stack()))
		return response, err
	}

	apiUrl := CREDMARK_GATEWAY_URL + endpoint
	queryUrl, err := url.Parse(apiUrl)
	if err != nil {
		log.Error(err, string(debug.Stack()))
		return response, err
	}

	values := queryUrl.Query()

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

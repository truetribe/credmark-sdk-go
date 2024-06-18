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
	CREDMARK_API_V1_URI_TOKEN_HOLDER_COUNT = "/v1/tokens/%s/%s/holders/count" //chainId, tokenAddress
)

type GetTokenHoldersCountPayload struct {
	ChainID      int    `validate:"required"`
	TokenAddress string `validate:"required"`
	BlockNumber  *uint64
	Timestamp    *uint64
}

type TokenHoldersCountResponse struct {
	ChainID      int    `json:"chainId"`
	BlockNumber  int64  `json:"blockNumber"`
	Timestamp    int64  `json:"blockTimestamp"`
	TokenAddress string `json:"tokenAddress"`
	Count        int64  `json:"count"`
}

func (c *Client) GetTokenHoldersCount(payload GetTokenHoldersCountPayload) (response TokenHoldersCountResponse, err error) {

	if err := ValidateStruct(payload); err != nil {
		log.Error(err, string(debug.Stack()))
		return response, err
	}

	endpoint := fmt.Sprintf(CREDMARK_API_V1_URI_TOKEN_HOLDER_COUNT, strconv.Itoa(payload.ChainID), payload.TokenAddress)
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
		return response, fmt.Errorf("%v: Response Error: %v", CREDMARK_API_V1_URI_TOKEN_HOLDER_COUNT, body)
	}

	err = json.NewDecoder(body).Decode(&response)
	if err != nil {
		return response, fmt.Errorf("%v: Decode Error: %v", CREDMARK_API_V1_URI_TOKEN_HOLDER_COUNT, err)
	}

	return
}

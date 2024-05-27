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
	CREDMARK_TOKEN_API_V1_VOLUME = "/v1/tokens/%s/%s/volume" //chainId, tokenAddress
)

type GetTokenVolumePayload struct {
	ChainID          int    `validate:"required"`
	TokenAddr        string `validate:"required"`
	Scaled           bool
	StartBlockNumber *uint64
	EndBlockNumber   *uint64
	StartTimestamp   *uint64
	EndTimestamp     *uint64
}

type GetTokenVolumeResponse struct {
	ChainID          int     `json:"chainId"`
	StartBlockNumber int     `json:"startBlockNumber"`
	EndBlockNumber   int     `json:"endBlockNumber"`
	StartTimestamp   int     `json:"startTimestamp"`
	EndTimestamp     int     `json:"endTimestamp"`
	TokenAddress     string  `json:"tokenAddress"`
	Scaled           bool    `json:"scaled"`
	Volume           float64 `json:"volume"`
}

func (c *Client) GetTokenVolume(payload GetTokenVolumePayload) (response GetTokenVolumeResponse, err error) {

	if err := ValidateStruct(payload); err != nil {
		log.Error(err, string(debug.Stack()))
		return response, err
	}

	uri := fmt.Sprintf(CREDMARK_API_V1_URI_TOKEN_PRICE, strconv.Itoa(payload.ChainID), payload.TokenAddr)
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

	if payload.StartBlockNumber != nil {
		values.Add("startBlockNumber", strconv.FormatUint(*payload.StartBlockNumber, 10))
	}

	if payload.EndBlockNumber != nil {
		values.Add("endBlockNumber", strconv.FormatUint(*payload.EndBlockNumber, 10))
	}

	if payload.StartTimestamp != nil {
		values.Add("startTimestamp", strconv.FormatUint(*payload.StartTimestamp, 10))
	}

	if payload.EndTimestamp != nil {
		values.Add("endTimestamp", strconv.FormatUint(*payload.EndTimestamp, 10))
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
		return response, fmt.Errorf("%v: Response Error: %v", CREDMARK_API_V1_URI_TOKEN_PRICE, body)
	}

	err = json.NewDecoder(body).Decode(&response)
	if err != nil {
		return response, fmt.Errorf("%v: Decode Error: %v", CREDMARK_API_V1_URI_TOKEN_PRICE, err)
	}

	return
}

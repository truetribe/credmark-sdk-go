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
	CREDMARK_API_V1_URI_TOKEN_VOLUME_HISTORICAL = "/v1/tokens/%s/%s/volume/historical" //chainId, tokenAddress
)

type GetTokenVolumeHistoricalPayload struct {
	ChainID          int    `validate:"required"`
	TokenAddress     string `validate:"required"`
	StartBlockNumber *uint64
	EndBlockNumber   *uint64
	BlockInterval    *uint64
	StartTimestamp   *uint64
	EndTimestamp     *uint64
	TimeInterval     *uint64
}

type TokenVolumeHistoricalResponse struct {
	ChainID          int                         `json:"chainId"`
	StartBlockNumber int                         `json:"startBlockNumber"`
	EndBlockNumber   int                         `json:"endBlockNumber"`
	StartTimestamp   int                         `json:"startTimestamp"`
	EndTimestamp     int                         `json:"endTimestamp"`
	TokenAddress     string                      `json:"tokenAddress"`
	Scaled           bool                        `json:"scaled"`
	Data             []TokenVolumeHistoricalItem `json:"data"`
}

type TokenVolumeHistoricalItem struct {
	StartBlockNumber int     `json:"startBlockNumber"`
	EndBlockNumber   int     `json:"endBlockNumber"`
	StartTimestamp   int     `json:"startTimestamp"`
	EndTimestamp     int     `json:"endTimestamp"`
	Volume           float64 `json:"volume"`
}

func (c *Client) GetTokenVolumeHistorical(payload GetTokenVolumeHistoricalPayload) (response TokenVolumeHistoricalResponse, err error) {

	slug := CREDMARK_API_V1_URI_TOKEN_VOLUME_HISTORICAL

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

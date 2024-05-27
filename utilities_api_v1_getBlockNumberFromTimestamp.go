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
	CREDMARK_UTIL_V1_GET_BLOCK_NUM_FROM_TIMESTAMP = "/v1/utilities/chains/%s/block/from-timestamp"
)

type GetBlockNumberFromTimestamp struct {
	ChainID   int    `validate:"required"`
	Timestamp string `validate:"required"`
}

type GetBlockNumberFromTimestampResponse struct {
	BlockNumber    uint `json:"blockNumber"`
	BlockTimestamp uint `json:"blockTimestamp"`
}

func (c *Client) GetBlockNumberFromTimestamp(payload GetBlockNumberFromTimestamp) (response GetBlockNumberFromTimestampResponse, err error) {

	slug := CREDMARK_UTIL_V1_GET_BLOCK_NUM_FROM_TIMESTAMP

	if err := ValidateStruct(payload); err != nil {
		log.Error(err, string(debug.Stack()))
		return response, err
	}

	endpoint := fmt.Sprintf(slug, strconv.Itoa(payload.ChainID))
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

	values.Add("timestamp", payload.Timestamp)

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

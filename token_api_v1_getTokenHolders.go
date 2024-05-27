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
	CREDMARK_API_V1_URI_TOKEN_HOLDER = "/v1/tokens/%s/%s/holders" //chainId, tokenAddress
)

type GetTokenHoldersPayload struct {
	ChainID   int    `validate:"required"` //Chain ID. Use 1 for mainnet.
	TokenAddr string `validate:"required"` //The address of the token requested.
	PageSize  int    `validate:"required"` //The size of the returned page. Do not change this from page to page when using a cursor.
	Cursor    string //The cursor from the results of a previous page. Use empty string (or undefined/null) for first page.
	QuoteAddr string //The address of the token/currency used as the currency of the returned price. Defaults to USD (address 0x0000000000000000000000000000000000000348)
	Scaled    bool   //Scale holders' balance by token decimals. Defaults to true.
	BlockNum  int    //Block number of the price quote. Defaults to the latest block.
	Timestamp int    //Timestamp of a block number can be specified instead of a block number. Finds a block at or before the number of seconds since January 1, 1970.
}

type GetTokenHoldersCredmarkResponse struct {
	ChainID        int    `json:"chainId"`
	BlockNumber    int    `json:"blockNumber"`
	BlockTimestamp int    `json:"blockTimestamp"`
	TokenAddress   string `json:"tokenAddress"`
	Scaled         bool   `json:"scaled"`
	QuoteAddress   string `json:"quoteAddress"`
	Data           []struct {
		Address string  `json:"address"`
		Balance float64 `json:"balance"`
		Value   float64 `json:"value"`
	} `json:"data"`
	Total  int    `json:"total"`
	Cursor string `json:"cursor"`
}

func (c *Client) GetTokenHolders(payload GetTokenHoldersPayload) (response GetTokenHoldersCredmarkResponse, err error) {

	if err := ValidateStruct(payload); err != nil {
		log.Error(err, string(debug.Stack()))
		return response, err
	}

	endpoint := fmt.Sprintf(CREDMARK_API_V1_URI_TOKEN_HOLDER, strconv.Itoa(payload.ChainID), payload.TokenAddr)
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

	values.Add("pageSize", strconv.Itoa(payload.PageSize))

	if len(payload.Cursor) > 0 {
		values.Add("cursor", payload.Cursor)
	}

	if len(payload.QuoteAddr) > 0 {
		values.Add("quoteAddress", payload.QuoteAddr)
	}

	if !payload.Scaled {
		values.Add("scaled", "false")
	}

	if payload.BlockNum > 0 {
		values.Add("blockNumber", strconv.Itoa(payload.BlockNum))
	}

	if payload.Timestamp > 0 {
		values.Add("timestamp", strconv.Itoa(payload.Timestamp))
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
		return response, fmt.Errorf("%v: Response Error: %v", CREDMARK_API_V1_URI_TOKEN_HOLDER, body)
	}

	err = json.NewDecoder(body).Decode(&response)
	if err != nil {
		return response, fmt.Errorf("%v: Decode Error: %v", CREDMARK_API_V1_URI_TOKEN_HOLDER, err)
	}

	return
}

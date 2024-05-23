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
	CREDMARK_API_V1_URI_TOKEN_PRICE = "/v1/tokens/%s/%s/price" //chainId, tokenAddress
)

type GetTokenPricePayload struct {
	ChainId   int    `validate:"required"` //Chain ID. Use 1 for mainnet.
	TokenAddr string `validate:"required"` //The address of the token requested.
	QuoteAddr string //The address of the token/currency used as the currency of the returned price. Defaults to USD (address 0x0000000000000000000000000000000000000348)
	BlockNum  int    //Block number of the price quote. Defaults to the latest block.
	Timestamp int    //Timestamp of a block number can be specified instead of a block number. Finds a block at or before the number of seconds since January 1, 1970.
	Source    string `validate:"oneof='' 'dex' 'cex'"` //(Optional) Available values : dex, cex, specify preferred source to be queried first, choices: "dex" (pre-calculated, default), or "cex" (from call to price.quote model)
	Align     string //(Optional) Available values : 5min, 15min; specify align block number or timestamp to nearest mark for pre-calculated price, choose "5min" for 0,5,10...45,50,55 minutes in every hour
}

type GetTokenPriceCredmarkResponse struct {
	ChainID        int     `json:"chainId"`
	BlockNumber    int     `json:"blockNumber"`
	BlockTimestamp int     `json:"blockTimestamp"`
	TokenAddress   string  `json:"tokenAddress"`
	QuoteAddress   string  `json:"quoteAddress"`
	Price          float64 `json:"price"`
	Src            string  `json:"src"`
	SrcInternal    string  `json:"srcInternal"`
}

type GetTokenPriceCredmarkResponseError struct {
	StatusCode int      `json:"statusCode"`
	Error      string   `json:"error"`
	Message    []string `json:"message"`
}

func (c *Client) GetTokenPrice(payload GetTokenPricePayload) (response GetTokenPriceCredmarkResponse, err error) {

	if err := ValidateStruct(payload); err != nil {
		log.Error(err, string(debug.Stack()))
		return response, err
	}

	uri := fmt.Sprintf(CREDMARK_API_V1_URI_TOKEN_PRICE, strconv.Itoa(payload.ChainId), payload.TokenAddr)
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

	if len(payload.QuoteAddr) > 0 {
		values.Add("quoteAddress", payload.QuoteAddr)
	}

	if payload.BlockNum > 0 {
		values.Add("blockNumber", strconv.Itoa(payload.BlockNum))
	}

	if payload.Timestamp > 0 {
		values.Add("timestamp", strconv.Itoa(payload.Timestamp))
	}

	if len(payload.Source) > 0 {
		values.Add("src", payload.Source)
	}

	if len(payload.Align) > 0 {
		values.Add("align", payload.Align)
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

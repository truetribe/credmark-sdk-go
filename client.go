package credmark

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/hashicorp/go-retryablehttp"

	log "github.com/sirupsen/logrus"
)

const (
	CLIENT_VERSION = "0.0.0"

	CREDMARK_API_V1_XR_URI = "/v1/model/run"
)

type ApiRequest struct {
	Endpoint string `validate:"required"`
	Method   string `validate:"required"`
	Body     string `validate:"required"`
	Header   RequestHeader
}

type RequestHeader struct {
	ContentType   string
	Authorization string
	Connection    string
}

type Client struct {
	HttpClient          *http.Client
	Config              *Config
	RetryablehttpClient *retryablehttp.Client
}

func NewClient(cfg *Config) (*Client, error) {

	if err := ValidateStruct(cfg); err != nil {
		log.Error(err, string(debug.Stack()))
		return nil, err
	}

	retryClient := retryablehttp.NewClient()
	if cfg.RetryWaitMin > 0 {
		retryClient.RetryWaitMin = time.Duration(cfg.RetryWaitMin) * time.Second
	}

	if cfg.RetryWaitMax > 0 {
		retryClient.RetryWaitMax = time.Duration(cfg.RetryWaitMax) * time.Second
	}

	if cfg.RetryMax > 0 {
		retryClient.RetryMax = cfg.RetryMax
	}

	client := &Client{
		Config:              cfg,
		HttpClient:          &http.Client{},
		RetryablehttpClient: retryClient,
	}

	if err := ValidateStruct(client); err != nil {
		log.Error(err, string(debug.Stack()))
		return nil, err
	}

	return client, nil
}

// build request and sent
func (c *Client) doHttpRequeset(apiReq ApiRequest, actionDesc string) (*http.Response, error) {
	//build request
	req, err := retryablehttp.NewRequest(apiReq.Method, apiReq.Endpoint, bytes.NewReader([]byte(apiReq.Body)))
	if err != nil {
		return nil, fmt.Errorf("%v: Not able to buildHttpRequest: %v", actionDesc, err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.Header.Set("Authorization", "Bearer "+c.Config.ApiKey)

	//send request
	res, err := c.RetryablehttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%v: Not able to send request: %v", actionDesc, err)
	}

	return res, nil
}

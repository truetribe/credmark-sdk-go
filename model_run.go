package credmark

import (
	"fmt"
	"io"
	"net/http"
)

type CredmarkModelResponse struct {
	Slug         string      `json:"slug"`
	Version      string      `json:"version"`
	ChainID      int         `json:"chainId"`
	BlockNumber  int         `json:"blockNumber"`
	Output       interface{} `json:"output"`
	Dependencies interface{} `json:"dependencies"`
	Cached       bool        `json:"cached"`
	Runtime      int         `json:"runtime"`
}

func (c *Client) RunModel(payload string) (response string, err error) {

	apiReq := ApiRequest{
		Endpoint: CREDMARK_GATEWAY_URL + CREDMARK_API_V1_XR_URI,
		Method:   "POST",
		Body:     payload,
	}

	res, err := c.doHttpRequeset(apiReq, GetCurrentFuncName())
	if err != nil {
		return response, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	response = string(body)

	if res.StatusCode != http.StatusOK {
		return response, fmt.Errorf("credmark model error http response status code: %v, %v", res.StatusCode, response)
	}

	return response, err
}

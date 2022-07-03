package aladhan

import (
	"encoding/json"
	"net/http"

	"github.com/gojek/heimdall/v7/httpclient"
)

const defaultCountry = "Russia"

type Client struct {
	Host   string
	Path   string
	Client *httpclient.Client
}

func (c *Client) GetTimeByCity(city string) (*Response, error) {
	return c.sendRequest(NewRequest(city, defaultCountry))
}

func (c *Client) sendRequest(request *Request) (*Response, error) {
	fullUrl := c.getFullUrl(request)
	req, err := http.NewRequest(request.httpMethod(), fullUrl, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	response := &Response{}

	if err = json.NewDecoder(resp.Body).Decode(response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) getFullUrl(request *Request) string {
	return c.Host + c.Path + request.getApiMethod() + request.string()
}

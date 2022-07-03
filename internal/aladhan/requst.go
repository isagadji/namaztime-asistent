package aladhan

import (
	"fmt"
	"net/http"
)

type Request struct {
	Country string `json:"country"`
	City    string `json:"address"`
	Method  int    `json:"method"`
}

func NewRequest(city string, country string) *Request {
	return &Request{
		Country: country,
		City:    city,
		Method:  1, //3 - Muslim World League, 14 - Spiritual Administration of Muslims of Russia
	}
}

func (receiver Request) httpMethod() string {
	return http.MethodGet
}

func (receiver Request) getApiMethod() string {
	return "timingsByCity"
}

func (r Request) string() string {
	return fmt.Sprintf("?city=%s&country=%s&method=%d", r.City, r.Country, r.Method)
}

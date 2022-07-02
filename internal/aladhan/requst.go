package aladhan

import (
	"net/http"
)

type Request struct {
	Country string `json:"country"`
	City    string `json:"address"`
	Method  int    `json:"method"`
}

func NewRequest(city string, country string) Request {
	return Request{
		Country: country,
		City:    city,
		Method:  1,
	}
}

func (receiver Request) httpMethod() string {
	return http.MethodGet
}

func (receiver Request) getPath() string {
	return "timingsByCity"
}

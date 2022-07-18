package aladhan

import (
	"time"

	"github.com/gojek/heimdall/v7/httpclient"
)

type Flags struct {
	APIHost    string        `kong:"required,name=aladhan-api-host,env=ALADHAN_API_HOST"`
	APIPath    string        `kong:"required,name=aladhan-api-path,env=ALADHAN_API_PATH"`
	APITimeout time.Duration `kong:"optional,name=aladhan-api-timeout,env=ALADHAN_API_TIMEOUT,default=1s"`
}

const defaultRetryCount = 1

func (f Flags) Init() *Service {
	client := &Client{
		Host: f.APIHost,
		Path: f.APIPath,
		Client: httpclient.NewClient(
			httpclient.WithHTTPTimeout(f.APITimeout),
			httpclient.WithRetryCount(defaultRetryCount),
		),
	}

	return NewService(client)
}

package httpclient

import (
	"crypto/tls"
	"time"

	"github.com/go-resty/resty/v2"
)

type HttpClient struct {
	c *resty.Client
}

func (h *HttpClient) GetClient() *resty.Client {

	return h.c
}

func (h *HttpClient) SetClient(c *resty.Client) {

	h.c = c
}

func NewHttpClient() *HttpClient {

	return &HttpClient{
		c: resty.New().SetDebug(false).SetTimeout(time.Second * 5).SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}),
	}
}

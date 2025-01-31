package httpclient

import (
	"github.com/go-resty/resty/v2"
)

func (h *HttpClient) Get(url string, params map[string]string) (*resty.Response, error) {

	return h.c.R().SetQueryParams(params).Get(url)
}

func (h *HttpClient) Post(url string, body map[string]string) (*resty.Response, error) {

	return h.c.R().SetQueryParams(body).Post(url)
}

func (h *HttpClient) PostData(url string, data interface{}) (*resty.Response, error) {

	return h.c.R().SetBody(data).Post(url)
}

func (h *HttpClient) GetWithBasicAuth(url string, params map[string]string, username string, password string) (*resty.Response, error) {

	return h.c.R().SetBasicAuth(username, password).SetQueryParams(params).Get(url)
}

func (h *HttpClient) PostWithBasicAuth(url string, body map[string]string, username string, password string) (*resty.Response, error) {

	return h.c.R().SetBasicAuth(username, password).SetQueryParams(body).Post(url)
}

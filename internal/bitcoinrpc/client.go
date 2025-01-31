package bitcoinrpc

import (
	"crypto/tls"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

type BitcoinRpc struct {
	protocol string // http
	host     string // 127.0.0.1
	port     string // 8545
	user     string // user
	pass     string // pass
}

func (br *BitcoinRpc) url() string {

	return fmt.Sprintf("%s://%s:%s", br.protocol, br.host, br.port)
}

func (br *BitcoinRpc) request() *resty.Request {

	return resty.New().SetDebug(false).SetTimeout(time.Second * 10).SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).R()
}

func NewBitcoinRPC(protocol, host, port, user, pass string) *BitcoinRpc {

	return &BitcoinRpc{
		protocol: protocol,
		host:     host,
		port:     port,
		user:     user,
		pass:     pass,
	}
}

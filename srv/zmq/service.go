package zmq

import (
	"btczmq/config"
	"btczmq/srv/gateway"
	"context"
)

type Zmq struct {
	address string
	ctx     context.Context
	g       *gateway.Gateway
}

func (z *Zmq) Start() {

	go z.g.Start()
}

func NewZmq(ctx context.Context, g *gateway.Gateway) *Zmq {

	return &Zmq{
		ctx:     ctx,
		g:       g,
		address: config.ZmqAddress(),
	}
}

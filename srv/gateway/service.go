package gateway

import (
	"btczmq/config"
	"btczmq/internal/wsserver"
	"btczmq/tools/logger"
	"context"
	"net"
	"sync"
	"time"
)

type Gateway struct {
	ctx      context.Context
	cancel   context.CancelFunc
	wsserver *wsserver.WsServer
	// state
	mutex       sync.Mutex
	connections []net.Conn
}

func (g *Gateway) Start() error {

	// start ws server to access ws connections
	err := g.wsserver.Serve(g.ctx)
	if err != nil {

		return err
	}

	logger.Info("[gateway] wsserver started").Log()

	return nil
}

func (g *Gateway) Stop() {

	g.cancel()
	time.Sleep(time.Second * 3) // wait for all go routines to stop
}

func NewGateway(ctx context.Context, cancel context.CancelFunc) *Gateway {

	g := &Gateway{
		ctx:      ctx,
		cancel:   cancel,
		wsserver: nil,
	}

	g.wsserver = wsserver.NewWsServer(config.GatewayHost(), config.GatewayPort(), g.validateConnection, g.onConnectionOpenned, g.onConnectionClosed, g.onConnectionMessage)

	return g
}

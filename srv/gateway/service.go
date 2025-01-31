package gateway

import (
	"btczmq/config"
	"btczmq/internal/wsserver"
	"btczmq/tools/logger"
	"btczmq/types"
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
	connectionsMutex  sync.Mutex
	connections       []net.Conn
	transactionsMutex sync.Mutex
	transactions      []types.Transaction
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
		ctx:          ctx,
		cancel:       cancel,
		wsserver:     nil,
		transactions: make([]types.Transaction, 0),
	}

	g.wsserver = wsserver.NewWsServer(config.GatewayHost(), config.GatewayPort(), g.validateConnection, g.onConnectionOpenned, g.onConnectionClosed, g.onConnectionMessage)

	return g
}

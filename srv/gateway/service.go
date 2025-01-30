package gateway

import (
	"btczmq/config"
	"btczmq/internal/wsserver"
	"btczmq/tools/logger"
	"btczmq/types"
	"context"
	"net"
	"time"
)

type Gateway struct {
	ctx      context.Context
	cancel   context.CancelFunc
	wsserver *wsserver.WsServer
	// jobs
	addConnectionJobs    chan net.Conn
	removeConnectionJobs chan net.Conn
	newTransactionJobs   chan types.Transaction
	// state
	connections []net.Conn
}

func (g *Gateway) Start() error {

	// start ws server to access ws connections
	err := g.wsserver.Serve(g.ctx)
	if err != nil {

		return err
	}

	logger.Info("[gateway] wsserver started").Log()

	// start new go routines to handle jobs
	go g.handleJobs()

	return nil
}

func (g *Gateway) Stop() {

	g.cancel()
	time.Sleep(time.Second * 3) // wait for all go routines to stop
}

func NewGateway(ctx context.Context, cancel context.CancelFunc) *Gateway {

	g := &Gateway{
		ctx:                  ctx,
		cancel:               cancel,
		wsserver:             nil,
		addConnectionJobs:    make(chan net.Conn, 100),
		removeConnectionJobs: make(chan net.Conn, 100),
		newTransactionJobs:   make(chan types.Transaction, 100),
	}

	g.wsserver = wsserver.NewWsServer(config.GatewayHost(), config.GatewayPort(), g.validateConnection, g.onConnectionOpenned, g.onConnectionClosed, g.onConnectionMessage)

	return g
}

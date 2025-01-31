package gateway

import (
	"btczmq/types"
	"net"
)

func (g *Gateway) BroadcastTransction(transaction types.Transaction) {

	g.transactionsMutex.Lock()
	defer g.transactionsMutex.Unlock()

	// always keep last 1000 transactions
	g.transactions = append(g.transactions, transaction)
	if len(g.transactions) > 1000 {

		g.transactions = g.transactions[1:]
	}

	g.connectionsMutex.Lock()
	defer g.connectionsMutex.Unlock()

	for _, connection := range g.connections {

		g.wsserver.SendServerMessage(connection, types.NewServerMessage(1, transaction))
	}
}

func (g *Gateway) AddConnection(conn net.Conn) {

	g.connectionsMutex.Lock()
	defer g.connectionsMutex.Unlock()

	g.connections = append(g.connections, conn)
}

func (g *Gateway) RemoveConnection(conn net.Conn) {

	g.connectionsMutex.Lock()
	defer g.connectionsMutex.Unlock()

	for i, client := range g.connections {

		if client == conn {

			g.connections = append(g.connections[:i], g.connections[i+1:]...)
			break
		}
	}
}

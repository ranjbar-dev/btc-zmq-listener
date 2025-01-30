package gateway

import (
	"btczmq/types"
	"net"
)

func (g *Gateway) BroadcastTransction(transaction types.Transaction) {

	g.mutex.Lock()
	defer g.mutex.Unlock()

	for _, connection := range g.connections {

		g.wsserver.SendServerMessage(connection, types.NewServerMessage(1, transaction))
	}
}

func (g *Gateway) AddConnection(conn net.Conn) {

	g.mutex.Lock()
	defer g.mutex.Unlock()

	g.connections = append(g.connections, conn)
}

func (g *Gateway) RemoveConnection(conn net.Conn) {

	g.mutex.Lock()
	defer g.mutex.Unlock()

	for i, client := range g.connections {

		if client == conn {

			g.connections = append(g.connections[:i], g.connections[i+1:]...)
			break
		}
	}
}

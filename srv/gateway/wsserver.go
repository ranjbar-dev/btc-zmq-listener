package gateway

import (
	"btczmq/types"
	"fmt"
	"net"
)

func (g *Gateway) validateConnection(conn net.Conn) bool {

	fmt.Println("LocalAddr", conn.LocalAddr().String())
	fmt.Println("RemoteAddr", conn.RemoteAddr().String())

	return true
}

func (g *Gateway) onConnectionOpenned(conn net.Conn) {

	g.AddConnection(conn)
}

func (g *Gateway) onConnectionClosed(conn net.Conn) {

	g.RemoveConnection(conn)
}

func (g *Gateway) onConnectionMessage(conn net.Conn, msg []byte) {

	g.wsserver.SendServerMessage(conn, types.NewServerMessage(-1, "rpc not implemented"))
}

package gateway

import (
	"btczmq/config"
	"btczmq/tools/logger"
	"btczmq/types"
	"net"
	"strings"
)

func (g *Gateway) validateConnection(conn net.Conn) bool {

	splitted := strings.Split(conn.RemoteAddr().String(), ":")
	ipAddress := splitted[0]

	for _, validIpAddress := range config.GatewayWhiteListIps() {

		if ipAddress == validIpAddress {

			logger.Debug("[gateway] accepted new connection from ip address " + ipAddress).Log()
			return true
		}
	}

	logger.Debug("[gateway] rejected new connection from ip address " + ipAddress).Log()
	return false
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

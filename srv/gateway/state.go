package gateway

import "net"

func (g *Gateway) AddConnection(conn net.Conn) {

	go func() {

		g.addConnectionJobs <- conn
	}()
}

func (g *Gateway) RemoveConnection(conn net.Conn) {

	go func() {

		g.removeConnectionJobs <- conn
	}()
}

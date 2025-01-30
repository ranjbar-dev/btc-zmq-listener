package wsserver

import (
	"btczmq/tools/logger"
	"context"
	"net"
)

type WsServer struct {
	terminating         bool
	host                string
	port                string
	validateConnection  func(conn net.Conn) bool
	onConnectionOpenned func(conn net.Conn)
	onConnectionClosed  func(conn net.Conn)
	onConnectionMessage func(conn net.Conn, msg []byte)
}

func (s *WsServer) Serve(ctx context.Context) error {

	listener, err := net.Listen("tcp", s.host+":"+s.port)
	if err != nil {
		return err
	}

	// handle incoming connections
	go func() {

		for {

			if s.terminating {
				return
			}

			conn, err := s.listenForConnections(listener)
			if err != nil {

				logger.Warn("wsserver upgrade error").Message(err.Error()).Log()
				continue
			}

			s.onConnectionOpenned(conn)

			go s.handleConnection(conn)
		}
	}()

	// handle context done
	go func() {

		<-ctx.Done()

		s.terminating = true

		err := listener.Close()
		if err != nil {

			logger.Error("wsserver close listener error").Message(err.Error()).Log()
		}
	}()

	return nil
}

func NewWsServer(host string, port string, validateConnection func(conn net.Conn) bool, onConnectionOpenned func(conn net.Conn), onConnectionClosed func(conn net.Conn), onConnectionMessage func(conn net.Conn, msg []byte)) *WsServer {

	return &WsServer{
		terminating:         false,
		host:                host,
		port:                port,
		validateConnection:  validateConnection,
		onConnectionOpenned: onConnectionOpenned,
		onConnectionClosed:  onConnectionClosed,
		onConnectionMessage: onConnectionMessage,
	}
}

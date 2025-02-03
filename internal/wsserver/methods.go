package wsserver

import (
	"btczmq/tools/logger"
	"fmt"
	"io"
	"net"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func (s *WsServer) listenForConnections(listener net.Listener) (net.Conn, error) {

	// accept connection
	conn, err := listener.Accept()
	if err != nil {

		return nil, err
	}

	// validate connection
	if !s.validateConnection(conn) {

		conn.Close()
		return nil, fmt.Errorf("connection validation failed")
	}

	// upgrade the connection to a WebSocket connection
	_, err = ws.Upgrade(conn)
	if err != nil {

		conn.Close()
		return nil, err
	}

	return conn, nil
}

func (s *WsServer) handleConnection(conn net.Conn) {

	defer func() {

		s.onConnectionClosed(conn)
		conn.Close()
	}()

	for {

		msg, op, err := wsutil.ReadClientData(conn)

		if err != nil {

			errorMessage := err.Error()

			// disconnect by client
			if err == io.EOF || errorMessage == "ws closed: 1000 " {

				return
			}

			// disconnect by server
			if errorMessage == "ws closed: 1006 " {

				return
			}

			logger.Error("wsserver error client data").Message(err.Error()).Log()
			return
		}

		if op == ws.OpText {

			s.onConnectionMessage(conn, msg)
		}
	}

}

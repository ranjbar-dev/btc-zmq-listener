package wsserver

import (
	"btczmq/types"
	"encoding/json"
	"net"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func (s *WsServer) SendPlainMessage(conn net.Conn, bytes []byte) error {

	return wsutil.WriteServerMessage(conn, ws.OpText, bytes)
}

func (s *WsServer) SendBinaryMessage(conn net.Conn, bytes []byte) error {

	return wsutil.WriteServerBinary(conn, bytes)
}

func (s *WsServer) SendServerMessage(conn net.Conn, message types.ServerMessage) error {

	data, err := json.Marshal(message)
	if err != nil {

		return err
	}

	return s.SendPlainMessage(conn, data) // TODO : change to binary message later
}

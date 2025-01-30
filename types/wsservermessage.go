package types

type ServerMessage struct {
	Code int `json:"c"`
	Data any `json:"d"`
}

func NewServerMessage(code int, data any) ServerMessage {

	return ServerMessage{
		Code: code,
		Data: data,
	}
}

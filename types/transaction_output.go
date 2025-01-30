package types

type TransactionInput struct {
	TxID    string `json:"txid"`
	Index   uint32 `json:"index"`
	Address string `json:"address"`
}

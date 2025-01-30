package types

type Transaction struct {
	TxID          string              `json:"txid"`
	Inputs        []TransactionInput  `json:"inputs"`
	Outputs       []TransactionOutput `json:"outputs"`
	Confirmations int                 `json:"confirmations"`
}

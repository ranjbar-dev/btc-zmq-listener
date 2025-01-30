package types

type TransactionOutput struct {
	Address string `json:"address"`
	Value   int64  `json:"value"`
}

package gateway

import "btczmq/types"

func (g *Gateway) NotifyTransactionToConnections(transaction types.Transaction) {

	go func() {

		g.newTransactionJobs <- transaction
	}()
}

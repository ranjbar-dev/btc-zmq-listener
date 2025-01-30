package zmq

import (
	"btczmq/tools/decoder"
	"btczmq/tools/logger"
	"encoding/hex"
	"time"

	"github.com/go-zeromq/zmq4"
)

func (z *Zmq) SubscribeToNewTransactions() {

	logger.Info("[zmq] subscribing to new transactions").Log()

	// Create a new SUB socket
	subscriber := zmq4.NewSub(z.ctx)
	defer subscriber.Close()

	// Connect to the publisher
	err := subscriber.Dial(z.address)
	if err != nil {

		logger.Info("[zmq] failed to dial").Message(err.Error()).Log()
		return
	}

	// Subscribe to all messages (empty topic)
	err = subscriber.SetOption(zmq4.OptionSubscribe, "")
	if err != nil {

		logger.Info("[zmq] failed to subscribe").Message(err.Error()).Log()
		return
	}

	logger.Info("[zmq] subscribed to new transactions").Log()

	// Loop to receive messages
	for {

		// Receive a message
		msg, err := subscriber.Recv()
		if err != nil {

			logger.Error("[zmq] failed to receive message").Message(err.Error()).Log()
			continue
		}

		// Convert the raw transaction data to hexadecimal
		rawTxHex := hex.EncodeToString(msg.Frames[1]) // The actual raw transaction is in the second frame

		transaction, err := decoder.DecodeTransactionHex(rawTxHex)
		if err != nil {

			logger.Error("[zmq] failed to decode transaction").Message(err.Error()).Log()
			continue
		}

		// Notify the transaction to all connections
		z.g.NotifyTransactionToConnections(transaction)

		// Sleep for a short duration to avoid busy-waiting
		time.Sleep(10 * time.Millisecond)
	}
}

package decoder

import (
	"btczmq/types"
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
)

func DecodeTransactionHex(rawTxHex string) (types.Transaction, error) {

	var transaction types.Transaction

	txBytes, err := hex.DecodeString(rawTxHex)
	if err != nil {

		return transaction, fmt.Errorf("error decoding hex: %v", err)
	}

	var tx wire.MsgTx
	err = tx.Deserialize(bytes.NewReader(txBytes))
	if err != nil {

		return transaction, fmt.Errorf("error deserializing transaction: %v", err)
	}

	var inputs []types.TransactionInput
	for _, input := range tx.TxIn {

		txID := input.PreviousOutPoint.Hash.String()
		index := input.PreviousOutPoint.Index

		// Extract address from input
		var address string
		script := input.SignatureScript
		witness := input.Witness

		// Try different extraction methods
		if len(witness) > 0 {

			// Handle SegWit inputs
			if len(witness) >= 2 {

				pubKeyHash := btcutil.Hash160(witness[len(witness)-1])
				addr, err := btcutil.NewAddressWitnessPubKeyHash(pubKeyHash, &chaincfg.MainNetParams)
				if err == nil {

					address = addr.String()
				}
			}
		} else if len(script) > 0 {

			// Handle legacy inputs
			if data, err := txscript.PushedData(script); err == nil && len(data) >= 2 {

				if pubKeyHash := btcutil.Hash160(data[len(data)-1]); len(pubKeyHash) == 20 {

					if addr, err := btcutil.NewAddressPubKeyHash(pubKeyHash, &chaincfg.MainNetParams); err == nil {

						address = addr.String()
					}
				}
			}
		}

		inputs = append(inputs, types.TransactionInput{
			TxID:    txID,
			Index:   index,
			Address: address,
		})
	}

	var outputs []types.TransactionOutput
	for _, output := range tx.TxOut {

		// Extract address from output
		var address string
		_, addresses, _, err := txscript.ExtractPkScriptAddrs(output.PkScript, &chaincfg.MainNetParams)
		if err == nil && len(addresses) != 0 {

			address = addresses[0].String()
		}

		outputs = append(outputs, types.TransactionOutput{
			Address: address,
			Value:   output.Value,
		})
	}

	transaction.TxID = tx.TxHash().String()
	transaction.Inputs = inputs
	transaction.Outputs = outputs

	return transaction, nil
}

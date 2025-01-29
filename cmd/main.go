package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/go-zeromq/zmq4"
)

type RpcRequest struct {
	Method string   `json:"method"`
	Params []string `json:"params"`
	ID     int      `json:"id"`
}

type RpcResponse struct {
	Result interface{} `json:"result"`
	Error  interface{} `json:"error"`
	ID     int         `json:"id"`
}

func main() {

	// Create a new ZeroMQ context
	ctx := context.Background()

	// Create a new SUB socket
	subscriber := zmq4.NewSub(ctx)
	defer subscriber.Close()

	// Connect to the publisher
	err := subscriber.Dial("tcp://127.0.0.1:28332")
	if err != nil {

		log.Fatalf("Failed to dial: %v", err)
	}

	// Subscribe to all messages (empty topic)
	err = subscriber.SetOption(zmq4.OptionSubscribe, "")
	if err != nil {

		log.Fatalf("Failed to subscribe: %v", err)
	}

	fmt.Println("Subscriber connected to tcp://127.0.0.1:28332 and waiting for messages...")

	// Loop to receive messages
	for {

		// Receive a message
		msg, err := subscriber.Recv()
		if err != nil {
			log.Fatalf("Failed to receive message: %v", err)
		}

		// Convert the raw transaction data to hexadecimal
		rawTxHex := hex.EncodeToString(msg.Frames[1]) // The actual raw transaction is in the second frame
		fmt.Println("Raw Transaction in Hexadecimal:")
		fmt.Println(rawTxHex)

		// Call Bitcoin Core's decoderawtransaction RPC method
		decodedTx, err := decodeRawTransaction(rawTxHex)
		if err != nil {

			log.Fatalf("Failed to decode transaction: %v", err)
		}

		// Print the decoded transaction (JSON)
		fmt.Println("Decoded Transaction:")
		fmt.Println(decodedTx)

		// Sleep for a short duration to avoid busy-waiting
		time.Sleep(100 * time.Millisecond)
	}
}

func decodeRawTransaction(rawTxHex string) (interface{}, error) {
	// Prepare the RPC request payload
	payload := RpcRequest{
		Method: "decoderawtransaction",
		Params: []string{rawTxHex},
		ID:     1,
	}

	// Marshal the request to JSON
	reqBody, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}

	// Prepare the HTTP request to the Bitcoin Core RPC server with authentication
	url := "http://127.0.0.1:8332" // Bitcoin Core RPC URL
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Set RPC credentials (user:password)
	rpcUser := "kahsdasgjasd"
	rpcPassword := "asdovegiyjsd"
	req.SetBasicAuth(rpcUser, rpcPassword)

	// Set necessary headers
	req.Header.Set("Content-Type", "application/json")

	// Make the HTTP request to Bitcoin Core RPC server
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send RPC request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Unmarshal the response JSON
	var rpcResp RpcResponse
	if err := json.Unmarshal(respBody, &rpcResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	if rpcResp.Error != nil {
		return nil, fmt.Errorf("RPC error: %v", rpcResp.Error)
	}

	// Return the decoded transaction
	return rpcResp.Result, nil
}

package main

import (
	"btczmq/srv/gateway"
	"btczmq/srv/zmq"
	"btczmq/tools/logger"
	"context"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	defer func() {

		if err := recover(); err != nil {

			logger.Error("recover function called").Log()
		}
	}()

	// Create a context with cancellation
	ctx, cancel := context.WithCancel(context.Background())

	// Create a channel to receive OS signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	waitChannel := make(chan struct{}, 1)
	go func() {

		// we can exit from app now
		defer func() {

			waitChannel <- struct{}{}
		}()

		// wait for signal to exit
		<-sigs
		cancel()
	}()

	// start gateway service
	gatewayService := gateway.NewGateway(ctx, cancel)
	gatewayService.Start()

	// start zmq service
	zmqService := zmq.NewZmq(ctx, gatewayService)
	zmqService.Start()

	// wait to exit from app
	<-waitChannel
}

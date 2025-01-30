package gateway

import (
	"btczmq/tools/logger"
	"btczmq/types"
)

func (g *Gateway) handleJobs() {

	logger.Info("[gateway] listening for jobs").Log()
	for {

		select {

		// add connection to the list of connections
		case conn := <-g.addConnectionJobs:

			g.connections = append(g.connections, conn)

		// remove connection from the list of connections
		case conn := <-g.removeConnectionJobs:

			for i, client := range g.connections {

				if client == conn {

					g.connections = append(g.connections[:i], g.connections[i+1:]...)
					break
				}
			}

		// send new transaction info to all connections
		case tx := <-g.newTransactionJobs:

			for _, connection := range g.connections {

				g.wsserver.SendServerMessage(connection, types.NewServerMessage(1, tx))
			}

		// handle g.ctx done
		case <-g.ctx.Done():

			logger.Info("[gateway] stopped listening for jobs").Log()
			return
		}
	}
}

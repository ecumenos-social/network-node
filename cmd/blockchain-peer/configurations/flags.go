package configurations

import (
	"time"

	cli "github.com/urfave/cli/v2"
)

var Flags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "nn-blockchain-peer-logger-production",
		Usage:   "make it true if you need logging on production environment",
		Value:   false,
		EnvVars: []string{"NETWORK_NODE_BLOCKCHAIN_PEER_LOGGER_PRODUCTION"},
	},
	&cli.StringFlag{
		Name:    "nn-blockchain-peer-grpc-host",
		Usage:   "it is gRPC server host",
		Value:   "0.0.0.0",
		EnvVars: []string{"NETWORK_NODE_BLOCKCHAIN_PEER_GRPC_HOST"},
	},
	&cli.StringFlag{
		Name:    "nn-blockchain-peer-grpc-port",
		Usage:   "it is gRPC server port",
		Value:   "8080",
		EnvVars: []string{"NETWORK_NODE_BLOCKCHAIN_PEER_GRPC_PORT"},
	},
	&cli.StringFlag{
		Name:    "nn-blockchain-peer-http-gateway-host",
		Usage:   "it is HTTP gateway host",
		Value:   "0.0.0.0",
		EnvVars: []string{"NETWORK_NODE_BLOCKCHAIN_PEER_HTTP_GATEWAY_HOST"},
	},
	&cli.StringFlag{
		Name:    "nn-blockchain-peer-http-gateway-port",
		Usage:   "it is HTTP gateway port",
		Value:   "9090",
		EnvVars: []string{"NETWORK_NODE_BLOCKCHAIN_PEER_HTTP_GATEWAY_PORT"},
	},
	&cli.BoolFlag{
		Name:    "nn-blockchain-peer-enabled-health-server",
		Usage:   "make it true if you need to enable health server",
		Value:   false,
		EnvVars: []string{"NETWORK_NODE_BLOCKCHAIN_PEER_ENABLED_HEALTH_SERVER"},
	},
	&cli.StringFlag{
		Name:    "nn-blockchain-peer-health-server-host",
		Usage:   "it is health server host",
		Value:   "0.0.0.0",
		EnvVars: []string{"NETWORK_NODE_BLOCKCHAIN_PEER_HEALTH_SERVER_HOST"},
	},
	&cli.StringFlag{
		Name:    "nn-blockchain-peer-health-server-port",
		Usage:   "it is health server port",
		Value:   "10010",
		EnvVars: []string{"NETWORK_NODE_BLOCKCHAIN_PEER_HEALTH_SERVER_PORT"},
	},
	&cli.StringFlag{
		Name:    "nn-blockchain-peer-liveness-gateway-host",
		Usage:   "it is liveness gateway host",
		Value:   "0.0.0.0",
		EnvVars: []string{"NETWORK_NODE_BLOCKCHAIN_PEER_LIVENESS_GATEWAY_HOST"},
	},
	&cli.StringFlag{
		Name:    "nn-blockchain-peer-liveness-gateway-port",
		Usage:   "it is liveness gateway port",
		Value:   "8086",
		EnvVars: []string{"NETWORK_NODE_BLOCKCHAIN_PEER_LIVENESS_GATEWAY_PORT"},
	},
	&cli.DurationFlag{
		Name:    "nn-blockchain-peer-grpc-max-conn-age",
		Usage:   "it is max age of connection with gRPC server",
		Value:   5 * time.Minute,
		EnvVars: []string{"NETWORK_NODE_BLOCKCHAIN_PEER_GRPC_MAX_CONNECTION_AGE"},
	},
}

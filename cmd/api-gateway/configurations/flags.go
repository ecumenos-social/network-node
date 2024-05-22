package configurations

import (
	"time"

	cli "github.com/urfave/cli/v2"
)

var Flags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "nn-api-gateway-logger-production",
		Usage:   "make it true if you need logging on production environment",
		Value:   false,
		EnvVars: []string{"NETWORK_NODE_API_GATEWAY_LOGGER_PRODUCTION"},
	},
	&cli.StringFlag{
		Name:    "nn-api-gateway-grpc-host",
		Usage:   "it is gRPC server host",
		Value:   "0.0.0.0",
		EnvVars: []string{"NETWORK_NODE_API_GATEWAY_GRPC_HOST"},
	},
	&cli.StringFlag{
		Name:    "nn-api-gateway-grpc-port",
		Usage:   "it is gRPC server port",
		Value:   "8080",
		EnvVars: []string{"NETWORK_NODE_API_GATEWAY_GRPC_PORT"},
	},
	&cli.StringFlag{
		Name:    "nn-api-gateway-http-gateway-host",
		Usage:   "it is HTTP gateway host",
		Value:   "0.0.0.0",
		EnvVars: []string{"NETWORK_NODE_API_GATEWAY_HTTP_GATEWAY_HOST"},
	},
	&cli.StringFlag{
		Name:    "nn-api-gateway-http-gateway-port",
		Usage:   "it is HTTP gateway port",
		Value:   "9090",
		EnvVars: []string{"NETWORK_NODE_API_GATEWAY_HTTP_GATEWAY_PORT"},
	},
	&cli.BoolFlag{
		Name:    "nn-api-gateway-enabled-health-server",
		Usage:   "make it true if you need to enable health server",
		Value:   false,
		EnvVars: []string{"NETWORK_NODE_API_GATEWAY_ENABLED_HEALTH_SERVER"},
	},
	&cli.StringFlag{
		Name:    "nn-api-gateway-health-server-host",
		Usage:   "it is health server host",
		Value:   "0.0.0.0",
		EnvVars: []string{"NETWORK_NODE_API_GATEWAY_HEALTH_SERVER_HOST"},
	},
	&cli.StringFlag{
		Name:    "nn-api-gateway-health-server-port",
		Usage:   "it is health server port",
		Value:   "10010",
		EnvVars: []string{"NETWORK_NODE_API_GATEWAY_HEALTH_SERVER_PORT"},
	},
	&cli.StringFlag{
		Name:    "nn-api-gateway-liveness-gateway-host",
		Usage:   "it is liveness gateway host",
		Value:   "0.0.0.0",
		EnvVars: []string{"NETWORK_NODE_API_GATEWAY_LIVENESS_GATEWAY_HOST"},
	},
	&cli.StringFlag{
		Name:    "nn-api-gateway-liveness-gateway-port",
		Usage:   "it is liveness gateway port",
		Value:   "8086",
		EnvVars: []string{"NETWORK_NODE_API_GATEWAY_LIVENESS_GATEWAY_PORT"},
	},
	&cli.DurationFlag{
		Name:    "nn-api-gateway-grpc-max-conn-age",
		Usage:   "it is max age of connection with gRPC server",
		Value:   5 * time.Minute,
		EnvVars: []string{"NETWORK_NODE_API_GATEWAY_GRPC_MAX_CONNECTION_AGE"},
	},
}

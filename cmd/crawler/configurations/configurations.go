package configurations

import (
	"github.com/ecumenos-social/toolkitfx/fxgrpc"
	"github.com/ecumenos-social/toolkitfx/fxlogger"
	cli "github.com/urfave/cli/v2"
	"go.uber.org/fx"
)

type fxConfig struct {
	fx.Out

	Logger *fxlogger.Config
	GRPC   *fxgrpc.Config
}

var Module = func(cctx *cli.Context) fx.Option {
	return fx.Options(
		fx.Provide(func() fxConfig {
			return fxConfig{
				Logger: &fxlogger.Config{
					Production: cctx.Bool("nn-crawler-logger-production"),
				},
				GRPC: &fxgrpc.Config{
					GRPC: fxgrpc.GRPCConfig{
						Host:                                    cctx.String("nn-crawler-grpc-host"),
						Port:                                    cctx.String("nn-crawler-grpc-port"),
						MaxConnectionAge:                        cctx.Duration("nn-crawler-grpc-max-conn-age"),
						KeepAliveEnforcementMinTime:             cctx.Duration("nn-crawler-grpc-keep-alive-enforcement-min-time"),
						KeepAliveEnforcementPermitWithoutStream: cctx.Bool("nn-crawler-grpc-keep-alive-enforcement-permit-without-stream"),
					},
					Health: fxgrpc.HealthConfig{
						Enabled: cctx.Bool("nn-crawler-enabled-health-server"),
						Host:    cctx.String("nn-crawler-health-server-host"),
						Port:    cctx.String("nn-crawler-health-server-port"),
					},
					HTTPGateway: fxgrpc.HTTPGatewayConfig{
						Host: cctx.String("nn-crawler-http-gateway-host"),
						Port: cctx.String("nn-crawler-http-gateway-port"),
					},
					LivenessGateway: fxgrpc.LivenessGatewayConfig{
						Host: cctx.String("nn-crawler-liveness-gateway-host"),
						Port: cctx.String("nn-crawler-liveness-gateway-port"),
					},
				},
			}
		}),
	)
}

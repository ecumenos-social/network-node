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
					Production: cctx.Bool("nn-admin-logger-production"),
				},
				GRPC: &fxgrpc.Config{
					GRPC: fxgrpc.GRPCConfig{
						Host:                                    cctx.String("nn-admin-grpc-host"),
						Port:                                    cctx.String("nn-admin-grpc-port"),
						MaxConnectionAge:                        cctx.Duration("nn-admin-grpc-max-conn-age"),
						KeepAliveEnforcementMinTime:             cctx.Duration("nn-admin-grpc-keep-alive-enforcement-min-time"),
						KeepAliveEnforcementPermitWithoutStream: cctx.Bool("nn-admin-grpc-keep-alive-enforcement-permit-without-stream"),
					},
					Health: fxgrpc.HealthConfig{
						Enabled: cctx.Bool("nn-admin-enabled-health-server"),
						Host:    cctx.String("nn-admin-health-server-host"),
						Port:    cctx.String("nn-admin-health-server-port"),
					},
					HTTPGateway: fxgrpc.HTTPGatewayConfig{
						Host: cctx.String("nn-admin-http-gateway-host"),
						Port: cctx.String("nn-admin-http-gateway-port"),
					},
					LivenessGateway: fxgrpc.LivenessGatewayConfig{
						Host: cctx.String("nn-admin-liveness-gateway-host"),
						Port: cctx.String("nn-admin-liveness-gateway-port"),
					},
				},
			}
		}),
	)
}

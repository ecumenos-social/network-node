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
					Production: cctx.Bool("nn-ingress-service-logger-production"),
				},
				GRPC: &fxgrpc.Config{
					GRPC: fxgrpc.GRPCConfig{
						Host:                                    cctx.String("nn-ingress-service-grpc-host"),
						Port:                                    cctx.String("nn-ingress-service-grpc-port"),
						MaxConnectionAge:                        cctx.Duration("nn-ingress-service-grpc-max-conn-age"),
						KeepAliveEnforcementMinTime:             cctx.Duration("nn-ingress-service-grpc-keep-alive-enforcement-min-time"),
						KeepAliveEnforcementPermitWithoutStream: cctx.Bool("nn-ingress-service-grpc-keep-alive-enforcement-permit-without-stream"),
					},
					Health: fxgrpc.HealthConfig{
						Enabled: cctx.Bool("nn-ingress-service-enabled-health-server"),
						Host:    cctx.String("nn-ingress-service-health-server-host"),
						Port:    cctx.String("nn-ingress-service-health-server-port"),
					},
					HTTPGateway: fxgrpc.HTTPGatewayConfig{
						Host: cctx.String("nn-ingress-service-http-gateway-host"),
						Port: cctx.String("nn-ingress-service-http-gateway-port"),
					},
					LivenessGateway: fxgrpc.LivenessGatewayConfig{
						Host: cctx.String("nn-ingress-service-liveness-gateway-host"),
						Port: cctx.String("nn-ingress-service-liveness-gateway-port"),
					},
				},
			}
		}),
	)
}

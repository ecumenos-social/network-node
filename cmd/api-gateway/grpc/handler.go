package grpc

import (
	"context"

	grpcutils "github.com/ecumenos-social/grpc-utils"
	pbv1 "github.com/ecumenos-social/schemas/proto/gen/networknode/v1"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Handler struct {
	pbv1.ApiGatewayServiceServer

	logger *zap.Logger
}

var _ pbv1.ApiGatewayServiceServer = (*Handler)(nil)

type handlerParams struct {
	fx.In

	Logger *zap.Logger
}

func NewHandler(params handlerParams) *Handler {
	return &Handler{
		logger: params.Logger,
	}
}

func (h *Handler) customizeLogger(ctx context.Context, operationName string) *zap.Logger {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return h.logger
	}

	l := h.logger.With(
		zap.String("operation-name", operationName),
	)
	if corrID := md.Get("correlation-id"); len(corrID) > 0 {
		l = l.With(zap.String("correlation-id", corrID[0]))
	}
	if ip := grpcutils.ExtractRemoteIPAddress(ctx); ip != "" {
		l = l.With(zap.String("remote-ip-address", ip))
	}

	return l
}

func (h *Handler) GetGroups(ctx context.Context, _ *pbv1.ApiGatewayServiceGetGroupsRequest) (*pbv1.ApiGatewayServiceGetGroupsResponse, error) {
	logger := h.customizeLogger(ctx, "GetGroups")
	defer logger.Info("request processed")

	return &pbv1.ApiGatewayServiceGetGroupsResponse{}, status.Error(codes.Unimplemented, "unimplemented")
}

func (h *Handler) GetGroupByID(ctx context.Context, _ *pbv1.ApiGatewayServiceGetGroupByIDRequest) (*pbv1.ApiGatewayServiceGetGroupByIDResponse, error) {
	logger := h.customizeLogger(ctx, "GetGroupByID")
	defer logger.Info("request processed")

	return &pbv1.ApiGatewayServiceGetGroupByIDResponse{}, nil
	//
}

func (h *Handler) GetGroupPublications(ctx context.Context, _ *pbv1.ApiGatewayServiceGetGroupPublicationsRequest) (*pbv1.ApiGatewayServiceGetGroupPublicationsResponse, error) {
	logger := h.customizeLogger(ctx, "GetGroupPublications")
	defer logger.Info("request processed")

	return &pbv1.ApiGatewayServiceGetGroupPublicationsResponse{}, nil
	//
}

func (h *Handler) GetChannels(ctx context.Context, _ *pbv1.ApiGatewayServiceGetChannelsRequest) (*pbv1.ApiGatewayServiceGetChannelsResponse, error) {
	logger := h.customizeLogger(ctx, "GetChannels")
	defer logger.Info("request processed")

	return &pbv1.ApiGatewayServiceGetChannelsResponse{}, nil
	//
}

func (h *Handler) GetChannelByID(ctx context.Context, _ *pbv1.ApiGatewayServiceGetChannelByIDRequest) (*pbv1.ApiGatewayServiceGetChannelByIDResponse, error) {
	logger := h.customizeLogger(ctx, "GetChannelByID")
	defer logger.Info("request processed")

	return &pbv1.ApiGatewayServiceGetChannelByIDResponse{}, nil
	//
}

func (h *Handler) GetChannelPublications(ctx context.Context, _ *pbv1.ApiGatewayServiceGetChannelPublicationsRequest) (*pbv1.ApiGatewayServiceGetChannelPublicationsResponse, error) {
	logger := h.customizeLogger(ctx, "GetChannelPublications")
	defer logger.Info("request processed")

	return &pbv1.ApiGatewayServiceGetChannelPublicationsResponse{}, nil
	//
}

func (h *Handler) GetPublicationByID(ctx context.Context, _ *pbv1.ApiGatewayServiceGetPublicationByIDRequest) (*pbv1.ApiGatewayServiceGetPublicationByIDResponse, error) {
	logger := h.customizeLogger(ctx, "GetPublicationByID")
	defer logger.Info("request processed")

	return &pbv1.ApiGatewayServiceGetPublicationByIDResponse{}, nil
	//
}

func (h *Handler) GetFeed(ctx context.Context, _ *pbv1.ApiGatewayServiceGetFeedRequest) (*pbv1.ApiGatewayServiceGetFeedResponse, error) {
	logger := h.customizeLogger(ctx, "GetFeed")
	defer logger.Info("request processed")

	return &pbv1.ApiGatewayServiceGetFeedResponse{}, nil
	//
}

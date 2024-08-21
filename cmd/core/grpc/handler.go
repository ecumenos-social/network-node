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
	pbv1.CoreServiceServer

	logger *zap.Logger
}

var _ pbv1.CoreServiceServer = (*Handler)(nil)

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

func (h *Handler) GetGroups(ctx context.Context, _ *pbv1.CoreServiceGetGroupsRequest) (*pbv1.CoreServiceGetGroupsResponse, error) {
	logger := h.customizeLogger(ctx, "GetGroups")
	defer logger.Info("request processed")

	return &pbv1.CoreServiceGetGroupsResponse{}, status.Error(codes.Unimplemented, "unimplemented")
}

func (h *Handler) GetGroupByID(ctx context.Context, _ *pbv1.CoreServiceGetGroupByIDRequest) (*pbv1.CoreServiceGetGroupByIDResponse, error) {
	logger := h.customizeLogger(ctx, "GetGroupByID")
	defer logger.Info("request processed")

	return &pbv1.CoreServiceGetGroupByIDResponse{}, status.Error(codes.Unimplemented, "unimplemented")
}

func (h *Handler) GetGroupPublications(ctx context.Context, _ *pbv1.CoreServiceGetGroupPublicationsRequest) (*pbv1.CoreServiceGetGroupPublicationsResponse, error) {
	logger := h.customizeLogger(ctx, "GetGroupPublications")
	defer logger.Info("request processed")

	return &pbv1.CoreServiceGetGroupPublicationsResponse{}, status.Error(codes.Unimplemented, "unimplemented")
}

func (h *Handler) GetChannels(ctx context.Context, _ *pbv1.CoreServiceGetChannelsRequest) (*pbv1.CoreServiceGetChannelsResponse, error) {
	logger := h.customizeLogger(ctx, "GetChannels")
	defer logger.Info("request processed")

	return &pbv1.CoreServiceGetChannelsResponse{}, status.Error(codes.Unimplemented, "unimplemented")
}

func (h *Handler) GetChannelByID(ctx context.Context, _ *pbv1.CoreServiceGetChannelByIDRequest) (*pbv1.CoreServiceGetChannelByIDResponse, error) {
	logger := h.customizeLogger(ctx, "GetChannelByID")
	defer logger.Info("request processed")

	return &pbv1.CoreServiceGetChannelByIDResponse{}, status.Error(codes.Unimplemented, "unimplemented")
}

func (h *Handler) GetChannelPublications(ctx context.Context, _ *pbv1.CoreServiceGetChannelPublicationsRequest) (*pbv1.CoreServiceGetChannelPublicationsResponse, error) {
	logger := h.customizeLogger(ctx, "GetChannelPublications")
	defer logger.Info("request processed")

	return &pbv1.CoreServiceGetChannelPublicationsResponse{}, status.Error(codes.Unimplemented, "unimplemented")
}

func (h *Handler) GetPublicationByID(ctx context.Context, _ *pbv1.CoreServiceGetPublicationByIDRequest) (*pbv1.CoreServiceGetPublicationByIDResponse, error) {
	logger := h.customizeLogger(ctx, "GetPublicationByID")
	defer logger.Info("request processed")

	return &pbv1.CoreServiceGetPublicationByIDResponse{}, status.Error(codes.Unimplemented, "unimplemented")
}

func (h *Handler) GetFeed(ctx context.Context, _ *pbv1.CoreServiceGetFeedRequest) (*pbv1.CoreServiceGetFeedResponse, error) {
	logger := h.customizeLogger(ctx, "GetFeed")
	defer logger.Info("request processed")

	return &pbv1.CoreServiceGetFeedResponse{}, status.Error(codes.Unimplemented, "unimplemented")
}

func (h *Handler) RegisterGroup(ctx context.Context, _ *pbv1.CoreServiceRegisterGroupRequest) (*pbv1.CoreServiceRegisterGroupResponse, error) {
	logger := h.customizeLogger(ctx, "RegisterGroup")
	defer logger.Info("request processed")

	return &pbv1.CoreServiceRegisterGroupResponse{}, status.Error(codes.Unimplemented, "unimplemented")
}

func (h *Handler) ModifyGroup(ctx context.Context, _ *pbv1.CoreServiceModifyGroupRequest) (*pbv1.CoreServiceModifyGroupResponse, error) {
	logger := h.customizeLogger(ctx, "ModifyGroup")
	defer logger.Info("request processed")

	return &pbv1.CoreServiceModifyGroupResponse{}, status.Error(codes.Unimplemented, "unimplemented")
}

func (h *Handler) JoinGroup(ctx context.Context, _ *pbv1.CoreServiceJoinGroupRequest) (*pbv1.CoreServiceJoinGroupResponse, error) {
	logger := h.customizeLogger(ctx, "JoinGroup")
	defer logger.Info("request processed")

	return &pbv1.CoreServiceJoinGroupResponse{}, status.Error(codes.Unimplemented, "unimplemented")
}

func (h *Handler) RequestJoinGroup(ctx context.Context, _ *pbv1.CoreServiceRequestJoinGroupRequest) (*pbv1.CoreServiceRequestJoinGroupResponse, error) {
	logger := h.customizeLogger(ctx, "RequestJoinGroup")
	defer logger.Info("request processed")

	return &pbv1.CoreServiceRequestJoinGroupResponse{}, status.Error(codes.Unimplemented, "unimplemented")
}

func (h *Handler) RegisterChannel(ctx context.Context, _ *pbv1.CoreServiceRegisterChannelRequest) (*pbv1.CoreServiceRegisterChannelResponse, error) {
	logger := h.customizeLogger(ctx, "RegisterChannel")
	defer logger.Info("request processed")

	return &pbv1.CoreServiceRegisterChannelResponse{}, status.Error(codes.Unimplemented, "unimplemented")
}

func (h *Handler) ModifyChannel(ctx context.Context, _ *pbv1.CoreServiceModifyChannelRequest) (*pbv1.CoreServiceModifyChannelResponse, error) {
	logger := h.customizeLogger(ctx, "ModifyChannel")
	defer logger.Info("request processed")

	return &pbv1.CoreServiceModifyChannelResponse{}, status.Error(codes.Unimplemented, "unimplemented")
}

func (h *Handler) FollowChannel(ctx context.Context, _ *pbv1.CoreServiceFollowChannelRequest) (*pbv1.CoreServiceFollowChannelResponse, error) {
	logger := h.customizeLogger(ctx, "FollowChannel")
	defer logger.Info("request processed")

	return &pbv1.CoreServiceFollowChannelResponse{}, status.Error(codes.Unimplemented, "unimplemented")
}

func (h *Handler) RequestFollowChannel(ctx context.Context, _ *pbv1.CoreServiceRequestFollowChannelRequest) (*pbv1.CoreServiceRequestFollowChannelResponse, error) {
	logger := h.customizeLogger(ctx, "RequestFollowChannel")
	defer logger.Info("request processed")

	return &pbv1.CoreServiceRequestFollowChannelResponse{}, status.Error(codes.Unimplemented, "unimplemented")
}

func (h *Handler) InviteToChannel(ctx context.Context, _ *pbv1.CoreServiceInviteToChannelRequest) (*pbv1.CoreServiceInviteToChannelResponse, error) {
	logger := h.customizeLogger(ctx, "InviteToChannel")
	defer logger.Info("request processed")

	return &pbv1.CoreServiceInviteToChannelResponse{}, status.Error(codes.Unimplemented, "unimplemented")
}

func (h *Handler) ReactOnPublication(ctx context.Context, _ *pbv1.CoreServiceReactOnPublicationRequest) (*pbv1.CoreServiceReactOnPublicationResponse, error) {
	logger := h.customizeLogger(ctx, "ReactOnPublication")
	defer logger.Info("request processed")

	return &pbv1.CoreServiceReactOnPublicationResponse{}, status.Error(codes.Unimplemented, "unimplemented")
}

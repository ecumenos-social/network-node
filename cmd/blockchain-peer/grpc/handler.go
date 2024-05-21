package grpc

import (
	pbv1 "github.com/ecumenos-social/schemas/proto/gen/networknode/v1"
)

type Handler struct {
	pbv1.BlockchainPeerServiceServer
}

var _ pbv1.BlockchainPeerServiceServer = (*Handler)(nil)

func NewHandler() *Handler {
	return &Handler{}
}

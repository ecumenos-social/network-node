package grpc

import (
	pbv1 "github.com/ecumenos-social/schemas/proto/gen/networknode/v1"
)

type Handler struct {
	pbv1.IngressServiceServer
}

var _ pbv1.IngressServiceServer = (*Handler)(nil)

func NewHandler() *Handler {
	return &Handler{}
}

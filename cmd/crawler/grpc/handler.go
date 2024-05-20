package grpc

import (
	pbv1 "github.com/ecumenos-social/schemas/proto/gen/networknode/v1"
)

type Handler struct {
	pbv1.CrawlerServiceServer
}

var _ pbv1.CrawlerServiceServer = (*Handler)(nil)

func NewHandler() *Handler {
	return &Handler{}
}

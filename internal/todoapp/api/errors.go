package api

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func toGrpcError(err error) error {
	return status.Error(codes.Internal, "internal error")
}

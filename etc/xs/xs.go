package xs

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

/* -~-~-~-~ Extra Small Types -~-~-~-~ */

// This wouldn't exist if context.Context had been called ctx.Ctx instead.

/* -~-~-~-~ Core -~-~-~-~ */

type Ctx = context.Context

type Num interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

/* -~-~-~-~ gRPC -~-~-~-~ */

// GrpcInfo is a grpc.UnaryServerInfo.
type GrpcInfo = grpc.UnaryServerInfo

// GrpcHandler is a grpc.UnaryHandler.
type GrpcHandler = grpc.UnaryHandler

// GrpcServerOpts is a []grpc.ServerOption.
type GrpcServerOpts = []grpc.ServerOption

// GrpcInterceptors is a []grpc.UnaryServerInterceptor.
type GrpcInterceptors = []grpc.UnaryServerInterceptor

// GrpcDialOpts is a []grpc.DialOption.
type GrpcDialOpts = []grpc.DialOption

// GrpcTLSCreds is a credentials.TransportCredentials.
type GrpcTLSCreds = credentials.TransportCredentials

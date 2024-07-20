package web

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

/* ~-~-~-~ God's gRPC Package -~-~-~-~ */

type (
	GrpcInfo         = grpc.UnaryServerInfo
	GrpcHandler      = grpc.UnaryHandler
	GrpcServerOpts   = []grpc.ServerOption
	GrpcInterceptors = []grpc.UnaryServerInterceptor
	GrpcDialOpts     = []grpc.DialOption
	GrpcTLSCreds     = credentials.TransportCredentials
)

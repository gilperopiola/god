package web

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type (
	// Type aliases
	GrpcInfo         = grpc.UnaryServerInfo
	GrpcHandler      = grpc.UnaryHandler
	GrpcInterceptors = []grpc.UnaryServerInterceptor
	GrpcServerOpts   = []grpc.ServerOption
	GrpcDialOpts     = []grpc.DialOption

	TLSCreds = credentials.TransportCredentials
)

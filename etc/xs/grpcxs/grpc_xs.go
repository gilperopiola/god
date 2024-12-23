package grpcxs

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type (
	ServerInfo       = grpc.UnaryServerInfo
	Handler          = grpc.UnaryHandler
	ServerOpts       = []grpc.ServerOption
	GrpcInterceptors = []grpc.UnaryServerInterceptor
	GrpcDialOpts     = []grpc.DialOption
	GrpcTLSCreds     = credentials.TransportCredentials
)

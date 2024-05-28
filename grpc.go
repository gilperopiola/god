package god

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type (
	GRPCInfo         = grpc.UnaryServerInfo
	GRPCHandler      = grpc.UnaryHandler
	GRPCInterceptors = []grpc.UnaryServerInterceptor
	GRPCServerOpts   = []grpc.ServerOption
	GRPCDialOpts     = []grpc.DialOption
	GRPCSvcRegistrar = grpc.ServiceRegistrar
	TLSCreds         = credentials.TransportCredentials
)

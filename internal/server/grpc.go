package server

import (
	"load-book/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(bc *conf.Bootstrap, pbServer *PbServer, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if bc.Server.Grpc.Network != "" {
		opts = append(opts, grpc.Network(bc.Server.Grpc.Network))
	}
	if bc.Server.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(bc.Server.Grpc.Addr))
	}
	if bc.Server.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(bc.Server.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	pbServer.RegisterGRPC(srv)
	return srv
}

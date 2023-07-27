package server

import (
	v1 "load-book/api/v1"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type PbServer struct {
	Hs v1.HealthServer
	Cs v1.CategoryServer
}

func (c *PbServer) RegisterGRPC(srv *grpc.Server) {
	v1.RegisterHealthServer(srv, c.Hs)
	v1.RegisterCategoryServer(srv, c.Cs)
}

func (c *PbServer) RegisterHTTP(srv *http.Server) {
	v1.RegisterHealthHTTPServer(srv, c.Hs)
	v1.RegisterCategoryHTTPServer(srv, c.Cs)
}

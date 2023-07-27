package server

import (
	"load-book/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func NewHTTPServer(bc *conf.Bootstrap, pbServer *PbServer, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if bc.Server.Http.Network != "" {
		opts = append(opts, http.Network(bc.Server.Http.Network))
	}
	if bc.Server.Http.Addr != "" {
		opts = append(opts, http.Address(bc.Server.Http.Addr))
	}
	// if bc.Server.Http.Timeout != nil {
	// 	opts = append(opts, http.Timeout(bc.Server.Http.Timeout.AsDuration()))
	// }
	srv := http.NewServer(opts...)
	pbServer.RegisterHTTP(srv)
	return srv
}

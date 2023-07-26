//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"load-book/internal/conf"
	"load-book/internal/data"
	"load-book/internal/server"
	"load-book/internal/service/category"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"load-book/internal/service/health"
)

func wireApp(*conf.Bootstrap, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(data.ProviderSet, server.ProviderSet, health.ProviderSet, category.ProviderSet, newApp))
}

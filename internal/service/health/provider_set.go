package health

import (
	"load-book/internal/service/health/biz"
	"load-book/internal/service/health/data"
	"load-book/internal/service/health/service"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(service.NewService, biz.NewUseCase, data.NewRepo)

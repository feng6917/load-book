package category

import (
	"load-book/internal/service/category/biz"
	"load-book/internal/service/category/data"
	"load-book/internal/service/category/service"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(service.NewService, biz.NewUseCase, data.NewRepo)

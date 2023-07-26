package biz

import (
	v1 "load-book/api/load-book/v1"
	"load-book/internal/model/health"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// UseCase is a Hs useCase.
type UseCase struct {
	repo health.Data
	log  *log.Helper
}

// NewUseCase new a Hs useCase.
func NewUseCase(repo health.Data, logger log.Logger) *UseCase {
	return &UseCase{repo: repo, log: log.NewHelper(logger)}
}

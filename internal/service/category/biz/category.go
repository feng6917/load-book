package biz

import (
	"context"
	v1 "load-book/api/load-book/v1"
	"load-book/internal/model/category"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// UseCase is a Hs useCase.
type UseCase struct {
	repo category.Data
	log  *log.Helper
}

// NewUseCase new a Hs useCase.
func NewUseCase(repo category.Data, logger log.Logger) *UseCase {
	return &UseCase{repo: repo, log: log.NewHelper(logger)}
}

func (c *UseCase) Create(ctx context.Context, req category.Form) (int32, error) {
	cf := category.Category{
		Name: req.Name,
		Desc: req.Desc,
	}
	return c.repo.Create(ctx, &cf)
}

func (c *UseCase) FindByID(ctx context.Context, id int32) (*category.Category, error) {
	return c.repo.FindByID(ctx, id)
}

func (c *UseCase) List(ctx context.Context, f *category.Form) ([]*category.Category, error) {
	return c.repo.List(ctx, f)
}

func (c *UseCase) Update(ctx context.Context, f *category.Form) error {
	up := make(map[string]interface{})
	if f.Name != "" {
		up["name"] = f.Name
	}
	if f.Desc != "" {
		up["desc"] = f.Desc
	}
	return c.repo.Update(ctx, up, f.ID)
}

func (c *UseCase) DeleteByID(ctx context.Context, id int32) error {
	return c.repo.DeleteByID(ctx, id)
}

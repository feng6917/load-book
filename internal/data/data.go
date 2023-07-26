package data

import (
	"context"
	"load-book/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData)

// Data .
type Data struct {
	postgres *gorm.DB
}

// NewData .
func NewData(bc *conf.Bootstrap, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
	}
	p := bc.Server.Postgres
	db, err := New(p.User, p.Pwd, p.Addr, p.Db, p.Schema, UpdateBeforeAutoUpdateAt(true))
	if err != nil {
		return nil, cleanup, err
	}

	return &Data{
		postgres: db,
	}, cleanup, nil
}

func (c *Data) PostgresDb(ctx context.Context) *gorm.DB {
	return c.postgres.WithContext(ctx)
}

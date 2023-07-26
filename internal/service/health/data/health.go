package data

import (
	data2 "load-book/internal/data"
	"load-book/internal/model/health"

	"github.com/go-kratos/kratos/v2/log"
)

type Repo struct {
	data *data2.Data
	log  *log.Helper
}

// NewRepo .
func NewRepo(data *data2.Data, logger log.Logger) health.Data {
	return &Repo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

//func (r *Repo) Save(ctx context.Context, g *model.Hs) (*model.Hs, error) {
//	return g, nil
//}
//
//func (r *Repo) Update(ctx context.Context, g *model.Hs) (*model.Hs, error) {
//	return g, nil
//}
//
//func (r *Repo) FindByID(context.Context, int64) (*model.Hs, error) {
//	return nil, nil
//}
//
//func (r *Repo) ListByHello(context.Context, string) ([]*model.Hs, error) {
//	return nil, nil
//}
//
//func (r *Repo) ListAll(context.Context) ([]*model.Hs, error) {
//	return nil, nil
//}

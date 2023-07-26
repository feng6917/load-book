package data

import (
	"context"
	data2 "load-book/internal/data"
	"load-book/internal/model/category"
	"load-book/internal/utils"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type Repo struct {
	data *data2.Data
	log  *log.Helper
}

// NewRepo .
func NewRepo(data *data2.Data, logger log.Logger) category.Data {
	return &Repo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c *Repo) InitDb(ctx context.Context) *gorm.DB {
	return c.data.PostgresDb(ctx)
}

func (c *Repo) InitCategory(ctx context.Context) *gorm.DB {
	return c.InitDb(ctx).Model(category.Category{})
}

func (c *Repo) Create(ctx context.Context, g *category.Category) (int32, error) {
	err := c.InitCategory(ctx).
		Create(g).Error
	if err != nil {
		return 0, err
	}
	return g.ID, nil
}

func (c *Repo) FindByID(ctx context.Context, id int32) (*category.Category, error) {
	var res category.Category
	err := c.InitCategory(ctx).
		Where("id = ?", id).
		First(&res).Error
	return &res, err
}

func (c *Repo) List(ctx context.Context, form *category.Form) ([]*category.Category, error) {
	var cs []*category.Category
	db := c.InitCategory(ctx)
	if form.Name != "" {
		db.Where("name ilike ?", "%"+form.Name+"%")
	}
	err := utils.Page(db, form.Limit, form.Offset).
		Find(&cs).Error
	return cs, err
}

func (r *Repo) Update(ctx context.Context, up map[string]interface{}, id int32) error {
	err := r.InitCategory(ctx).
		Where("id = ?", id).
		Updates(up).Error
	return err
}

func (c *Repo) DeleteByID(ctx context.Context, id int32) error {
	var res category.Category
	err := c.InitCategory(ctx).
		Where("id = ?", id).
		Delete(&res).Error
	return err
}

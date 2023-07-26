package category

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Data interface {
	Create(ctx context.Context, g *Category) (int32, error)
	List(ctx context.Context, form *Form) ([]*Category, error)
	FindByID(ctx context.Context, id int32) (*Category, error)
	Update(ctx context.Context, up map[string]interface{}, id int32) error
	DeleteByID(ctx context.Context, id int32) error
}

type Category struct {
	ID       int32     `json:"id"`
	Name     string    `json:"name"`
	Desc     string    `json:"desc"`
	CreateAt time.Time `json:"createAt" gorm:"column:create_at"`
	UpdateAt time.Time `json:"updateAt" gorm:"column:update_at"`
}

func (c *Category) TableName() string {
	return "categories"
}

func (c *Category) BeforeCreate(*gorm.DB) error {
	t := time.Now()
	c.CreateAt = t
	return nil
}

type Form struct {
	Limit  int
	Offset int
	ID     int32
	Name   string
	Desc   string
}

package data

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type registerCallback interface {
	Register(name string, fn func(*gorm.DB)) error
}

type gormHookFunc func(tx *gorm.DB)

func extractQuery(tx *gorm.DB) string {
	return tx.Dialector.Explain(tx.Statement.SQL.String(), tx.Statement.Vars...)
}

var _ gorm.Plugin = &PTime{}

func NewPTime(uu bool) *PTime {
	return &PTime{
		UpdateBeforeAutoUpdateAt: uu,
	}
}

const (
	updateAtName string = "update_at"
)

type PTime struct {
	UpdateBeforeAutoUpdateAt bool // time.now
}

func (p *PTime) Name() string {
	return "gorm:PTime"
}

func (p *PTime) Initialize(db *gorm.DB) error {
	registerHooks := []struct {
		callback registerCallback
		hook     gormHookFunc
		name     string
	}{
		// before hooks
		{db.Callback().Update().Before("gorm:before_update"), p.beforeUpdate, p.beforeName("update")},
	}

	for _, h := range registerHooks {
		if err := h.callback.Register(h.name, h.hook); err != nil {
			return fmt.Errorf("register %s hook: %w", h.name, err)
		}
	}

	return nil
}

func (p *PTime) beforeUpdate(tx *gorm.DB) {
	if tx.Statement.Parse(tx.Statement.Model) == nil {
		if p.UpdateBeforeAutoUpdateAt {
			p.SetColumn(tx, updateAtName, p.nowTimeAt())
		}
	}
}

func (p *PTime) beforeName(op string) string {
	return "Pg:before:" + op
}

func (p *PTime) SetColumn(tx *gorm.DB, k string, v time.Time) {
	// 查询字段是否存在
	if f := tx.Statement.Schema.LookUpField(k); f != nil {
		// 查询字段值是否为0
		_, isZero := f.ValueOf(tx.Statement.Context, tx.Statement.ReflectValue)
		if isZero {
			tx.Statement.Statement.SetColumn(k, v, true)
		}
	}
}

// func (p *PTime) zeroTimeAt() time.Time {
//	// return "0001-01-01 00:00:00"
//	return time.Time{}
// }

func (p *PTime) nowTimeAt() time.Time {
	// return time.Now().Format("2006-01-02 15:04:05")
	return time.Now()
}

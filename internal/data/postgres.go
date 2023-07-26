package data

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

const NumberTwo = 2

type Postgres struct {
	Dsn                      string        // 指定格式 postgres://user:pwd@addr/db?search_path=public&TimeZone=Asia/Shanghai&sslmode=disable&connect_timeout=5
	PreferSimpleProtocol     bool          // 默认情况下 启用 prepared statement 缓存
	MaxIdleConn              int           // 最大空闲连接数 默认情况下 5
	MaxOpenConn              int           // 最大连接数 默认情况下 50
	MaxLifeTime              time.Duration // 连接最大存活时间 默认情况下 1*time.hour
	MaxIdleTime              time.Duration // 空闲连接最大存活时间 默认情况下 300*time.Second
	Logger                   log.Logger
	Replicas                 []*Postgres
	UpdateBeforeAutoUpdateAt bool // 更新 自动更新updateAt时间
}

// PostgresOptions options
type PostgresOptions func(*Postgres)

func New(user, pwd, addr, db, schema string, opts ...PostgresOptions) (*gorm.DB, error) {
	return NewWithDsn(fmt.Sprintf("postgres://%s:%s@%s/%s?search_path=%s", user, pwd, addr, db, schema), opts...)
}

func NewWithDsn(dsn string, opts ...PostgresOptions) (*gorm.DB, error) {
	srv := &Postgres{
		Dsn:                  dsn,
		PreferSimpleProtocol: false,
		MaxIdleConn:          5,
		MaxOpenConn:          50,
		MaxLifeTime:          time.Hour,
	}
	for _, o := range opts {
		o(srv)
	}
	dial := postgres.New(postgres.Config{DSN: srv.Dsn, PreferSimpleProtocol: srv.PreferSimpleProtocol})

	var err error

	cfg := &gorm.Config{}
	if srv.Logger != nil {
		//cfg.Logger = &aclConn.Logger
	}

	var client *gorm.DB
	client, err = gorm.Open(dial, cfg)
	if err != nil {
		return nil, err
	}

	if srv.UpdateBeforeAutoUpdateAt {
		_ = client.Use(NewPTime(true))
	}

	var instance *sql.DB
	instance, err = client.DB()
	if err != nil {
		return nil, err
	}

	instance.SetMaxIdleConns(srv.MaxIdleConn)
	instance.SetMaxOpenConns(srv.MaxOpenConn)
	instance.SetConnMaxLifetime(srv.MaxLifeTime)
	instance.SetConnMaxIdleTime(srv.MaxIdleTime)
	if len(srv.Replicas) > 0 {
		var replicaDials []gorm.Dialector
		for _, sr := range srv.Replicas {
			tempDial := postgres.New(postgres.Config{DSN: sr.Dsn, PreferSimpleProtocol: sr.PreferSimpleProtocol})
			replicaDials = append(replicaDials, tempDial)
		}

		if len(replicaDials) > 0 {
			if err = client.Use(dbresolver.Register(dbresolver.Config{
				Sources:  []gorm.Dialector{dial},
				Replicas: replicaDials,
				Policy:   nil,
			})); err != nil {
				return nil, err
			}
		}
	}
	return client, nil
}

func Schema(schema string) PostgresOptions {
	return func(c *Postgres) {
		c.setDsnValue("search_path", schema)
	}
}

func SslMode(sslMode string) PostgresOptions {
	return func(c *Postgres) {
		c.setDsnValue("sslmode", sslMode)
	}
}

func Timezone(timezone string) PostgresOptions {
	return func(c *Postgres) {
		c.setDsnValue("TimeZone", timezone)
	}
}

func Timeout(timeout string) PostgresOptions {
	return func(c *Postgres) {
		c.setDsnValue("connect_timeout", timeout)
	}
}

// PreferSimpleProtocol with conn preferSimpleProtocol
func PreferSimpleProtocol(p bool) PostgresOptions {
	return func(c *Postgres) {
		c.PreferSimpleProtocol = p
	}
}

// MaxIdleConn with conn maxIdleConn
func MaxIdleConn(mc int) PostgresOptions {
	return func(c *Postgres) {
		c.MaxIdleConn = mc
	}
}

// MaxOpenConn with conn maxOpenConn
func MaxOpenConn(mc int) PostgresOptions {
	return func(c *Postgres) {
		c.MaxOpenConn = mc
	}
}

// MaxLifetime with conn maxLifetime
func MaxLifetime(ml time.Duration) PostgresOptions {
	return func(c *Postgres) {
		c.MaxLifeTime = ml
	}
}

// Logger with conn logger
func Logger(l log.Logger) PostgresOptions {
	return func(c *Postgres) {
		c.Logger = l
	}
}

// Replicas with conn replicas
func Replicas(cs []*Postgres) PostgresOptions {
	return func(c *Postgres) {
		c.Replicas = cs
	}
}

// UpdateBeforeAutoUpdateAt with conn updateBeforeAutoUpdateAt
func UpdateBeforeAutoUpdateAt(u bool) PostgresOptions {
	return func(c *Postgres) {
		c.UpdateBeforeAutoUpdateAt = u
	}
}

// GetAddr GetAddr
func (c *Postgres) GetAddr() string {
	var addr string
	l := strings.SplitN(c.Dsn, "@", NumberTwo)
	if len(l) == NumberTwo {
		al := strings.SplitN(l[1], "/", NumberTwo)
		if len(al) == NumberTwo {
			addr = al[0]
		}
	}
	return addr
}

func (c *Postgres) setDsnValue(k, v string) {
	if strings.Contains(c.Dsn, "?") {
		dsnList := strings.Split(c.Dsn, "?")
		if len(dsnList) == NumberTwo {
			if strings.Contains(dsnList[1], k) {
				fieldList := strings.Split(dsnList[1], k+"=")
				if len(fieldList) == NumberTwo {
					var setValue string
					if strings.Contains(fieldList[1], "&") {
						vIndex := strings.Index(fieldList[1], "&")
						setValue = fieldList[1][vIndex:]
					}
					c.Dsn = dsnList[0] + "?" + fieldList[0] + k + "=" + v + setValue
				}
			}
		}
	} else {
		c.Dsn = c.Dsn + "?" + k + "=" + v
	}
}

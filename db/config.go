package db

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/jmoiron/sqlx"
	"golang.org/x/sync/singleflight"
	"sync"
	"time"
)

var dbHandler *DatabaseHandler

var initOnce sync.Once

func GetDbInstance() *DatabaseHandler {
	initOnce.Do(func() {
		dbInstance, err := sqlx.Open("mysql", "root:_pc233508@tcp(127.0.0.1:3306)/vhr")
		if err != nil {
			panic(err)
		}
		err = dbInstance.Ping()
		if err != nil {
			panic(err)
		}
		dbInstance.SetConnMaxLifetime(time.Minute * 5)
		dbInstance.SetMaxOpenConns(2)
		dbInstance.SetMaxIdleConns(1)
		dbInstance.SetConnMaxIdleTime(time.Minute)
		dbHandler = &DatabaseHandler{
			DBInstance: dbInstance,
			DbCache: Cache{
				store: nil,
			},
			Rq: singleflight.Group{},
		}
	})
	return dbHandler
}

type DatabaseHandler struct {
	DBInstance *sqlx.DB
	DbCache    Cache
	Rq         singleflight.Group
}

func (d *DatabaseHandler) Select(dest any, query string, args ...interface{}) error {
	err := d.DBInstance.Select(dest, query, args...)
	return err
}

func (d *DatabaseHandler) Exec(sql string, args ...interface{}) error {
	_, err := d.DBInstance.Exec(sql, args)
	return err
}

type Cache struct {
	store      cacheStore
	expireTime time.Duration
}

type cacheStore interface {
	Get(string) (string, error)
	Set(string, string, int) error
	Del(string) error
}

func (c *Cache) Del(ctx context.Context, keys ...string) error {
	for i := 0; i < len(keys); i++ {
		err := c.store.Del(keys[i])
		if err != nil {
			c.DelRetry(ctx, keys[i], err)
		}
	}
	return nil
}

func (c *Cache) DelRetry(ctx context.Context, key string, err error) {
}

var CacheIsNil = errors.New("nil Cache")

func (c *Cache) Get(ctx context.Context, key string, val any) error {
	value, err := c.store.Get(key)
	err = json.Unmarshal([]byte(value), val)
	return err
}

func (c *Cache) Set(ctx context.Context, key string, val any) error {
	marshal, err := json.Marshal(val)
	if err != nil {
		return err
	}
	return c.store.Set(key, string(marshal), int(c.expireTime))
}

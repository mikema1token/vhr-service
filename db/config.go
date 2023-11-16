package db

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"golang.org/x/sync/singleflight"
	"reflect"
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
			cache:      sync.Map{},
		}
	})
	return dbHandler
}

type DatabaseHandler struct {
	DBInstance *sqlx.DB
	cache      sync.Map
}

func (d *DatabaseHandler) SelectWithCache(dest any, query string, args ...interface{}) error {
	key := fmt.Sprintf("%s #@$ %s", query, fmt.Sprint(args))
	value, ok := d.cache.Load(key)
	if ok {
		reflect.Copy(reflect.ValueOf(dest), reflect.ValueOf(value))
		return nil
	} else {
		err := d.DBInstance.Select(dest, query, args...)
		if err != nil {
			return err
		}
		d.cache.Store(key, dest)
		return nil
	}
}

func (d *DatabaseHandler) Exec(sql string, tableName string, args ...interface{}) error {
	result, err := d.DBInstance.Exec(sql, args)
	if err != nil {
		return err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if lastInsertId != 0 || rowsAffected != 0 {

	}
	return nil
}

type cache struct {
	store      sync.Map
	expireTime time.Duration
	barrier    singleflight.Group
}

func (c *cache) Del(ctx context.Context, keys ...string) error {

}

func (c *cache) Get(ctx context.Context, key string, val any) error {

}

func (c *cache) Set(ctx context.Context, key string, val any) error {

}

func (c *cache) GetAndUpdate() {

}

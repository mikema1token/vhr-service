package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"golang.org/x/sync/singleflight"
	"log"
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
			DbCache: Cache{
				store:      GetRedisInstance(),
				expireTime: 0,
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

func (d *DatabaseHandler) SelectWithCache(dest any, key, query string, args ...interface{}) error {
	if key == "" {
		err3 := d.DBInstance.Select(dest, query, args...)
		return err3
	} else {
		err2 := d.DbCache.Get(key, dest)
		if err2 == nil {
			return nil
		}
		if !errors.Is(err2, redis.Nil) {
			return err2
		}
		_, err, _ := d.Rq.Do(key, func() (interface{}, error) {
			err3 := d.DbCache.Get(key, dest)
			if errors.Is(err3, redis.Nil) {
				err3 := d.DBInstance.Select(dest, query, args...)
				if err3 != nil {
					return nil, err3
				} else {
					if 0 == reflect.ValueOf(dest).Elem().Len() {
						return nil, sql.ErrNoRows
					}
					err := d.DbCache.Set(key, dest)
					return nil, err
				}
			} else {
				return nil, err3
			}
		})
		return err
	}
}

func (d *DatabaseHandler) Exec(sql, key string, args ...interface{}) error {
	_, err := d.DBInstance.Exec(sql, args)
	if key != "" {
		return d.DbCache.Del(key)
	} else {
		return err
	}
}

type Cache struct {
	store      *redis.Client
	expireTime time.Duration
}

type cacheStore interface {
	Get(string) (string, error)
	Set(string, string, int) error
	Del(string) error
}

func (c *Cache) Del(keys ...string) error {
	intCmd := c.store.Del(context.Background(), keys...)
	return intCmd.Err()
}

func (c *Cache) DelRetry(key string, err error) {
}

var CacheIsNil = errors.New("nil Cache")

func (c *Cache) Get(key string, val any) error {
	cmd := c.store.Get(context.Background(), key)
	if cmd.Err() != nil {
		return cmd.Err()
	}
	log.Print("cache hit")
	bytes, _ := cmd.Bytes()
	return json.Unmarshal(bytes, val)
}

func (c *Cache) Set(key string, val any) error {
	b, _ := json.Marshal(val)
	return c.store.Set(context.Background(), key, b, 0).Err()
}

func GetRedisInstance() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	ping := client.Ping(context.Background())
	if ping.Err() != nil {
		panic(ping.Err())
	}
	return client
}

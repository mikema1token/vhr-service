package db

import (
	"github.com/jmoiron/sqlx"
	"sync"
	"time"
)

var dbInstance *sqlx.DB

var initOnce sync.Once

func GetDbInstance() *sqlx.DB {
	initOnce.Do(func() {
		open, err := sqlx.Open("mysql", "root:_pc233508@tcp(127.0.0.1:3306)/vhr")
		if err != nil {
			panic(err)
		}
		err = open.Ping()
		if err != nil {
			panic(err)
		}
		open.SetConnMaxLifetime(time.Minute * 5)
		open.SetMaxOpenConns(2)
		open.SetMaxIdleConns(1)
		open.SetConnMaxIdleTime(time.Minute)
		dbInstance = open
	})
	return dbInstance
}

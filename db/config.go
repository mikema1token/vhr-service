package db

import (
	"github.com/jmoiron/sqlx"
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

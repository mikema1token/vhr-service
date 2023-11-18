package db

import (
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRedisInstance(t *testing.T) {
	instance := GetRedisInstance()
	assert.True(t, instance != nil)
}

func TestCache(t *testing.T) {
	c := Cache{
		store:      GetRedisInstance(),
		expireTime: 10,
	}
	err := c.Set("afuuu", 1)
	assert.Nil(t, err)
	v := 2
	err = c.Get("afuuu", &v)
	assert.Nil(t, err)
	assert.Equal(t, 1, v)

	err = c.Del("afuuu")
	assert.Nil(t, err)

	err = c.Get("afuuu", 0)
	assert.True(t, errors.Is(err, redis.Nil))
}

func TestDatabaseHandler(t *testing.T) {
	instance := GetDbInstance()

	selectSql := "select * from place where country = ?"
	type place struct {
		Country string  `db:"country"`
		City    *string `db:"city"`
		Telcode *int    `db:"telcode"`
	}
	var dest []place
	err2 := instance.SelectWithCache(&dest, "country:Hong Kong", selectSql, "Hong Kong")
	assert.Nil(t, err2)
	var dest2 []place
	err2 = instance.SelectWithCache(&dest2, "country:Hong Kong", selectSql, "Hong Kong")
	assert.Nil(t, err2)
	t.Log(dest2)
	update := "update place set telcode = ? where country = ?"
	err := instance.Exec(update, "country:Hong Kong", 21, "Hong Kong")
	assert.Nil(t, err)

}

func TestCache_Get(t *testing.T) {
	c := Cache{
		store:      GetRedisInstance(),
		expireTime: 10,
	}
	strVal := "str value"
	err := c.Set("hello", strVal)
	assert.Nil(t, err)
	strCache := ""
	err = c.Get("hello", &strCache)
	assert.Nil(t, err)
	t.Log(strCache)
	type Object struct {
		Name string
		Age  int
	}
	obj := Object{
		Name: "zk",
		Age:  10,
	}
	err = c.Set("obj", obj)
	assert.Nil(t, err)

	obj2 := Object{}
	err = c.Get("obj", &obj2)
	assert.Nil(t, err)
	t.Log(obj2)
}

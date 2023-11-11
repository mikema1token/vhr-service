package db

import (
	"database/sql"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDbInstance(t *testing.T) {
	instance := GetDbInstance()
	assert.True(t, instance != nil)
}

//func TestSqlxExec(t *testing.T) {
//	cityState := `insert into place(country, telcode) values (?,?)`
//	exec, err := GetDbInstance().Exec(cityState, "Hong Kong", 852)
//	assert.Nil(t, err)
//	t.Log(exec.LastInsertId())
//	t.Log(exec.RowsAffected())
//}

type Place struct {
	Country       string
	City          sql.NullString
	TelephoneCode int `db:"telcode""`
}

func TestSqlxQuery(t *testing.T) {
	//1:错误检查
	//2:必须对返回的结果进行检查，避免连接泄露
	//3:以游标的方式工作，避免大内存的占用
	dbInstance := GetDbInstance()

	rows, err := dbInstance.Query("select country, city, telcode from place")
	assert.Nil(t, err)
	for rows.Next() {
		var country string
		var city sql.NullString
		var telcode int
		err = rows.Scan(&country, &city, &telcode)
		assert.Nil(t, err)
		t.Log(country, city, telcode)
	}
	rows.Close()
	assert.Nil(t, rows.Err())

	bytes, err := json.Marshal(Place{})
	assert.Nil(t, err)
	t.Log(string(bytes))
	queryx, err := dbInstance.Queryx("select * from place")
	assert.Nil(t, err)
	for queryx.Next() {
		var p Place
		//1:默认是用小写的域名或者db标签来匹配列
		//如果是嵌套的结构体，含有重名的字段，是在结构体便利中通过广度优先找到的第一个字段来赋值
		err := queryx.StructScan(&p)
		assert.Nil(t, err)
		t.Log(p)
	}
	queryx.Close()
}

func TestQueryRow(t *testing.T) {
	instance := GetDbInstance()
	row := instance.QueryRow("select telcode from place where telcode=?", 852)
	var telcode int
	err := row.Scan(&telcode)
	assert.Nil(t, err)
	t.Log(telcode)
	var p Place
	row2 := instance.QueryRowx("select * from place where telcode=?", 852)
	err = row2.StructScan(&p)
	assert.Nil(t, err)
	t.Log(p)
}

func TestGetAndSelect(t *testing.T) {
	//select会一次加载整个结果集，可能会占用比较多的内存
	name := ""
	names := []string{}
	instance := GetDbInstance()
	err := instance.Get(&name, "select country from place limit 1")
	assert.Nil(t, err)
	t.Log(name)

	err = instance.Select(&names, "select country from place")
	assert.Nil(t, err)
	t.Log(names)

}

func TestTransaction(t *testing.T) {
	//持有一个连接，并在其上begin，不用并发的使用这个连接，这个连接上的操作应该是顺序的处理，
	instance := GetDbInstance()
	beginx, err := instance.Beginx()
	assert.Nil(t, err)
	beginx.Commit()
}

func TestQueryIn(t *testing.T) {
	pp := []Place{}
	query, args, err := sqlx.In("select * from place where telcode in (?)", []int{27, 65, 852})
	assert.Nil(t, err)
	query = GetDbInstance().Rebind(query)
	err = GetDbInstance().Select(&pp, query, args...)
	assert.Nil(t, err)
	t.Log(pp)
}

func TestNameQuery(t *testing.T) {
	arg := map[string]interface{}{
		"country": "South Africa",
		"telcode": []int{27, 65, 852},
	}
	query, args, err := sqlx.Named("select * from place where country=:country and telcode in (:telcode)", arg)
	assert.Nil(t, err)
	query, args, err = sqlx.In(query, args...)
	assert.Nil(t, err)
	pp := []Place{}

	err = GetDbInstance().Select(&pp, query, args...)
	assert.Nil(t, err)
	t.Log(pp)
}

package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateSqlByStructFieldNames(t *testing.T) {
	type Person struct {
		Name    string  `db:"name"`
		Age     int     `db:"age"`
		Sex     string  `db:"sex"`
		Balance float64 `db:"balance"`
	}
	type Student struct {
		Person
		Score float32 `db:"score"`
	}
	query := "select {{field}} from student"
	dest := []Student{}
	query = GenerateQueryFieldByStructFieldNames(dest, query)
	t.Log(query)
}

func TestSqlHelper_DoUpdate(t *testing.T) {
	sql := "update place {{update_field}}"
	sqlHelper := NewSqlHelper(sql)
	sqlHelper.AddUpdateField("country", "中国")
	err := sqlHelper.DoUpdate()
	assert.Nil(t, err)
}

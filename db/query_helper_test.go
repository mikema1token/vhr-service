package db

import "testing"

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

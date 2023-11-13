package db

import (
	"reflect"
	"strings"
)

func GenerateQueryFieldByStructFieldNames(dest any, querySql string) string {
	elem := reflect.TypeOf(dest).Elem()
	dbTagList := make([]string, 0)
	for i := 0; i < elem.NumField(); i++ {
		structField := elem.Field(i)
		dbTagList = append(dbTagList, getDbTagFromField(structField)...)
	}
	return strings.ReplaceAll(querySql, "{{field}}", strings.Join(dbTagList, ","))
}

func getDbTagFromField(field reflect.StructField) []string {
	if field.Anonymous {
		dbTagList := make([]string, 0)
		fieldType := field.Type
		for i := 0; i < fieldType.NumField(); i++ {
			dbTagList = append(dbTagList, getDbTagFromField(fieldType.Field(i))...)
		}
		return dbTagList
	} else {
		return []string{field.Tag.Get("db")}
	}
}

func GenerateUpdateField() {

}

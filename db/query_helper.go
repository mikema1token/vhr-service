package db

import (
	"fmt"
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

type SqlHelper struct {
	sql             string
	updateParamList []UpdateParam
}

type UpdateParam struct {
	FieldName string
	Value     any
}

func (h *SqlHelper) AddUpdateField(fieldName string, value any) {
	h.updateParamList = append(h.updateParamList, UpdateParam{
		FieldName: fieldName,
		Value:     value,
	})
}

func (h *SqlHelper) getUpdateFieldNameAndValue() ([]string, []any) {
	updateField := make([]string, 0)
	value := make([]any, 0)
	for _, param := range h.updateParamList {
		updateField = append(updateField, fmt.Sprintf("%s = ?", param.FieldName))
		value = append(value, param.Value)
	}
	return updateField, value
}

func (h *SqlHelper) DoUpdate() error {
	fieldNames, values := h.getUpdateFieldNameAndValue()
	updateSql := strings.ReplaceAll(h.sql, "{{update_field}}", "set "+strings.Join(fieldNames, ","))
	_, err := GetDbInstance().Exec(updateSql, values...)
	return err
}

func NewSqlHelper(sql string) SqlHelper {
	return SqlHelper{
		sql: sql,
	}
}

package db

import (
	"fmt"
	"reflect"
	"strings"
)

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
	databaseHandler *DatabaseHandler
	sql             string
	updateParamList []Param
	whereParam      []Param
}

type Param struct {
	fieldName string
	cmp       string
	value     any
}

func (h *SqlHelper) replaceSelectField(dest any) {
	ptrElem := reflect.TypeOf(dest).Elem()
	sliceElem := ptrElem.Elem()
	dbTagList := make([]string, 0)
	for i := 0; i < sliceElem.NumField(); i++ {
		structField := sliceElem.Field(i)
		dbTagList = append(dbTagList, getDbTagFromField(structField)...)
	}
	h.sql = strings.ReplaceAll(h.sql, "{{select_field}}", strings.Join(dbTagList, ","))
}

func NewSqlHelper(sql string) *SqlHelper {
	return &SqlHelper{
		databaseHandler: GetDbInstance(),
		sql:             sql,
	}
}

//func (h *SqlHelper) DoQuery(dest any) error {
//	h.replaceSelectField(dest)
//	args := h.replaceWhereField()
//	if len(h.tableCache) != 0 {
//		return h.databaseHandler.SelectWithCache(dest, h.sql, args...)
//	} else {
//		return h.databaseHandler.DBInstance.Select(dest, h.sql, args...)
//	}
//}

func (h *SqlHelper) replaceWhereField() []any {
	conditions := make([]string, 0)
	args := make([]any, 0)
	for _, param := range h.whereParam {
		conditions = append(conditions, fmt.Sprintf("%s %s ?", param.fieldName, param.cmp))
		args = append(args, param.value)
	}
	h.sql = strings.ReplaceAll(h.sql, "{{where_field}}", strings.Join(conditions, " and "))
	return args
}

func (h *SqlHelper) Update() error {
	args := h.replaceUpdateField()
	t := h.replaceWhereField()
	args = append(args, t...)
	_, err := h.databaseHandler.DBInstance.Exec(h.sql, args...)
	return err
}

func (h *SqlHelper) AddUpdateField(fieldName, cmp string, value any) {
	h.updateParamList = append(h.updateParamList, Param{
		fieldName: fieldName,
		cmp:       cmp,
		value:     value,
	})
}

func (h *SqlHelper) AddWhereParam(key, cmp string, value any) {
	h.whereParam = append(h.whereParam, Param{
		fieldName: key,
		cmp:       cmp,
		value:     value,
	})
}

func (h *SqlHelper) replaceUpdateField() []any {
	conditions := make([]string, 0)
	args := make([]any, 0)
	for _, param := range h.updateParamList {
		conditions = append(conditions, fmt.Sprintf("%s = ?", param.fieldName))
		args = append(args, param.value)
	}
	h.sql = strings.ReplaceAll(h.sql, "{{update_field}}", "set "+strings.Join(conditions, ","))
	return args
}

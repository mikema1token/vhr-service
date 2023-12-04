package db

import _ "github.com/go-sql-driver/mysql"

type UserModel struct {
	Id        int     `json:"id" db:"id"`
	Name      string  `json:"name" db:"name"`
	Phone     string  `json:"phone" db:"phone"`
	Telephone string  `json:"telephone" db:"telephone"`
	Address   string  `json:"address" db:"address"`
	Enabled   bool    `json:"enabled" db:"enabled"`
	Username  string  `json:"username" db:"username"`
	Password  string  `json:"password" db:"password"`
	UserFace  string  `json:"userface" db:"userface"`
	Remark    *string `json:"remark" db:"remark"`
}

func GetUserModelByName(name string) (UserModel, error) {
	query := `select {{select_field}} from hr where {{where_field}}`
	sqlHelper := NewSqlHelper(query)
	sqlHelper.AddWhereParam("username", "=", name)
	var r []UserModel
	err := sqlHelper.DoQuery(&r, "hr.username:"+name)
	if err != nil {
		return UserModel{}, err
	} else {
		return r[0], err
	}
}

func GetUserList() ([]UserModel, error) {
	dbInstance := GetDbInstance()
	var userList []UserModel
	querySql := "select * from hr"
	err := dbInstance.DBInstance.Select(&userList, querySql)
	return userList, err
}

type Role struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	NameZh string `db:"nameZh" json:"nameZh"`
}

func GetRoles() ([]Role, error) {
	var roles []Role
	err := GetDbInstance().DBInstance.Select(&roles, "select * from role")
	return roles, err
}

func GetRoleMenus(id int) ([]Menu, error) {
	query := `WITH RECURSIVE MenuCTE AS (
				  SELECT id, name, parentId
				  FROM menu
				  WHERE id IN (SELECT mid FROM menu_role WHERE rid = ?)
				  UNION
				  SELECT m.id, m.name, m.parentId
				  FROM menu m
				  JOIN MenuCTE cte ON m.parentId = cte.id
				)
				SELECT * FROM MenuCTE`

	var menus []Menu
	err := GetDbInstance().DBInstance.Select(&menus, query, id)
	return menus, err
}

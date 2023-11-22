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
	sqlHelper.AddWhereParam("name", "=", name)
	var r []UserModel
	err := sqlHelper.DoQuery(&r, "hr.name:"+name)
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

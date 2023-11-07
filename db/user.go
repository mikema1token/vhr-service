package db

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type UserModel struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Phone     string  `json:"phone"`
	Telephone string  `json:"telephone"`
	Address   string  `json:"address"`
	Enabled   bool    `json:"enabled"`
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	UserFace  string  `json:"userface"`
	Remark    *string `json:"remark"`
}

func GetUserModelByName(name string) (UserModel, error) {
	query := `select * from hr where username = ?`
	db, err := sql.Open("mysql", "root:_pc233508@tcp(127.0.0.1:3306)/vhr")
	if err != nil {
		return UserModel{}, err
	}
	rows, err := db.Query(query, name)
	if err != nil {
		return UserModel{}, err
	}
	for rows.Next() {
		user := UserModel{}
		err = rows.Scan(&user.Id, &user.Name, &user.Phone, &user.Telephone, &user.Address, &user.Enabled, &user.Username, &user.Password, &user.UserFace, &user.Remark)
		if err != nil {
			return UserModel{}, err
		}
		return user, nil
	}
	return UserModel{}, nil
}

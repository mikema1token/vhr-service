package db

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type Position struct {
	Id         int       `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	CreateDate time.Time `json:"createDate" db:"createDate"`
	Enabled    bool      `json:"enabled" db:"enabled"`
}

func ListPosition() ([]Position, error) {
	dbInstance := GetDbInstance()
	var posList []Position
	querySql := "Select * from position"
	err := dbInstance.DBInstance.Select(&posList, querySql)
	return posList, err
}

func AddPosition(name string) error {
	dbInstance := GetDbInstance()
	_, err := dbInstance.DBInstance.Exec("insert into position (name,createDate,enabled) values (?,?,?)", name, time.Now(), 1)
	return err
}

func DelPosition(id int) error {
	instance := GetDbInstance()
	_, err := instance.DBInstance.Exec("delete from position where id = ?", id)
	return err
}

func UpdateMenu(id int, name string) error {
	_, err := GetDbInstance().DBInstance.Exec("update position set name = ? where id = ?", name, id)
	return err
}

func UpdateMenu2(ids []int) error {
	in, i, err := sqlx.In("update position set enabled = 0 where id in(?)", ids)
	if err != nil {
		return err
	}
	in = GetDbInstance().DBInstance.Rebind(in)
	_, err = GetDbInstance().DBInstance.Exec(in, i...)
	return err
}

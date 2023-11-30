package db

import "time"

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
	_, err := dbInstance.DBInstance.Exec("insert into position(name) values(?)", name)
	return err
}

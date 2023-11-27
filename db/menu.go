package db

type Menu struct {
	ID          int     `json:"id" db:"id"`
	URL         *string `json:"url" db:"url"`
	Path        *string `json:"path" db:"path"`
	Component   *string `json:"component" db:"component"`
	Name        *string `json:"name" db:"name"`
	IconCls     *string `json:"iconCls" db:"iconCls"`
	KeepAlive   *int    `json:"keepAlive" db:"keepAlive"`
	RequireAuth *int    `json:"requireAuth" db:"requireAuth"`
	ParentID    *int    `json:"parentId" db:"parentId"`
	Enabled     *int    `json:"enabled" db:"enabled"`
	Children    []Menu  `json:"children"`
}

func ListMenu() ([]Menu, error) {
	var dest []Menu
	err := GetDbInstance().DBInstance.Select(&dest, "select * from menu")
	return dest, err
}

package service

import (
	"github.com/gin-gonic/gin"
	"vhr-service/db"
)

func GetMenuTree(c *gin.Context) {
	menu, err := db.ListMenu()
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	tree := BuildMenuTree(menu, 1)
	c.JSON(200, gin.H{"data": tree})
}

func BuildMenuTree(menuList []db.Menu, parentId int) []db.Menu {
	var tree []db.Menu
	for _, menu := range menuList {
		if menu.ParentID != nil && *menu.ParentID == parentId {
			menu.Children = BuildMenuTree(menuList, menu.ID)
			tree = append(tree, menu)
		}
	}
	return tree
}

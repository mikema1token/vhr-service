package service

import (
	"github.com/gin-gonic/gin"
	"strconv"
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

func UpdateMenu(c *gin.Context) {
	var r struct {
		Id   int
		Name string
	}
	err := c.ShouldBind(&r)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	err = db.UpdateMenu(r.Id, r.Name)
	if err != nil {
		c.String(500, err.Error())
	} else {
		c.JSON(200, gin.H{"code": "ok"})
	}
}

func DeleteMenu(c *gin.Context) {
	var ids []int
	err := c.ShouldBind(&ids)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	err = db.UpdateMenu2(ids)
	if err != nil {
		c.String(500, err.Error())
	} else {
		c.JSON(200, gin.H{"code": "ok"})
	}
}

func GetMenuTreeByRole(c *gin.Context) {
	id := c.Query("id")
	idInt, _ := strconv.Atoi(id)
	menus, err := db.GetRoleMenus(idInt)
	if err != nil {
		c.JSON(200, gin.H{"code": "fail", "msg": err.Error()})
	} else {
		tree := BuildMenuTree(menus, 1)
		c.JSON(200, gin.H{"code": "ok", "data": tree})
	}
}

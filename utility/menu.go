package utility

import (
	v1 "gf-admin/api/device_menu/v1"
	"gf-admin/internal/model/entity"
)

// BuildMenuTree 构建菜单树
func BuildMenuTree(menuList []*entity.DeviceMenu, pid int) []v1.Menu {
	var tree []v1.Menu
	for _, menu := range menuList {
		if menu.Pid == pid {
			node := v1.Menu{
				Id:     menu.Id,
				Pid:    menu.Pid,
				Level:  menu.Level,
				Code:   menu.Code,
				Name:   menu.Name,
				Path:   menu.Path,
				Icon:   menu.Icon,
				Order:  menu.Order,
				Type:   menu.Type,
				Select: menu.Select == 1,
			}
			// 递归构建子菜单
			children := BuildMenuTree(menuList, menu.Id)
			if len(children) > 0 {
				node.Children = children
			}
			tree = append(tree, node)
		}
	}
	return tree
}

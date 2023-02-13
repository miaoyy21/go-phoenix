package xsys

import (
	"database/sql"
	"encoding/json"
	"go-phoenix/asql"
	"go-phoenix/handle"
)

type SysRoleMenus struct {
}

func (o *SysRoleMenus) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	roleId := ctx.FormValue("role_id")
	res, err := asql.Select(tx, "SELECT menu_id_ FROM sys_role_menu WHERE role_id_ = ?", roleId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (o *SysRoleMenus) PostPatch(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	roleId := ctx.PostFormValue("role_id")

	var menus []string
	if err := json.Unmarshal([]byte(ctx.PostFormValue("menus")), &menus); err != nil {
		return nil, err
	}

	// 删除原角色关联的菜单
	if err := asql.Delete(tx, "DELETE FROM sys_role_menu WHERE role_id_ = ?", roleId); err != nil {
		return nil, err
	}

	// 更新新角色关联的菜单
	now := asql.GetNow()
	query := "INSERT INTO sys_role_menu(id, role_id_, menu_id_, create_at_) VALUES (?,?,?,?)"
	for _, menuId := range menus {
		if err := asql.Insert(tx, query, asql.GenerateId(), roleId, menuId, now); err != nil {
			return nil, err
		}
	}

	return map[string]string{"status": "success"}, nil
}

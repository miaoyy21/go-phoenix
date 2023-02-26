package xsys

import (
	"database/sql"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/base"
	"strings"
)

func menusByOrg(tx *sql.Tx, org ...interface{}) ([]map[string]string, error) {
	if len(org) <= 0 {
		return make([]map[string]string, 0), nil
	}

	// 查询用户和所有上级部门的权限
	query := fmt.Sprintf("SELECT role_id_ FROM sys_organization_role WHERE organization_id_ IN (?%s)", strings.Repeat(", ?", len(org)-1))
	rolesRes, err := asql.Select(tx, query, org...)
	if err != nil {
		return nil, err
	}

	// 没有对应角色
	if len(rolesRes) <= 0 {
		return make([]map[string]string, 0), nil
	}

	// 根据角色查询菜单
	sRoles := base.NewStringSet(base.ResAsSliceString(rolesRes, "role_id_"))
	roles := make([]interface{}, 0, len(sRoles.Values()))
	for _, role := range sRoles.Values() {
		roles = append(roles, role)
	}

	query = fmt.Sprintf("SELECT menu_id_ FROM sys_role_menu,sys_menu WHERE sys_role_menu.menu_id_ = sys_menu.id AND role_id_ IN (?%s)", strings.Repeat(", ?", len(roles)-1))
	menusRes, err := asql.Select(tx, query, roles...)
	if err != nil {
		return nil, err
	}

	// 查询所有菜单，构建菜单树
	menusFull, err := asql.Select(tx, "SELECT id,menu_,name_,parent_id_,icon_,order_ FROM sys_menu")
	if err != nil {
		return nil, err
	}

	// 构建
	rte := base.NewRelationTree(menusFull, base.ResAsSliceString(menusRes, "menu_id_"))
	return rte.Build(), nil
}

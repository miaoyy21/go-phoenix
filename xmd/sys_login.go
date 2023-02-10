package xmd

import (
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/base"
	"go-phoenix/handle"
	"strconv"
	"strings"
)

type SysLogin struct {
}

func (m *SysLogin) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	//scope := ctx.FormValue("scope")

	query, args := "invalid", make([]interface{}, 0)
	//if strings.EqualFold(scope, "ALL") {
	//	query = `
	//	SELECT sys_dict_kind.code_ AS code, sys_dict_item.code_ AS id, sys_dict_item.name_ AS value
	//	FROM sys_dict_kind,sys_dict_item
	//	WHERE sys_dict_kind.id = sys_dict_item.kind_id_
	//`
	//}

	res, err := asql.Select(tx, query, args...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *SysLogin) GetByToken(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {

	// 获取用户的部门ID
	departIds, err := asql.QueryRelationParents(tx, "sys_depart", ctx.GetDepartId())
	if err != nil {
		return nil, err
	}

	// 查询用户和所有上级部门的权限
	query := fmt.Sprintf("SELECT role_id_ FROM sys_organization_role WHERE organization_id_ IN (?%s)", strings.Repeat(", ?", len(departIds)))
	org := append(departIds, ctx.GetUserId())
	rolesRes, err := asql.Select(tx, query, org...)
	if err != nil {
		return nil, err
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

func (m *SysLogin) PostByPassword(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	accountId := ctx.PostFormValue("account_id")
	password := ctx.PostFormValue("password")
	departId := ctx.PostFormValue("depart_id")

	var uPwd string
	var userId, userCode, userName, departCode, departName string
	query := `
		SELECT sys_user.id, sys_user.user_code_, sys_user.user_name_, sys_depart.code_, sys_depart.name_, sys_user.password_ 
		FROM sys_user, sys_depart
		WHERE sys_user.depart_id_ = sys_depart.id
			AND sys_user.account_id_ = ? AND sys_user.depart_id_ = ?
	`
	args := []interface{}{accountId, departId}
	if err := asql.SelectRow(tx, query, args...).Scan(&userId, &userCode, &userName, &departCode, &departName, &uPwd); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("无效的登录用户")
		}

		return nil, err
	}

	// 密码比较
	ePwd := base.Config.AesStream([]byte(password))
	if !strings.EqualFold(uPwd, hex.EncodeToString(ePwd)) {
		return nil, errors.New("登录密码不正确")
	}

	// 获取设定的有效时限
	var sExpire string

	query = "SELECT value_ FROM sys_setting WHERE field_ = ?"
	if err := asql.SelectRow(tx, query, "token_expire").Scan(&sExpire); err != nil {
		return nil, err
	}

	iExpire, err := strconv.ParseInt(sExpire, 10, 64)
	if err != nil {
		return nil, err
	}

	res := base.GenerateToken(userId, userCode, userName, departId, departCode, departName, uPwd, ctx.UserAgent(), iExpire)
	return res, nil
}

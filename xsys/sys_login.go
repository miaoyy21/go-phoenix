package xsys

import (
	"database/sql"
	"errors"
	"go-phoenix/asql"
	"go-phoenix/base"
	"go-phoenix/handle"
	"strconv"
	"strings"
)

type SysLogin struct {
}

func (m *SysLogin) GetByToken(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {

	// 获取用户的部门ID
	org, err := asql.QueryRelationParents(tx, "sys_depart", ctx.GetDepartId())
	if err != nil {
		return nil, err
	}

	org = append(org, ctx.GetUserId())
	menus, err := menusByOrg(tx, org...)
	if err != nil {
		return nil, err
	}

	var tasks int

	query := "SELECT COUNT(1) AS count_ FROM wf_flow_task WHERE executor_user_id_ = ? AND status_ = ?"
	if err := asql.SelectRow(tx, query, ctx.GetUserId(), "Executing").Scan(&tasks); err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}

		tasks = 0
	}

	return map[string]interface{}{"menus": menus, "tasks": tasks}, nil
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
	if !strings.EqualFold(uPwd, base.Config.AesEncodeString(password)) {
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

	// 记录用户登录时间
	query = "UPDATE sys_user SET login_at_ = ? WHERE id = ?"
	if err := asql.Update(tx, query, asql.GetNow(), userId); err != nil {
		return nil, err
	}

	res := base.GenerateToken(userId, userCode, userName, departId, departCode, departName, uPwd, ctx.UserAgent(), iExpire)
	return res, nil
}

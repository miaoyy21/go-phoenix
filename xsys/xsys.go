package xsys

import (
	"database/sql"
	"encoding/base64"
	"errors"
	"go-phoenix/asql"
	"go-phoenix/base"
	"go-phoenix/handle"
	"strconv"
	"strings"
	"time"
)

type Sys struct {
}

func (o *Sys) GetSync(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	var tasks int

	// 查询待办事项条数
	if err := tx.QueryRow("SELECT COUNT(1) AS count_ FROM wf_flow_task WHERE executor_user_id_ = ? AND status_ = ?", ctx.UserId(), "Executing").Scan(&tasks); err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}

		tasks = 0
	}

	// 同时记录用户的活跃时间，用于判断用户是否在线
	if _, err := tx.Exec("UPDATE sys_user SET online_at_ = ? WHERE id = ?", asql.GetNow(), ctx.UserId()); err != nil {
		return nil, err
	}

	return map[string]interface{}{"tasks": tasks}, nil
}

func (o *Sys) GetDictionary(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	query := `
		SELECT sys_dict_item.code_ AS id, sys_dict_kind.code_ AS code, sys_dict_item.name_ AS value
		FROM sys_dict_kind, sys_dict_item
		WHERE sys_dict_kind.id = sys_dict_item.kind_id_
		ORDER BY sys_dict_kind.order_ ASC, sys_dict_item.order_ ASC
	`

	return asql.Select(tx, query)
}

func (o *Sys) GetSetting(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	res, err := asql.Select(tx, "SELECT field_, value_ FROM sys_setting")
	if err != nil {
		return nil, err
	}

	return base.ResAsMap(res, true, "field_", "value_"), nil
}

func (o *Sys) GetDepart(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	accountId := ctx.FormValue("account_id")

	query := `	
		SELECT sys_depart.id, sys_depart.name_
		FROM sys_user,sys_depart
		WHERE sys_user.depart_id_ = sys_depart.id
			AND sys_user.account_id_ = ?
		ORDER BY sys_depart.order_ ASC
	`

	return asql.Select(tx, query, accountId)
}

func (o *Sys) PostLoginByPassword(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
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

	// Token
	token := base.GenerateToken(userId, userCode, userName, departId, departCode, departName, uPwd, ctx.UserAgent(), int(iExpire))

	expire := time.Now().Add(time.Duration(iExpire) * time.Second)
	setCookie(ctx, "PHOENIX_LOGIN_TOKEN", token, expire)
	setCookie(ctx, "PHOENIX_USER_ID", userId, expire)
	setCookie(ctx, "PHOENIX_USER_CODE", userCode, expire)
	setCookie(ctx, "PHOENIX_USER_NAME", base64.StdEncoding.EncodeToString([]byte(userName)), expire)
	setCookie(ctx, "PHOENIX_DEPART_ID", departId, expire)
	setCookie(ctx, "PHOENIX_DEPART_CODE", departCode, expire)
	setCookie(ctx, "PHOENIX_DEPART_NAME", base64.StdEncoding.EncodeToString([]byte(departName)), expire)

	return map[string]string{"status": "success"}, nil
}

package xmd

import (
	"crypto/md5"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/base"
	"go-phoenix/handle"
	"strconv"
	"strings"
	"time"
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
	// 获取用户的菜单
	// id,menu_,name_,parent_id_,icon_

	return map[string]string{"status": "success"}, nil
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

	src := make([]string, 0, 3)

	// 0 用户ID
	src = append(src, userId)

	// 1 附加信息
	ext := fmt.Sprintf("%s_%s_%s_%s_%s_%s_%s", ctx.UserAgent(), uPwd, userCode, userName, departId, departCode, departName)
	md5Ext := md5.Sum([]byte(ext))
	src = append(src, base64.StdEncoding.EncodeToString(md5Ext[:]))

	// 获取设定的有效时限
	var e0 string

	query = "SELECT value_ FROM sys_setting WHERE field_ = ?"
	if err := asql.SelectRow(tx, query, "token_expire").Scan(&e0); err != nil {
		return nil, err
	}

	e1, err := strconv.ParseInt(e0, 10, 64)
	if err != nil {
		return nil, err
	}

	// 2 失效时间
	expire := strconv.FormatInt(time.Now().Add(time.Duration(e1)*time.Second).Unix(), 10)
	src = append(src, expire)

	bytes := base.Config.AesStream([]byte(strings.Join(src, ",")))
	res := map[string]string{
		"status":      "success",
		"token":       base64.StdEncoding.EncodeToString(bytes),
		"user_id":     userId,
		"user_code":   userCode,
		"user_name":   userName,
		"depart_id":   departId,
		"depart_code": departCode,
		"depart_name": departName,
		"expire":      expire,
	}

	return res, nil
}

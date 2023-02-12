package xsys

import (
	"database/sql"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/base"
	"go-phoenix/handle"
	"strings"
)

type SysUsers struct {
}

func (o *SysUsers) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	accountId := ctx.FormValue("account_id")
	departId := ctx.FormValue("depart_id")
	scope := ctx.FormValue("scope")

	query, args := "invalid", make([]interface{}, 0)
	if strings.EqualFold(scope, "LOGIN") && len(accountId) > 0 {
		query = `	
			SELECT sys_depart.id, sys_depart.name_
			FROM sys_user,sys_depart
			WHERE sys_user.depart_id_ = sys_depart.id
				AND sys_user.account_id_ = ?
			ORDER BY sys_depart.order_ ASC
		`
		args = append(args, accountId)
	} else if len(departId) > 0 && strings.EqualFold(scope, "ALL") {
		query = `
			SELECT sys_user.id, sys_user.user_code_, sys_user.user_name_, sys_user.account_id_, 
				sys_user.depart_id_, sys_depart.name_ AS depart_name_, 
				sys_user.sex_, sys_user.is_depart_leader_, sys_user.valid_, 
				sys_user.classification_, sys_user.telephone_, sys_user.email_, 
				sys_user.birth_, sys_user.description_, sys_user.create_at_, sys_user.login_at_
			FROM sys_user 
				LEFT JOIN sys_depart ON sys_user.depart_id_ = sys_depart.id
			WHERE sys_user.depart_id_ = ?
			ORDER BY sys_user.order_ ASC
		`
		args = append(args, departId)
	}

	res, err := asql.Select(tx, query, args...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (o *SysUsers) Post(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	operation := ctx.PostFormValue("operation")

	id := ctx.PostFormValue("id")
	userCode := ctx.PostFormValue("user_code_")
	userName := ctx.PostFormValue("user_name_")
	accountId := ctx.PostFormValue("account_id_")
	departId := ctx.PostFormValue("depart_id_")
	sex := ctx.PostFormValue("sex_")
	isDepartLeader := ctx.PostFormValue("is_depart_leader_")
	valid := ctx.PostFormValue("valid_")
	classification := ctx.PostFormValue("classification_")
	telephone := ctx.PostFormValue("telephone_")
	email := ctx.PostFormValue("email_")
	birth := ctx.PostFormValue("birth_")
	description := ctx.PostFormValue("description_")

	moveId := ctx.PostFormValue("webix_move_id")
	moveIndex := ctx.PostFormValue("webix_move_index")
	moveParent := ctx.PostFormValue("webix_move_parent")

	now := asql.GetNow()
	switch operation {
	case "insert":
		newId := asql.GenerateId()
		query := `
		INSERT INTO sys_user(id, user_code_, user_name_, account_id_,
			depart_id_, sex_, is_depart_leader_, valid_, 
			classification_, telephone_, email_, 
			birth_, description_, order_, create_at_
		) VALUES (?,?,?,?, ?,?,?,?, ?,?,?, ?,?,?,?)
		`
		args := []interface{}{
			newId, userCode, userName, accountId,
			departId, sex, isDepartLeader, valid,
			classification, telephone, email,
			birth, description, asql.GenerateOrderId(), now,
		}
		if err := asql.Insert(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "newid": newId, "create_at_": now}, nil
	case "update":
		query := `
		UPDATE sys_user 
		SET user_code_ = ?, user_name_ = ?, account_id_ = ?,  
			depart_id_ = ?, sex_ = ?, is_depart_leader_ = ?, valid_ = ?, 
			classification_ = ?, telephone_ = ?, email_ = ?, 
			birth_ = ?, description_ = ?, update_at_ = ? 
		WHERE id = ?
		`
		args := []interface{}{
			userCode, userName, accountId,
			departId, sex, isDepartLeader, valid,
			classification, telephone, email,
			birth, description, now,
			id,
		}
		if err := asql.Update(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "id": id, "update_at_": now}, nil
	case "delete":
		if err := asql.Delete(tx, "DELETE FROM sys_user WHERE id = ?", id); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "id": id}, nil
	case "order":
		if err := asql.Order(tx, "sys_user", id, moveId, moveIndex, moveParent); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success"}, nil
	}

	return nil, fmt.Errorf("unrecognizable operation %s ", operation)
}

func (o *SysUsers) PostResetPassword(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	id := ctx.PostFormValue("id")

	// 获取默认密码
	var dPwd string
	if err := asql.SelectRow(tx, "SELECT value_ FROM sys_setting WHERE field_ = ?", "password_default").Scan(&dPwd); err != nil {
		return nil, err
	}

	// 密码加密
	ePwd := base.Config.AesEncodeString(dPwd)
	if err := asql.Update(tx, "UPDATE sys_user SET password_ = ? WHERE id = ?", ePwd, id); err != nil {
		return nil, err
	}

	return map[string]interface{}{"status": "success", "id": id}, nil
}

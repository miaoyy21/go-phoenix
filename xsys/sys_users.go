package xsys

import (
	"database/sql"
	"errors"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/base"
	"go-phoenix/handle"
	"strings"
	"time"
)

type SysUsers struct {
}

func (o *SysUsers) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	departId := ctx.FormValue("depart_id")
	scope := ctx.FormValue("scope")

	query, args := "invalid", make([]interface{}, 0)
	if len(departId) > 0 && strings.EqualFold(scope, "SIMPLE") {
		query = `
			SELECT id, user_code_, user_name_, sex_, is_depart_leader_ 
			FROM sys_user
			WHERE depart_id_ = ?
			ORDER BY order_ ASC
		`
		args = append(args, departId)
	} else if len(departId) > 0 && strings.EqualFold(scope, "ALL") {
		query = `
			SELECT sys_user.id, sys_user.user_code_, sys_user.user_name_, sys_user.account_id_, 
				sys_user.depart_id_, sys_depart.name_ AS depart_name_, 
				sys_user.sex_, sys_user.is_depart_leader_, sys_user.valid_, 
				sys_user.classification_, sys_user.telephone_, sys_user.email_, 
				sys_user.birth_, sys_user.description_, sys_user.signer_ AS signer_, 
				sys_user.create_at_, sys_user.login_at_, sys_user.password_at_
			FROM sys_user 
				LEFT JOIN sys_depart ON sys_user.depart_id_ = sys_depart.id
			WHERE sys_user.depart_id_ = ?
			ORDER BY sys_user.order_ ASC
		`
		args = append(args, departId)
	}

	return asql.Select(tx, query, args...)
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
	signer := ctx.PostFormValue("signer_")

	moveId := ctx.PostFormValue("webix_move_id")
	moveIndex := ctx.PostFormValue("webix_move_index")
	moveParent := ctx.PostFormValue("webix_move_parent")

	now := asql.GetNow()
	switch operation {
	case "insert":
		newId := asql.GenerateId()

		// 默认密码
		var dPwd string
		if err := asql.SelectRow(tx, "SELECT value_ FROM sys_setting WHERE field_ = ?", "password_default").Scan(&dPwd); err != nil {
			return nil, err
		}

		ePwd := base.Config.AesEncodeString(dPwd)

		query := `
		INSERT INTO sys_user(id, user_code_, user_name_, 
			account_id_, password_, depart_id_, sex_, is_depart_leader_, 
			valid_, classification_, telephone_, email_, 
			birth_, description_, signer_, order_, create_at_
		) VALUES (?,?,?, ?,?,?,?,?, ?,?,?,?, ?,?,?,?,?)
		`
		args := []interface{}{
			newId, userCode, userName,
			accountId, ePwd, departId, sex, isDepartLeader,
			valid, classification, telephone, email,
			birth, description, signer, asql.GenerateOrderId(), now,
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
			birth_ = ?, description_ = ?, signer_ = ?, update_at_ = ? 
		WHERE id = ?
		`
		args := []interface{}{
			userCode, userName, accountId,
			departId, sex, isDepartLeader, valid,
			classification, telephone, email,
			birth, description, signer, now,
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
	if err := asql.Update(tx, "UPDATE sys_user SET password_ = ?, password_at_ = NULL WHERE id = ?", ePwd, id); err != nil {
		return nil, err
	}

	return map[string]interface{}{"status": "success", "id": id}, nil
}

func (o *SysUsers) GetLoginByToken(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	userId, departId := ctx.UserId(), ctx.DepartId()

	// 获取用户的部门ID
	org, err := asql.QueryRelationParents(tx, "sys_depart", departId)
	if err != nil {
		return nil, err
	}

	org = append(org, userId)
	menus, err := menusByOrg(tx, org...)
	if err != nil {
		return nil, err
	}

	var tasksMaxActivated sql.NullString
	var tasksCount int

	// 查询待办事项条数
	if err := tx.QueryRow("SELECT MAX(activated_at_) AS activated_at,COUNT(1) AS count FROM wf_flow_task WHERE executor_user_id_ = ? AND status_ = ?", userId, "Executing").Scan(&tasksMaxActivated, &tasksCount); err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}

		tasksMaxActivated = sql.NullString{Valid: true, String: "1900-01-01 01:01:01"}
		tasksCount = 0
	}

	users, err := asql.Select(tx, "SELECT login_at_ AS login_at, password_at_ AS password_at FROM sys_user WHERE id = ?", userId)
	if err != nil || len(users) != 1 {
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("invalid user result set of size %d", len(users))
	}

	return map[string]interface{}{"user": users[0], "menus": menus, "tasks_max_activated": tasksMaxActivated.String, "tasks_count": tasksCount}, nil
}

func (o *SysUsers) PostChangedPassword(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	oPwd := ctx.PostFormValue("old_password")
	nPwd := ctx.PostFormValue("new_password")

	var uPwd string
	if err := asql.SelectRow(tx, "SELECT sys_user.password_ FROM sys_user WHERE sys_user.id = ?", ctx.UserId()).Scan(&uPwd); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("无效的登录用户")
		}

		return nil, err
	}

	// 密码比较
	if !strings.EqualFold(uPwd, base.Config.AesEncodeString(oPwd)) {
		return nil, errors.New("登录密码不正确")
	}

	// 配置信息
	sRes, err := asql.Select(tx, "SELECT field_, value_ FROM sys_setting WHERE field_ IN (?,?,?)", "password_min_length", "password_max_length", "token_expire")
	if err != nil {
		return nil, err
	}
	setting := base.ResAsMapStringInt(sRes, "field_", "value_")

	// 密码长度
	if len(nPwd) < setting["password_min_length"] || len(nPwd) > setting["password_max_length"] {
		return nil, fmt.Errorf("密码长度必须为%d位至%d位", setting["password_min_length"], setting["password_max_length"])
	}

	// 密码加密
	ePwd := base.Config.AesEncodeString(nPwd)
	if err := asql.Update(tx, "UPDATE sys_user SET password_ = ?, password_at_ = ? WHERE id = ?", ePwd, asql.GetNow(), ctx.UserId()); err != nil {
		return nil, err
	}

	// Token
	token := base.GenerateToken(ctx.UserId(), ctx.UserCode(), ctx.UserName(), ctx.DepartId(), ctx.DepartCode(), ctx.DepartName(), ePwd, ctx.UserAgent(), setting["token_expire"])
	expire := time.Now().Add(time.Duration(setting["token_expire"]) * time.Second)
	setCookie(ctx, "PHOENIX_LOGIN_TOKEN", token, expire)

	return map[string]string{"status": "success"}, nil
}

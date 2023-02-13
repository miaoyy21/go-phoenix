package xsys

import (
	"database/sql"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/handle"
)

type SysRoles struct {
}

func (o *SysRoles) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	res, err := asql.Select(tx, "SELECT id, code_, name_, description_, create_at_ FROM sys_role ORDER BY order_ ASC")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (o *SysRoles) Post(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	operation := ctx.PostFormValue("operation")

	id := ctx.PostFormValue("id")
	code := ctx.PostFormValue("code_")
	name := ctx.PostFormValue("name_")
	description := ctx.PostFormValue("description_")

	moveId := ctx.PostFormValue("webix_move_id")
	moveIndex := ctx.PostFormValue("webix_move_index")
	moveParent := ctx.PostFormValue("webix_move_parent")

	now := asql.GetNow()
	switch operation {
	case "insert":
		newId := asql.GenerateId()
		query := "INSERT INTO sys_role(id, code_, name_, description_, order_, create_at_) VALUES (?,?,?,?,?,?)"
		args := []interface{}{newId, code, name, description, asql.GenerateOrderId(), now}
		if err := asql.Insert(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "newid": newId, "create_at_": now}, nil
	case "update":
		query := "UPDATE sys_role SET code_ = ?, name_ = ?, description_ = ?, update_at_ = ? WHERE id = ?"
		args := []interface{}{code, name, description, now, id}
		if err := asql.Update(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "id": id, "update_at_": now}, nil
	case "delete":
		// 删除关联菜单
		if err := asql.Delete(tx, "DELETE FROM sys_role_menu WHERE role_id_ = ?", id); err != nil {
			return nil, err
		}

		// 删除角色
		if err := asql.Delete(tx, "DELETE FROM sys_role WHERE id = ?", id); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "id": id}, nil
	case "order":
		if err := asql.Order(tx, "sys_role", id, moveId, moveIndex, moveParent); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success"}, nil
	}

	return nil, fmt.Errorf("unrecognizable operation %s ", operation)
}

package xsys

import (
	"database/sql"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/handle"
)

type SysAutoNoKinds struct {
}

func (o *SysAutoNoKinds) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	query := "SELECT id, code_, name_, description_, create_at_ FROM sys_auto_no_kind ORDER BY order_ ASC"
	res, err := asql.Select(tx, query)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (o *SysAutoNoKinds) Post(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
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
		query := "INSERT INTO sys_auto_no_kind(id, code_, name_, description_, order_, create_at_) VALUES (?,?,?,?,?,?)"
		args := []interface{}{newId, code, name, description, asql.GenerateOrderId(), now}
		if err := asql.Insert(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "newid": newId, "create_at_": now}, nil
	case "update":
		query := "UPDATE sys_auto_no_kind SET code_ = ?, name_ = ?, description_ = ?, update_at_ = ? WHERE id = ?"
		args := []interface{}{code, name, description, now, id}
		if err := asql.Update(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "id": id, "update_at_": now}, nil
	case "delete":
		// 删除自动编码明细
		if err := asql.Delete(tx, "DELETE FROM sys_auto_no_item WHERE kind_id_ = ?", id); err != nil {
			return nil, err
		}

		// 删除自动编码
		if err := asql.Delete(tx, "DELETE FROM sys_auto_no_kind WHERE id = ?", id); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success"}, nil
	case "order":
		if err := asql.Order(tx, "sys_auto_no_kind", id, moveId, moveIndex, moveParent); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success"}, nil
	}

	return nil, fmt.Errorf("unrecognizable operation %s ", operation)
}

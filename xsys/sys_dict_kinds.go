package xsys

import (
	"database/sql"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/handle"
)

type SysDictKinds struct {
}

func (o *SysDictKinds) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	query := "SELECT id, code_, name_, description_, create_at_ FROM sys_dict_kind ORDER BY order_ ASC"

	return asql.Select(tx, query)
}

func (o *SysDictKinds) Post(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
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
		query := "INSERT INTO sys_dict_kind(id, code_, name_, description_, order_, create_at_) VALUES (?,?,?,?,?,?)"
		args := []interface{}{newId, code, name, description, asql.GenerateOrderId(), now}
		if err := asql.Insert(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "newid": newId, "create_at_": now}, nil
	case "update":
		query := "UPDATE sys_dict_kind SET code_ = ?, name_ = ?, description_ = ?, update_at_ = ? WHERE id = ?"
		args := []interface{}{code, name, description, now, id}
		if err := asql.Update(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "id": id, "update_at_": now}, nil
	case "delete":
		// 删除数据字典明细
		if err := asql.Delete(tx, "DELETE FROM sys_dict_item WHERE kind_id_ = ?", id); err != nil {
			return nil, err
		}

		// 删除数据字典
		if err := asql.Delete(tx, "DELETE FROM sys_dict_kind WHERE id = ?", id); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success"}, nil
	case "order":
		if err := asql.Order(tx, "sys_dict_kind", id, moveId, moveIndex, moveParent); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success"}, nil
	}

	return nil, fmt.Errorf("unrecognizable operation %s ", operation)
}

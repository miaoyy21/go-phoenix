package xsys

import (
	"database/sql"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/handle"
)

type SysAutoNoItems struct {
}

func (o *SysAutoNoItems) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	kindId := ctx.FormValue("kind_id")

	query := "SELECT id, kind_id_, code_, value_ FROM sys_auto_no_item WHERE kind_id_ = ? ORDER BY order_ ASC"
	res, err := asql.Select(tx, query, kindId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (o *SysAutoNoItems) Post(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	operation := ctx.PostFormValue("operation")

	id := ctx.PostFormValue("id")
	kindId := ctx.PostFormValue("kind_id_")
	code := ctx.PostFormValue("code_")
	value := ctx.PostFormValue("value_")

	moveId := ctx.PostFormValue("webix_move_id")
	moveIndex := ctx.PostFormValue("webix_move_index")
	moveParent := ctx.PostFormValue("webix_move_parent")

	now := asql.GetNow()
	switch operation {
	case "insert":
		newId := asql.GenerateId()
		query := "INSERT INTO sys_auto_no_item(id, kind_id_, code_, value_, order_, create_at_) VALUES (?,?,?,?,?,?)"
		args := []interface{}{newId, kindId, code, value, asql.GenerateOrderId(), now}
		if err := asql.Insert(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "newid": newId, "create_at_": now}, nil
	case "update":
		query := "UPDATE sys_auto_no_item SET code_ = ?, value_ = ?, update_at_ = ? WHERE id = ?"
		args := []interface{}{code, value, now, id}
		if err := asql.Update(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "id": id, "update_at_": now}, nil
	case "delete":
		if err := asql.Delete(tx, "DELETE FROM sys_auto_no_item WHERE id = ?", id); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success"}, nil
	case "order":
		if err := asql.Order(tx, "sys_auto_no_item", id, moveId, moveIndex, moveParent); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success"}, nil
	}

	return nil, fmt.Errorf("unrecognizable operation %s ", operation)
}

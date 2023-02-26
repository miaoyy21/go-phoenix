package xsys

import (
	"database/sql"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/handle"
)

type SysDictItems struct {
}

func (o *SysDictItems) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	kindId := ctx.FormValue("kind_id")
	kindCode := ctx.FormValue("kind_code")

	query, args := "invalid", make([]interface{}, 0)
	if len(kindId) > 0 {
		query = "SELECT id, kind_id_, code_, name_ FROM sys_dict_item WHERE kind_id_ = ? ORDER BY order_ ASC"
		args = append(args, kindId)
	} else if len(kindCode) > 0 {
		query = `
			SELECT sys_dict_item.code_ AS id, sys_dict_item.name_ AS value
			FROM sys_dict_kind, sys_dict_item
			WHERE sys_dict_kind.id = sys_dict_item.kind_id_
				AND sys_dict_kind.code_ = ?
		`
		args = append(args, kindCode)
	}

	return asql.Select(tx, query, args...)
}

func (o *SysDictItems) Post(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	operation := ctx.PostFormValue("operation")

	id := ctx.PostFormValue("id")
	kindId := ctx.PostFormValue("kind_id_")
	code := ctx.PostFormValue("code_")
	name := ctx.PostFormValue("name_")

	moveId := ctx.PostFormValue("webix_move_id")
	moveIndex := ctx.PostFormValue("webix_move_index")
	moveParent := ctx.PostFormValue("webix_move_parent")

	now := asql.GetNow()
	switch operation {
	case "insert":
		newId := asql.GenerateId()
		query := "INSERT INTO sys_dict_item(id, kind_id_, code_, name_, order_, create_at_) VALUES (?,?,?,?,?,?)"
		args := []interface{}{newId, kindId, code, name, asql.GenerateOrderId(), now}
		if err := asql.Insert(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "newid": newId, "create_at_": now}, nil
	case "update":
		query := "UPDATE sys_dict_item SET code_ = ?, name_ = ?, update_at_ = ? WHERE id = ?"
		args := []interface{}{code, name, now, id}
		if err := asql.Update(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "id": id, "update_at_": now}, nil
	case "delete":
		if err := asql.Delete(tx, "DELETE FROM sys_dict_item WHERE id = ?", id); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success"}, nil
	case "order":
		if err := asql.Order(tx, "sys_dict_item", id, moveId, moveIndex, moveParent); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success"}, nil
	}

	return nil, fmt.Errorf("unrecognizable operation %s ", operation)
}

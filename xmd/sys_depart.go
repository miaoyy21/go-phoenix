package xmd

import (
	"database/sql"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/handle"
)

type SysDeparts struct {
}

func (o *SysDeparts) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	query := `
		SELECT id, code_, name_, parent_id_, valid_, description_
		FROM sys_depart
		ORDER BY order_ ASC
	`
	res, err := asql.Select(tx, query)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (o *SysDeparts) Post(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	operation := ctx.PostFormValue("operation")

	id := ctx.PostFormValue("id")
	code := ctx.PostFormValue("code_")
	name := ctx.PostFormValue("name_")
	parentId := ctx.GetNullableFormValue("parent")
	description := ctx.PostFormValue("description_")
	valid := ctx.PostFormValue("valid_")

	moveId := ctx.PostFormValue("webix_move_id")
	moveIndex := ctx.PostFormValue("webix_move_index")
	moveParent := ctx.PostFormValue("webix_move_parent")

	now := asql.GetNow()
	switch operation {
	case "insert":
		newId := asql.GenerateId()
		query := "INSERT INTO sys_depart(id, code_, name_, parent_id_, valid_, description_, order_, create_at_) VALUES (?,?,?,?,?,?,?,?)"
		args := []interface{}{newId, code, name, parentId, valid, description, asql.GenerateOrderId(), now}
		if err := asql.Insert(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "newid": newId, "create_at_": now}, nil
	case "update":
		query := "UPDATE sys_depart SET code_ = ?, name_ = ?, valid_ = ?, description_ = ?, update_at_ = ? WHERE id = ?"
		args := []interface{}{code, name, valid, description, now, id}
		if err := asql.Update(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "id": id, "update_at_": now}, nil
	case "delete":
		if err := asql.Delete(tx, "DELETE FROM sys_depart WHERE id = ?", id); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "id": id}, nil
	case "order":
		if err := asql.Order(tx, "sys_depart", id, moveId, moveIndex, moveParent); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success"}, nil
	}

	return nil, fmt.Errorf("unrecognizable operation %s ", operation)
}
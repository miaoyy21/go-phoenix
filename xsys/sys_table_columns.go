package xsys

import (
	"database/sql"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/handle"
)

type SysTableColumns struct {
}

func (o *SysTableColumns) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	tableId := ctx.FormValue("table_id")

	query := `
		SELECT id, code_, name_, type_, is_sys_, description_, create_at_, update_at_ 
		FROM sys_table_column
		WHERE table_id_ = ? 
		ORDER BY order_ ASC
	`
	args := []interface{}{tableId}
	res, err := asql.Select(tx, query, args...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (o *SysTableColumns) Post(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	operation := ctx.PostFormValue("operation")

	id := ctx.PostFormValue("id")
	tableId := ctx.PostFormValue("table_id_")
	code := ctx.PostFormValue("code_")
	name := ctx.PostFormValue("name_")
	xType := ctx.PostFormValue("type_")
	description := ctx.PostFormValue("description_")

	moveId := ctx.PostFormValue("webix_move_id")
	moveIndex := ctx.PostFormValue("webix_move_index")
	moveParent := ctx.PostFormValue("webix_move_parent")

	now := asql.GetNow()
	switch operation {
	case "insert":
		newid := asql.GenerateId()

		query := "INSERT INTO sys_table_column(id, table_id_, is_sys_, code_, name_, type_, description_, order_, create_at_) VALUES (?,?,?,?,?,?,?,?,?)"
		args := []interface{}{newid, tableId, false, code, name, xType, description, asql.GenerateOrderId(), now}
		if err := asql.Insert(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "newid": newid, "create_at_": now}, nil
	case "update":
		query := "UPDATE sys_table_column SET code_ = ?, name_ = ?, type_ = ?, description_ = ?, update_at_ = ? WHERE id = ?"
		args := []interface{}{code, name, xType, description, now, id}
		if err := asql.Update(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "id": id, "update_at_": now}, nil
	case "delete":
		if err := asql.Delete(tx, "DELETE FROM sys_table_column WHERE id = ?", id); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "id": id}, nil
	case "order":
		if err := asql.Order(tx, "sys_table_column", id, moveId, moveIndex, moveParent); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success"}, nil
	}

	return nil, fmt.Errorf("unrecognizable operation %s ", operation)
}

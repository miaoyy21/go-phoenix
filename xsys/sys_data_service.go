package xsys

import (
	"database/sql"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/handle"
	"go-phoenix/rujs"
)

type SysDataService struct {
}

func (o *SysDataService) GetByTableId(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	tableId := ctx.FormValue("table_id")

	query := `
		SELECT id, code_, name_, method_, 
			source_, timeout_, create_at_, update_at_ 
		FROM sys_data_service 
		WHERE table_id_ = ? 
		ORDER BY order_ ASC
	`
	res, err := asql.Select(tx, query, tableId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (o *SysDataService) PostByTableId(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	operation := ctx.PostFormValue("operation")

	id := ctx.PostFormValue("id")
	tableId := ctx.PostFormValue("table_id_")
	code := ctx.PostFormValue("code_")
	name := ctx.PostFormValue("name_")
	method := ctx.PostFormValue("method_")
	timeout := ctx.PostFormValue("timeout_")
	source := ctx.PostFormValue("source_")

	moveId := ctx.PostFormValue("webix_move_id")
	moveIndex := ctx.PostFormValue("webix_move_index")
	moveParent := ctx.PostFormValue("webix_move_parent")

	now := asql.GetNow()
	switch operation {
	case "insert":
		newId := asql.GenerateId()
		query := "INSERT INTO sys_data_service(id, table_id_, code_, name_, method_, timeout_, source_, order_, create_at_) VALUES (?,?,?,?,?,?,?,?,?)"
		args := []interface{}{newId, tableId, code, name, method, timeout, source, asql.GenerateOrderId(), now}
		if err := asql.Insert(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "newid": newId, "create_at_": now}, nil
	case "update":
		query := "UPDATE sys_data_service SET code_ = ?, name_ = ?, method_ = ?, timeout_ = ?, source_ = ?, update_at_ = ? WHERE id = ?"
		args := []interface{}{code, name, method, timeout, source, now, id}
		if err := asql.Update(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "id": id, "update_at_": now}, nil
	case "delete":
		if err := asql.Delete(tx, "DELETE FROM sys_data_service WHERE id = ?", id); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success"}, nil
	case "order":
		if err := asql.Order(tx, "sys_data_service", id, moveId, moveIndex, moveParent); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success"}, nil
	}

	return nil, fmt.Errorf("unrecognizable operation %s ", operation)
}

func (o *SysDataService) PostParse(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	source := ctx.PostFormValue("source_")

	// 事务控制
	newTx, err := ctx.Begin()
	if err != nil {
		return nil, err
	}

	// 检测
	ctx.Reset(map[string]string{}, map[string]string{"id": "******", "operation": "delete"})
	if _, err := rujs.Run(newTx, ctx, source, 0, nil); err != nil {
		return nil, err
	}

	// 回滚
	if err := newTx.Rollback(); err != nil {
		return nil, err
	}

	return map[string]string{"status": "success"}, nil
}

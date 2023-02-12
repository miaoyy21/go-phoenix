package xsys

import (
	"database/sql"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/handle"
	"strings"
	"time"
)

type SysTables struct {
}

func (o *SysTables) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	query := "SELECT id, code_, name_, sync_status_, description_, create_at_, update_at_ FROM sys_table ORDER BY order_ ASC"
	res, err := asql.Select(tx, query)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (o *SysTables) Post(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	operation := ctx.PostFormValue("operation")

	id := ctx.PostFormValue("id")
	code := ctx.PostFormValue("code_")
	name := ctx.PostFormValue("name_")
	syncStatus := ctx.GetNullableFormValue("sync_status_")
	description := ctx.PostFormValue("description_")

	moveId := ctx.PostFormValue("webix_move_id")
	moveIndex := ctx.PostFormValue("webix_move_index")
	moveParent := ctx.PostFormValue("webix_move_parent")

	now := asql.GetNow()
	switch operation {
	case "insert":
		tableId := asql.GenerateId()

		// 数据库表
		query := "INSERT INTO sys_table(id, code_, name_, sync_status_, description_, order_, create_at_) VALUES (?,?,?,?,?,?,?)"
		args := []interface{}{tableId, code, name, syncStatus, description, asql.GenerateOrderId(), now}
		if err := asql.Insert(tx, query, args...); err != nil {
			return nil, err
		}

		// 数据库表的默认列
		for _, col := range asql.SysColumns {
			order := asql.GenerateOrderId()
			if !strings.EqualFold(col.Code, "id") {
				order = order + (time.Hour * 24 * 30).Nanoseconds()
			}

			query := "INSERT INTO sys_table_column(id, table_id_, is_sys_, code_, name_, type_, description_, order_, create_at_) VALUES (?,?,?,?,?,?,?,?,?)"
			args := []interface{}{asql.GenerateId(), tableId, '1', col.Code, col.Name, col.Type, col.Description, order, now}
			if err := asql.Insert(tx, query, args...); err != nil {
				return nil, err
			}
		}

		return map[string]interface{}{"status": "success", "newid": tableId, "create_at_": now}, nil
	case "update":
		query := "UPDATE sys_table SET code_ = ?, name_ = ?, sync_status_ = ?, description_ = ?, update_at_ = ? WHERE id = ?"
		args := []interface{}{code, name, syncStatus, description, now, id}
		if err := asql.Update(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "id": id, "update_at_": now}, nil
	case "delete":
		// 数据库表
		if err := asql.Delete(tx, "DELETE FROM sys_table WHERE id = ?", id); err != nil {
			return nil, err
		}

		// 数据库表的所有列
		if err := asql.Delete(tx, "DELETE FROM sys_table_column WHERE table_id_ = ?", id); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "id": id}, nil
	case "order":
		if err := asql.Order(tx, "sys_table", id, moveId, moveIndex, moveParent); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success"}, nil
	}

	return nil, fmt.Errorf("unrecognizable operation %s ", operation)
}

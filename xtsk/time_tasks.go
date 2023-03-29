package xtsk

import (
	"database/sql"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/handle"
	"strings"
	"time"
)

type TimeTasks struct {
}

func (o *TimeTasks) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	query := `
		SELECT sys_dict_item.code_ AS id, sys_dict_kind.code_ AS code, sys_dict_item.name_ AS value
		FROM sys_dict_kind, sys_dict_item
		WHERE sys_dict_kind.id = sys_dict_item.kind_id_
	`

	return asql.Select(tx, query)
}

func (o *TimeTasks) Post(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	operation := ctx.PostFormValue("operation")

	id := ctx.PostFormValue("id")
	code := ctx.PostFormValue("code_")
	name := ctx.PostFormValue("name_")
	syncStatus := ctx.PostFormValue("sync_status_")
	description := ctx.PostFormValue("description_")

	moveId := ctx.PostFormValue("webix_move_id")
	moveIndex := ctx.PostFormValue("webix_move_index")
	moveParent := ctx.PostFormValue("webix_move_parent")

	now := asql.GetNow()
	switch operation {
	case "insert":
		tableId := asql.GenerateId()
		creating := "Creating"

		// 数据库表
		query := "INSERT INTO sys_table(id, code_, name_, sync_status_, description_, order_, create_at_) VALUES (?,?,?,?,?,?,?)"
		args := []interface{}{tableId, code, name, creating, description, asql.GenerateOrderId(), now}
		if err := asql.Insert(tx, query, args...); err != nil {
			return nil, err
		}

		// 数据库表的默认列
		for _, col := range asql.SysColumns {
			order := asql.GenerateOrderId()
			if !strings.EqualFold(col.Code, "id") {
				order = order + (time.Hour * 24 * 365).Nanoseconds()
			}

			query := "INSERT INTO sys_table_column(id, table_id_, is_sys_, code_, name_, type_, description_, order_, create_at_) VALUES (?,?,?,?,?,?,?,?,?)"
			args := []interface{}{asql.GenerateId(), tableId, true, col.Code, col.Name, col.Type, col.Description, order, now}
			if err := asql.Insert(tx, query, args...); err != nil {
				return nil, err
			}
		}

		return map[string]interface{}{"status": "success", "newid": tableId, "sync_status_": creating, "create_at_": now}, nil
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

		return map[string]interface{}{"status": "success"}, nil
	case "order":
		if err := asql.Order(tx, "sys_table", id, moveId, moveIndex, moveParent); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success"}, nil
	}

	return nil, fmt.Errorf("unrecognizable operation %s ", operation)
}

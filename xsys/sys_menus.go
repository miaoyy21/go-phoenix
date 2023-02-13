package xsys

import (
	"database/sql"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/handle"
)

type SysMenus struct {
}

func (o *SysMenus) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	res, err := asql.Select(tx, "SELECT id, name_, parent_id_, menu_, icon_, valid_, description_ FROM sys_menu ORDER BY order_ ASC")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (o *SysMenus) Post(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	operation := ctx.PostFormValue("operation")

	id := ctx.PostFormValue("id")
	name := ctx.PostFormValue("name_")
	menu := ctx.PostFormValue("menu_")
	icon := ctx.PostFormValue("icon_")
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
		query := "INSERT INTO sys_menu(id, name_, parent_id_, menu_, icon_, valid_, description_, order_, create_at_) VALUES (?,?,?,?,?,?,?,?,?)"
		args := []interface{}{newId, name, parentId, menu, icon, valid, description, asql.GenerateOrderId(), now}
		if err := asql.Insert(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "newid": newId, "create_at_": now}, nil
	case "update":
		query := "UPDATE sys_menu SET name_ = ?, menu_ = ?, icon_ = ?, valid_ = ?, description_ = ?, update_at_ = ? WHERE id = ?"
		args := []interface{}{name, menu, icon, valid, description, now, id}
		if err := asql.Update(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "id": id, "update_at_": now}, nil
	case "delete":
		if err := asql.Delete(tx, "DELETE FROM sys_menu WHERE id = ?", id); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "id": id}, nil
	case "order":
		if err := asql.Order(tx, "sys_menu", id, moveId, moveIndex, moveParent); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success"}, nil
	}

	return nil, fmt.Errorf("unrecognizable operation %s ", operation)
}

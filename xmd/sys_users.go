package xmd

import (
	"database/sql"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/handle"
	"strings"
)

type SysUsers struct {
}

func (m *SysUsers) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	accountId := ctx.FormValue("account_id")
	scope := ctx.FormValue("scope")

	query, args := "invalid", make([]interface{}, 0)
	if strings.EqualFold(scope, "LOGIN") && len(accountId) > 0 {
		query = `	
			SELECT sys_depart.id, sys_depart.name_
			FROM sys_user,sys_depart
			WHERE sys_user.depart_id_ = sys_depart.id
				AND sys_user.account_id_ = ?
			ORDER BY sys_depart.order_ ASC
		`
		args = append(args, accountId)
	} else {
		//query = `
		//	SELECT id, code_, name_, description_, create_at_, update_at_
		//	FROM wf_diagram
		//	ORDER BY order_ ASC
		//`
	}

	res, err := asql.Select(tx, query, args...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *SysUsers) Post(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	operation := ctx.PostFormValue("operation")
	//
	//id := ctx.PostFormValue("id")
	//code := ctx.PostFormValue("code_")
	//name := ctx.PostFormValue("name_")
	//description := ctx.PostFormValue("description_")
	//
	//moveId := ctx.PostFormValue("webix_move_id")
	//moveIndex := ctx.PostFormValue("webix_move_index")
	//moveParent := ctx.PostFormValue("webix_move_parent")
	//
	//now := asql.GetNow()
	//switch operation {
	//case "insert":
	//	newId := asql.GenerateId()
	//	query := "INSERT INTO wf_diagram(id, code_, name_, description_, order_, create_at_) VALUES (?,?,?,?,?,?,?)"
	//	args := []interface{}{newId, code, name, description, asql.GenerateOrderId(), now}
	//	if err := asql.Insert(tx, query, args...); err != nil {
	//		return nil, err
	//	}
	//
	//	return map[string]interface{}{"status": "success", "newid": newId, "create_at_": now}, nil
	//case "update":
	//	query := "UPDATE wf_diagram SET code_ = ?, name_ = ?, description_ = ?, update_at_ = ? WHERE id = ?"
	//	args := []interface{}{code, name, description, now, id}
	//	if err := asql.Update(tx, query, args...); err != nil {
	//		return nil, err
	//	}
	//
	//	return map[string]interface{}{"status": "success", "id": id, "update_at_": now}, nil
	//case "delete":
	//	if err := asql.Delete(tx, "DELETE FROM wf_diagram WHERE id = ?", id); err != nil {
	//		return nil, err
	//	}
	//
	//	return map[string]interface{}{"status": "success", "id": id}, nil
	//case "order":
	//	if err := asql.Order(tx, "wf_diagram", id, moveId, moveIndex, moveParent); err != nil {
	//		return nil, err
	//	}
	//
	//	return map[string]interface{}{"status": "success"}, nil
	//}

	return nil, fmt.Errorf("unrecognizable operation %s ", operation)
}

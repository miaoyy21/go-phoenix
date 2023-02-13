package xsys

import (
	"database/sql"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/handle"
	"strings"
)

type SysDeparts struct {
}

func (o *SysDeparts) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	parentId := ctx.FormValue("parent_id")
	scope := ctx.FormValue("scope")

	query, args := "invalid", make([]interface{}, 0)
	if strings.EqualFold(scope, "KIDS") {
		if len(parentId) > 0 {
			query = `
			SELECT id, code_, name_, type_, kids_
			FROM (
				SELECT T.id, T.code_, T.name_, T.order_, 'depart' AS type_,
					CASE WHEN EXISTS (SELECT 1 FROM sys_depart X WHERE X.parent_id_ = T.id) 
						OR EXISTS (SELECT 1 FROM sys_user X WHERE X.depart_id_ = T.id) THEN '1' ELSE '0' END AS kids_
				FROM sys_depart T
				WHERE T.parent_id_ = ?
				UNION ALL
				SELECT T.id, T.user_code_, T.user_name_, T.order_, 'user','0'
				FROM sys_user T
				WHERE T.depart_id_ = ?
			) TX
			ORDER BY type_ ASC, order_ ASC
		`
			args = append(args, parentId, parentId)
		} else {
			query = `
			SELECT T.id, T.code_, T.name_, 'depart' AS type_,
				CASE WHEN EXISTS (SELECT 1 FROM sys_depart X WHERE X.parent_id_ = T.id) 
					OR EXISTS (SELECT 1 FROM sys_user X WHERE X.depart_id_ = T.id) THEN '1' ELSE '0' END AS kids_
			FROM sys_depart T
			WHERE T.parent_id_ IS NULL
			ORDER BY T.order_ ASC
			`
		}
	} else if strings.EqualFold(scope, "ORGANIZATION") {
		query = `
			SELECT id, name_, type_, parent_name_
			FROM (
				SELECT T.id, T.name_, T.order_, 'depart' AS type_, CASE WHEN X.id IS NULL THEN '-' ELSE X.name_ END AS parent_name_
				FROM sys_depart T
					LEFT JOIN sys_depart X ON T.parent_id_ = X.id
				UNION ALL
				SELECT T.id, T.user_name_, T.order_, 'user', CASE WHEN X.id IS NULL THEN '-' ELSE X.name_ END
				FROM sys_user T
					LEFT JOIN sys_depart X ON T.depart_id_ = X.id
			) TX
			ORDER BY type_ ASC, order_ ASC
		`
	} else {
		query = "SELECT id, code_, name_, parent_id_, valid_, description_ FROM sys_depart ORDER BY order_ ASC"
	}
	res, err := asql.Select(tx, query, args...)
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

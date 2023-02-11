package xmd

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/handle"
)

type SysOrganizationRoles struct {
}

func (o *SysOrganizationRoles) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	organizationId := ctx.FormValue("organization_id")
	//scope := ctx.FormValue("scope")

	query, args := "invalid", make([]interface{}, 0)
	if len(organizationId) > 0 {
		query = "SELECT role_id_ FROM sys_organization_role WHERE organization_id_ = ?"
		args = append(args, organizationId)
	} else {
		//query = `
		//		SELECT id, code_, name_
		//		FROM sys_role
		//		ORDER BY order_ ASC
		//	`
	}
	res, err := asql.Select(tx, query, args...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (o *SysOrganizationRoles) Post(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	operation := ctx.PostFormValue("operation")

	//id := ctx.PostFormValue("id")
	//code := ctx.PostFormValue("code_")
	//name := ctx.PostFormValue("name_")
	//parentId := ctx.GetNullableFormValue("parent")
	//description := ctx.PostFormValue("description_")
	//valid := ctx.PostFormValue("valid_")
	//
	//moveId := ctx.PostFormValue("webix_move_id")
	//moveIndex := ctx.PostFormValue("webix_move_index")
	//moveParent := ctx.PostFormValue("webix_move_parent")
	//
	//now := asql.GetNow()
	//switch operation {
	//case "insert":
	//	newId := asql.GenerateId()
	//	query := "INSERT INTO sys_depart(id, code_, name_, parent_id_, valid_, description_, order_, create_at_) VALUES (?,?,?,?,?,?,?,?)"
	//	args := []interface{}{newId, code, name, parentId, valid, description, asql.GenerateOrderId(), now}
	//	if err := asql.Insert(tx, query, args...); err != nil {
	//		return nil, err
	//	}
	//
	//	return map[string]interface{}{"status": "success", "newid": newId, "create_at_": now}, nil
	//case "update":
	//	query := "UPDATE sys_depart SET code_ = ?, name_ = ?, valid_ = ?, description_ = ?, update_at_ = ? WHERE id = ?"
	//	args := []interface{}{code, name, valid, description, now, id}
	//	if err := asql.Update(tx, query, args...); err != nil {
	//		return nil, err
	//	}
	//
	//	return map[string]interface{}{"status": "success", "id": id, "update_at_": now}, nil
	//case "delete":
	//	if err := asql.Delete(tx, "DELETE FROM sys_depart WHERE id = ?", id); err != nil {
	//		return nil, err
	//	}
	//
	//	return map[string]interface{}{"status": "success", "id": id}, nil
	//case "order":
	//	if err := asql.Order(tx, "sys_depart", id, moveId, moveIndex, moveParent); err != nil {
	//		return nil, err
	//	}
	//
	//	return map[string]interface{}{"status": "success"}, nil
	//}

	return nil, fmt.Errorf("unrecognizable operation %s ", operation)
}

func (o *SysOrganizationRoles) PostPatchRoles(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	organizationId := ctx.PostFormValue("organization_id")
	sRoles := ctx.PostFormValue("roles")

	var roles []string
	if err := json.Unmarshal([]byte(sRoles), &roles); err != nil {
		return nil, err
	}

	// 删除原有的组织权限
	dQuery := "DELETE FROM sys_organization_role WHERE organization_id_ = ?"
	if err := asql.Delete(tx, dQuery, organizationId); err != nil {
		return nil, err
	}

	// 添加新的组织权限
	iQuery := "INSERT INTO sys_organization_role(id, organization_id_, role_id_, create_at_) VALUES (?,?,?,?)"
	for _, role := range roles {
		if err := asql.Insert(tx, iQuery, asql.GenerateId(), organizationId, role, asql.GetNow()); err != nil {
			return nil, err
		}
	}

	return map[string]interface{}{"status": "success"}, nil
}

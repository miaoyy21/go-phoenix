package xsys

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/handle"
	"strings"
)

type SysOrganizationRoles struct {
}

func (o *SysOrganizationRoles) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	orgId := ctx.FormValue("organization_id")
	roleId := ctx.FormValue("role_id")

	query, args := "invalid", make([]interface{}, 0)
	if len(orgId) > 0 {
		query = "SELECT role_id_ FROM sys_organization_role WHERE organization_id_ = ?"
		args = append(args, orgId)
	} else if len(roleId) > 0 {
		query = `
			SELECT id, name_, type_, parent_name_
			FROM (
				SELECT sys_depart.id, sys_depart.name_, sys_depart.order_, 'depart' AS type_, '-' AS parent_name_
				FROM sys_organization_role, sys_depart
				WHERE sys_organization_role.organization_id_ = sys_depart.id 
					AND sys_organization_role.role_id_ = ?
				UNION ALL
				SELECT sys_user.id, sys_user.user_name_, sys_user.order_, 'user', sys_depart.name_
				FROM sys_organization_role, sys_user, sys_depart
				WHERE sys_organization_role.organization_id_ = sys_user.id 
					AND sys_user.depart_id_ = sys_depart.id
					AND sys_organization_role.role_id_ = ?
			) T
			ORDER BY type_ ASC, order_ ASC
		`
		args = append(args, roleId, roleId)
	}
	res, err := asql.Select(tx, query, args...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (o *SysOrganizationRoles) GetPermissions(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	orgId := ctx.FormValue("organization_id")

	var departId string

	sOrg := make([]interface{}, 0, 7)
	if err := asql.SelectRow(tx, "SELECT depart_id_ FROM sys_user WHERE id = ?", orgId).Scan(&departId); err != nil {
		if err == sql.ErrNoRows {
			// 说明查询部门的权限菜单
			org, err := asql.QueryRelationParents(tx, "sys_depart", orgId)
			if err != nil {
				return nil, err
			}

			sOrg = append(sOrg, org...)
		}
	} else {
		// 说明查询用户的权限菜单
		org, err := asql.QueryRelationParents(tx, "sys_depart", departId)
		if err != nil {
			return nil, err
		}

		sOrg = append(sOrg, orgId)
		sOrg = append(sOrg, org...)
	}

	return menusByOrg(tx, sOrg...)
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
	orgId := ctx.PostFormValue("organization_id")
	sRoles := ctx.PostFormValue("roles")

	var roles []string
	if err := json.Unmarshal([]byte(sRoles), &roles); err != nil {
		return nil, err
	}

	// 删除原有的组织权限
	dQuery := "DELETE FROM sys_organization_role WHERE organization_id_ = ?"
	if err := asql.Delete(tx, dQuery, orgId); err != nil {
		return nil, err
	}

	// 添加新的组织权限
	iQuery := "INSERT INTO sys_organization_role(id, organization_id_, role_id_, create_at_) VALUES (?,?,?,?)"
	for _, role := range roles {
		if err := asql.Insert(tx, iQuery, asql.GenerateId(), orgId, role, asql.GetNow()); err != nil {
			return nil, err
		}
	}

	return map[string]interface{}{"status": "success"}, nil
}

func (o *SysOrganizationRoles) PostPatchOrganization(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	operation := ctx.PostFormValue("operation")
	roleId := ctx.PostFormValue("role_id")
	sOrg := ctx.PostFormValue("organization")

	var org []string
	if err := json.Unmarshal([]byte(sOrg), &org); err != nil {
		return nil, err
	}

	switch operation {
	case "insert":
		if len(org) > 0 {
			// 构建参数
			args := make([]interface{}, 0, len(org)+1)
			args = append(args, roleId)
			for _, orgId := range org {
				args = append(args, orgId)
			}

			// 删除原有的组织权限
			dQuery := fmt.Sprintf("DELETE FROM sys_organization_role WHERE role_id_ = ? AND organization_id_ IN(?%s)", strings.Repeat(", ?", len(org)-1))
			if err := asql.Delete(tx, dQuery, args...); err != nil {
				return nil, err
			}

			// 添加新的组织权限
			iQuery := "INSERT INTO sys_organization_role(id, organization_id_, role_id_, create_at_) VALUES (?,?,?,?)"
			for _, orgId := range org {
				if err := asql.Insert(tx, iQuery, asql.GenerateId(), orgId, roleId, asql.GetNow()); err != nil {
					return nil, err
				}
			}
		}
	case "delete":
		// 删除原有的组织权限
		query := "DELETE FROM sys_organization_role WHERE organization_id_ = ? AND role_id_ = ?"
		for _, orgId := range org {
			if err := asql.Delete(tx, query, orgId, roleId); err != nil {
				return nil, err
			}
		}
	default:
		return nil, fmt.Errorf("unrecognizable operation %s ", operation)
	}

	return map[string]interface{}{"status": "success"}, nil
}

package xwf

import (
	"database/sql"
	"encoding/json"
	"go-phoenix/asql"
	"go-phoenix/handle"
)

type DiagramRoles struct {
}

func (o *DiagramRoles) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	roleId := ctx.FormValue("role_id")

	query, args := "invalid", make([]interface{}, 0)
	if len(roleId) > 0 {
		query = "SELECT diagram_id_ FROM sys_diagram_role WHERE role_id_ = ?"
		args = append(args, roleId)
	}

	return asql.Select(tx, query, args...)
}

func (o *DiagramRoles) PostPatchDiagrams(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	roleId := ctx.PostFormValue("role_id")
	sRoles := ctx.PostFormValue("diagrams")

	var diagrams []string
	if err := json.Unmarshal([]byte(sRoles), &diagrams); err != nil {
		return nil, err
	}

	// 删除原有的流程权限
	dQuery := "DELETE FROM sys_diagram_role WHERE role_id_ = ?"
	if err := asql.Delete(tx, dQuery, roleId); err != nil {
		return nil, err
	}

	// 添加新的流程权限
	iQuery := "INSERT INTO sys_diagram_role(id, role_id_, diagram_id_, create_at_) VALUES (?,?,?,?)"
	for _, diagram := range diagrams {
		if err := asql.Insert(tx, iQuery, asql.GenerateId(), roleId, diagram, asql.GetNow()); err != nil {
			return nil, err
		}
	}

	return map[string]interface{}{"status": "success"}, nil
}

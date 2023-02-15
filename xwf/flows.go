package xwf

import (
	"database/sql"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/handle"
	"go-phoenix/xwf/enum"
)

type Flows struct {
}

func (r *Flows) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	query, args := "invalid", make([]interface{}, 0)

	res, err := asql.Select(tx, query, args...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *Flows) Post(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	operation := ctx.PostFormValue("operation")

	id := ctx.PostFormValue("id")
	values := ctx.PostFormValue("values_")
	keyword := ctx.PostFormValue("keyword_")
	diagramId := ctx.PostFormValue("diagram_id_")

	switch operation {
	case "insert":
		now := asql.GetNow()
		newId := asql.GenerateId()

		// 流程配置
		var key int
		query := "SELECT start_key_ FROM wf_options_diagram WHERE diagram_id_ = ?"
		if err := asql.SelectRow(tx, query, diagramId).Scan(&key); err != nil {
			return nil, err
		}

		// 创建流程实例
		query = `
			INSERT INTO wf_flow(
				id, values_, diagram_id_, keyword_, 
				start_key_, status_, order_, create_at_,
				create_depart_id_, create_depart_code_, create_depart_name_,
				create_user_id_, create_user_code_, create_user_name_
			)
			VALUES(?,?,?,?, ?,?,?,?, ?,?,?, ?,?,?)
		`
		args := []interface{}{
			newId, values, diagramId, keyword,
			key, enum.FlowStatusDraft, asql.GenerateOrderId(), now,
			ctx.GetDepartId(), ctx.GetDepartCode(), ctx.GetDepartName(),
			ctx.GetUserId(), ctx.GetUserCode(), ctx.GetUserName(),
		}

		if err := asql.Insert(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "id": newId}, nil
	case "update":
		query := "UPDATE wf_diagram SET values = ?, keyword_ = ? WHERE id = ?"
		args := []interface{}{values, keyword, id}
		if err := asql.Update(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "id": id}, nil
	}

	return nil, fmt.Errorf("unrecognizable operation %s ", operation)
}

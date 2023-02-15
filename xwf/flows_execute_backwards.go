package xwf

import (
	"database/sql"
	"errors"
	"go-phoenix/asql"
	"go-phoenix/handle"
	"go-phoenix/xwf/enum"
)

// PostExecuteBackwards 用户流程的向后流程查询
func (o *Flows) PostExecuteBackwards(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	id := ctx.PostFormValue("id")         // 流转节点ID
	values := ctx.PostFormValue("values") // 表单数据

	var diagramId, flowId string
	var key int
	query := `
		SELECT diagram_id_, flow_id_, key_
		FROM wf_flow_node 
		WHERE id = ? AND executor_user_id_ = ? AND status_ = ?
	`
	if err := asql.SelectRow(tx, query, id, ctx.GetUserId(), enum.FlowNodeStatusExecuting).Scan(&diagramId, &flowId, &key); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("没有处理该待办事项权限")
		}

		return nil, err
	}

	return backwards(tx, ctx, diagramId, key, flowId, values)
}

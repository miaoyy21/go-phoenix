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
	id := ctx.PostFormValue("id") // 流转节点ID

	var diagramId, flowId, values string
	var key int
	query := `
		SELECT wf_flow_task.diagram_id_, wf_flow_task.flow_id_, wf_flow_task.key_, wf_flow.values_
		FROM wf_flow_task, wf_flow 
		WHERE wf_flow_task.flow_id_ = wf_flow.id 
			AND wf_flow_task.id = ? AND wf_flow_task.executor_user_id_ = ? AND wf_flow_task.status_ = ?
	`
	if err := asql.SelectRow(tx, query, id, ctx.UserId(), enum.FlowNodeStatusExecuting).Scan(&diagramId, &flowId, &key, &values); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("没有处理该待办事项权限")
		}

		return nil, err
	}

	return backwards(tx, ctx, diagramId, key, flowId, values)
}

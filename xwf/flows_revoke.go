package xwf

import (
	"database/sql"
	"errors"
	"go-phoenix/asql"
	"go-phoenix/base"
	"go-phoenix/handle"
	"go-phoenix/xwf/enum"
	"go-phoenix/xwf/flow"
)

// PostRevoke 流程撤回
func (r *Flows) PostRevoke(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	instanceId := ctx.PostFormValue("instanceId") // 流程实例ID
	values := ctx.PostFormValue("values")         // 表单数据

	// 是否有权限撤回
	var diagramId string
	var key int
	var status enum.FlowStatus
	query := `
		SELECT wf_flow.diagram_id_, wf_flow.start_key_, wf_flow.status_
		FROM wf_flow,wf_options_diagram 
		WHERE wf_flow.diagram_id_ = wf_options_diagram.diagram_id_
			AND wf_flow.instance_id_ = ? AND wf_flow.status_ = ? AND wf_flow.create_user_id_ = ?
	`
	args := []interface{}{instanceId, enum.FlowNodeStatusExecuting, ctx.GetUserId()}
	if err := asql.SelectRow(tx, query, args...).Scan(&diagramId, &key, &status); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("没有权限撤回流程实例")
		}

		return nil, err
	}

	// 只允许撤回执行中的流程实例
	if status != enum.FlowStatusExecuting {
		return nil, errors.New("只允许撤销执行中的流程实例")
	}

	// 节点
	node, err := flow.NewNode(tx, ctx, diagramId, key)
	if err != nil {
		return nil, err
	}

	// 必须是开始节点
	start, ok := node.(flow.StartFlowable)
	if !ok {
		return nil, errors.New("无法确定有效的开始节点")
	}

	// 流程是否允许撤回
	if !start.Revocable() {
		return nil, errors.New("流程已设定在启动后不允许撤回")
	}

	// 撤回
	if err := start.Revoke(instanceId, values); err != nil {
		return nil, err
	}

	// 更新流程状态
	empty := base.NewIntSet([]int{}).String()
	queryUpdate := "UPDATE wf_flow SET values_ = ?, executed_keys_ = ?, activated_keys_ = ?, active_at_ = ?, status_ = ? WHERE instance_id_ = ?"
	argsUpdate := []interface{}{values, empty, empty, asql.GetNow(), enum.FlowStatusRevoked, instanceId}
	if err := asql.Update(tx, queryUpdate, argsUpdate...); err != nil {
		return nil, err
	}

	return map[string]interface{}{"status": "success"}, nil
}

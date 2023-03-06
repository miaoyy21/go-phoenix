package xwf

import (
	"database/sql"
	"errors"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/handle"
	"go-phoenix/xwf/enum"
	"go-phoenix/xwf/flow"
)

// PostRemove 流程删除
func (o *Flows) PostRemove(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	id := ctx.PostFormValue("id") // 流程实例ID

	// 是否有权限撤回
	var diagramId, values string
	var key int
	var status enum.FlowStatus
	query := "SELECT diagram_id_, values_, start_key_, status_ FROM wf_flow WHERE id = ? AND create_user_id_ = ?"
	args := []interface{}{id, ctx.UserId()}
	if err := asql.SelectRow(tx, query, args...).Scan(&diagramId, &values, &key, &status); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("没有权限删除流程实例")
		}

		return nil, err
	}

	// 只允许删除 草稿、撤回和驳回 的流程实例
	if status != enum.FlowStatusRevoked && status != enum.FlowStatusDraft && status != enum.FlowStatusRejected {
		return nil, fmt.Errorf("只允许删除 草稿、撤回和驳回 的流程实例，当前状态为%q", status)
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

	// 删除
	if err := start.Remove(id, values); err != nil {
		return nil, err
	}

	// 更新流程状态
	queryUpdate := "UPDATE wf_flow SET active_at_ = ?, status_ = ?, status_text_ = ? WHERE id = ?"
	argsUpdate := []interface{}{asql.GetNow(), enum.FlowStatusDiscarded, "流程发起者已删除流程实例", id}
	if err := asql.Update(tx, queryUpdate, argsUpdate...); err != nil {
		return nil, err
	}

	return map[string]interface{}{"status": "success"}, nil
}

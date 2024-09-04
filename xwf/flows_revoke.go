package xwf

import (
	"database/sql"
	"errors"
	"go-phoenix/asql"
	"go-phoenix/handle"
	"go-phoenix/xwf/enum"
	"go-phoenix/xwf/flow"
)

// PostRevoke 流程撤回
func (o *Flows) PostRevoke(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	id := ctx.PostFormValue("id") // 流程实例ID

	// 是否有权限撤回
	var diagramId, values string
	var key int
	var status enum.FlowStatus
	query := "SELECT diagram_id_, values_, start_key_, status_ FROM wf_flow WHERE id = ? AND status_ = ? AND create_user_id_ = ?"
	args := []interface{}{id, enum.FlowNodeStatusExecuting, ctx.UserId()}
	if err := asql.SelectRow(tx, query, args...).Scan(&diagramId, &values, &key, &status); err != nil {
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
	node, err := flow.NewNode(tx, ctx, id, diagramId, key)
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
	if err := start.Revoke(id, values); err != nil {
		return nil, err
	}

	// 更新流程状态
	queryUpdate := "UPDATE wf_flow SET active_at_ = ?, status_ = ?, status_text_ = ? WHERE id = ?"
	argsUpdate := []interface{}{asql.GetNow(), enum.FlowStatusRevoked, "流程实例发起者 已撤回", id}
	if err := asql.Update(tx, queryUpdate, argsUpdate...); err != nil {
		return nil, err
	}

	return map[string]interface{}{"status": "success"}, nil
}

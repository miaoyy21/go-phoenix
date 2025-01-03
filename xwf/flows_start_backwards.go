package xwf

import (
	"database/sql"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/handle"
	"go-phoenix/xwf/enum"
)

// PostStartBackwards 启动流程的向后流程查询
func (o *Flows) PostStartBackwards(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	id := ctx.PostFormValue("id") // 流程实例ID

	// 流程实例是否已启动
	var diagramId, values string
	var status enum.FlowStatus
	if err := asql.SelectRow(tx, "SELECT diagram_id_, values_, status_ FROM wf_flow WHERE id = ?", id).Scan(&diagramId, &values, &status); err != nil {
		return nil, err
	}

	// 只能启动 草稿、撤回和驳回的流程
	if status != enum.FlowStatusRevoked && status != enum.FlowStatusDraft && status != enum.FlowStatusRejected {
		return nil, fmt.Errorf("流程实例已启动，当前状态为%q", status)
	}

	// 流程配置
	var key int
	query := "SELECT start_key_ FROM wf_options_diagram WHERE diagram_id_ = ?"
	if err := asql.SelectRow(tx, query, diagramId).Scan(&key); err != nil {
		return nil, err
	}

	return backwards(tx, ctx, diagramId, key, id, values)
}

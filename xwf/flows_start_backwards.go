package xwf

import (
	"database/sql"
	"errors"
	"go-phoenix/asql"
	"go-phoenix/handle"
	"go-phoenix/xwf/enum"
)

// PostStartBackwards 启动流程的向后流程查询
func (r *Flows) PostStartBackwards(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	instanceId := ctx.PostFormValue("instanceId")   // 流程实例ID
	diagramCode := ctx.PostFormValue("diagramCode") // 流程编码
	values := ctx.PostFormValue("values")           // 表单数据

	// 流程实例是否已启动
	var status enum.FlowStatus
	if err := asql.SelectRow(tx, "SELECT status_ FROM wf_flow WHERE instance_id_ = ?", instanceId).Scan(&status); err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
	} else if status != enum.FlowStatusRevoked {
		return nil, errors.New("流程实例已启动")
	}

	// 流程配置
	var diagramId string
	var key int
	query := "SELECT diagram_id_, start_key_ FROM wf_options_diagram WHERE diagram_code_ = ?"
	if err := asql.SelectRow(tx, query, diagramCode).Scan(&diagramId, &key); err != nil {
		return nil, err
	}

	return backwards(tx, ctx, diagramId, key, instanceId, values)
}

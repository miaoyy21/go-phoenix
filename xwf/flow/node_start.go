package flow

import (
	"go-phoenix/asql"
	"go-phoenix/rujs"
	"go-phoenix/xwf/enum"
)

// NodeStart 开始
type NodeStart struct {
	Node

	revocable      bool
	onRevokeScript string
}

func (node *NodeStart) Revocable() bool {
	return node.revocable
}

func (node *NodeStart) start(flowId string) error {
	now := asql.GetNow()

	// 创建开始节点
	queryNode := `
		INSERT INTO wf_flow_node(
			id, flow_id_, diagram_id_, 
			key_, category_, code_, name_, order_, 
			executor_user_id_, executor_user_name_, 
			activated_at_, executed_at_, status_, 
			executed_depart_id_, executed_depart_code_, executed_depart_name_, 
			executed_user_id_, executed_user_code_, executed_user_name_
		)
		VALUES(?,?,?, ?,?,?,?,?, ?,?, ?,?,?, ?,?,?, ?,?,?)
	`
	argsNode := []interface{}{
		asql.GenerateId(), flowId, node.diagramId,
		node.key, node.category, node.code, node.name, asql.GenerateOrderId(),
		node.ctx.GetUserId(), node.ctx.GetUserName(),
		now, now, enum.FlowNodeStatusExecutedAuto,
		node.ctx.GetDepartId(), node.ctx.GetDepartCode(), node.ctx.GetDepartName(),
		node.ctx.GetUserId(), node.ctx.GetUserCode(), node.ctx.GetUserName(),
	}

	return asql.Insert(node.tx, queryNode, argsNode...)
}

func (node *NodeStart) Start(flowId string, values string) error {
	// 执行开始前脚本
	if len(node.onBeforeScript) > 0 {
		if _, err := rujs.Run(node.tx, node.ctx, node.onBeforeScript, 0, flowReg(node, values)); err != nil {
			return err
		}
	}

	// 开始节点
	if err := node.start(flowId); err != nil {
		return err
	}

	// 执行开始后脚本
	if len(node.onAfterScript) > 0 {
		if _, err := rujs.Run(node.tx, node.ctx, node.onAfterScript, 0, flowReg(node, values)); err != nil {
			return err
		}
	}

	return nil
}

func (node *NodeStart) Revoke(flowId string, values string) error {

	// 驳回脚本
	if len(node.onRevokeScript) > 0 {
		if _, err := rujs.Run(node.tx, node.ctx, node.onRevokeScript, 0, flowReg(node, values)); err != nil {
			return err
		}
	}

	now := asql.GetNow()

	// 将激活节点作废
	queryNode := "UPDATE wf_flow_node SET canceled_at_ = ?, status_ = ? WHERE flow_id_ = ? AND status_ = ?"
	argsNode := []interface{}{now, enum.FlowNodeStatusCanceled, flowId, enum.FlowNodeStatusExecuting}
	if err := asql.Update(node.tx, queryNode, argsNode...); err != nil {
		return err
	}

	return nil
}

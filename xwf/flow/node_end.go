package flow

import (
	"go-phoenix/asql"
	"go-phoenix/rujs"
	"go-phoenix/xwf/enum"
)

// NodeEnd 结束
type NodeEnd struct {
	Node
}

func (node *NodeEnd) End(flowId string, values string) error {

	// 执行结束前脚本
	if len(node.onBeforeScript) > 0 {
		if _, err := rujs.Run(node.tx, node.ctx, node.onBeforeScript, 0, flowReg(node, values)); err != nil {
			return err
		}
	}

	now := asql.GetNow()

	// 创建结束节点
	query := `
		INSERT INTO wf_flow_task(
			id, flow_id_, diagram_id_, 
			key_, category_, code_, name_, order_, 
			executor_user_id_, executor_user_name_, 
			activated_at_, executed_at_, status_, 
			executed_depart_id_, executed_depart_code_, executed_depart_name_, 
			executed_user_id_, executed_user_code_, executed_user_name_
		)
		VALUES(?,?,?, ?,?,?,?,?, ?,?, ?,?,?, ?,?,?, ?,?,?)
	`
	args := []interface{}{
		asql.GenerateId(), flowId, node.diagramId,
		node.key, node.category, node.code, node.name, asql.GenerateOrderId(),
		node.ctx.UserId(), node.ctx.UserName(),
		now, now, enum.FlowNodeStatusExecutedAuto,
		node.ctx.DepartId(), node.ctx.DepartCode(), node.ctx.DepartName(),
		node.ctx.UserId(), node.ctx.UserCode(), node.ctx.UserName(),
	}
	if err := asql.Insert(node.tx, query, args...); err != nil {
		return err
	}

	// 执行结束后脚本
	if len(node.onAfterScript) > 0 {
		if _, err := rujs.Run(node.tx, node.ctx, node.onAfterScript, 0, flowReg(node, values)); err != nil {
			return err
		}
	}
	return nil
}

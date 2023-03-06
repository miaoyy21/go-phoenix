package flow

import (
	"go-phoenix/asql"
	"go-phoenix/xwf/enum"
)

// NodeBranch 结束
type NodeBranch struct {
	Node
}

func (node *NodeBranch) Branch(flowId string) error {
	now := asql.GetNow()
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

	return asql.Insert(node.tx, query, args...)
}

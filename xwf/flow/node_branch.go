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
	args := []interface{}{
		asql.GenerateId(), flowId, node.diagramId,
		node.key, node.category, node.code, node.name, asql.GenerateOrderId(),
		node.ctx.GetUserId(), node.ctx.GetUserName(),
		now, now, enum.FlowNodeStatusExecutedAuto,
		node.ctx.GetDepartId(), node.ctx.GetDepartCode(), node.ctx.GetDepartName(),
		node.ctx.GetUserId(), node.ctx.GetUserCode(), node.ctx.GetUserName(),
	}

	return asql.Insert(node.tx, query, args...)
}

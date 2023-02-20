package flow

import (
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/rujs"
	"go-phoenix/xwf/enum"
)

// NodeExecute 执行环节
type NodeExecute struct {
	Node

	rejectable           bool
	requireRejectComment bool

	onRejectScript string

	executorCustomNum     int
	executorSelectableNum int
	executorSavable       bool

	executorUsers       string
	executorNameUsers   string
	executorDeparts     string
	executorNameDeparts string
	executorRoles       string
	executorNameRoles   string
	executorPolicy      enum.ExecutorPolicy
	executorScript      string
}

func (node *NodeExecute) Rejectable() bool {
	return node.rejectable
}

func (node *NodeExecute) RequireRejectComment() bool {
	return node.requireRejectComment
}

func (node *NodeExecute) ExecuteStart(flowId string, executors []Executor) error {
	if len(executors) < 1 {
		return fmt.Errorf("%q 没有选择执行者", node.name)
	}

	// 创建待办事项
	now := asql.GetNow()
	executedId := asql.GenerateId()
	for _, executor := range executors {
		query := `
		INSERT INTO wf_flow_node(
			id, flow_id_, executed_id_, 
			diagram_id_, key_, category_, code_, name_, order_, 
			executor_user_id_, executor_user_name_, 
			activated_at_,  status_
		)
		VALUES(?,?,?, ?,?,?,?,?,?, ?,?, ?,?)
	`
		args := []interface{}{
			asql.GenerateId(), flowId, executedId,
			node.diagramId, node.key, node.category, node.code, node.name, asql.GenerateOrderId(),
			executor.Id, executor.Name,
			now, enum.FlowNodeStatusExecuting,
		}

		if err := asql.Insert(node.tx, query, args...); err != nil {
			return err
		}
	}

	return nil
}

func (node *NodeExecute) ExecuteAccept(id string, values string, comment string) error {
	// 执行前脚本
	if len(node.onBeforeScript) > 0 {
		if _, err := rujs.Run(node.tx, node.ctx, node.onBeforeScript, 0, flowReg(node, values)); err != nil {
			return err
		}
	}

	var executedId string

	query := "SELECT executed_id_ FROM wf_flow_node WHERE id = ?"
	if err := asql.SelectRow(node.tx, query, id).Scan(&executedId); err != nil {
		return err
	}

	now := asql.GetNow()

	// Executed
	queryExecuted := `
		UPDATE wf_flow_node 
		SET executed_at_ = ?, status_ = ?, comment_ = ?, 
			executed_depart_id_ = ?, executed_depart_code_ = ?, executed_depart_name_ = ?,
			executed_user_id_ = ?, executed_user_code_ = ?, executed_user_name_ = ?
		WHERE id = ?
	`
	argsExecuted := []interface{}{
		now, enum.FlowNodeStatusExecutedAccepted, comment,
		node.ctx.GetDepartId(), node.ctx.GetDepartCode(), node.ctx.GetDepartName(),
		node.ctx.GetUserId(), node.ctx.GetUserCode(), node.ctx.GetUserName(),
		id,
	}
	if err := asql.Update(node.tx, queryExecuted, argsExecuted...); err != nil {
		return err
	}

	// Canceled
	queryCanceled := "UPDATE wf_flow_node  SET canceled_at_ = ?, status_ = ? WHERE executed_id_ = ? AND id <> ? "
	if err := asql.Update(node.tx, queryCanceled, now, enum.FlowNodeStatusCanceled, executedId, id); err != nil {
		return err
	}

	// 执行后脚本
	if len(node.onAfterScript) > 0 {
		if _, err := rujs.Run(node.tx, node.ctx, node.onAfterScript, 0, flowReg(node, values)); err != nil {
			return err
		}
	}

	return nil
}

func (node *NodeExecute) ExecuteReject(id string, flowId string, values string, comment string) error {

	// 驳回脚本
	if len(node.onRejectScript) > 0 {
		if _, err := rujs.Run(node.tx, node.ctx, node.onRejectScript, 0, flowReg(node, values)); err != nil {
			return err
		}
	}

	now := asql.GetNow()

	// Rejected
	queryRejected := `
		UPDATE wf_flow_node 
		SET executed_at_ = ?, status_ = ?, comment_ = ?, 
			executed_depart_id_ = ?, executed_depart_code_ = ?, executed_depart_name_ = ?,
			executed_user_id_ = ?, executed_user_code_ = ?, executed_user_name_ = ?
		WHERE id = ?
	`
	argsRejected := []interface{}{
		now, enum.FlowNodeStatusExecutedRejected, comment,
		node.ctx.GetDepartId(), node.ctx.GetDepartCode(), node.ctx.GetDepartName(),
		node.ctx.GetUserId(), node.ctx.GetUserCode(), node.ctx.GetUserName(),
		id,
	}
	if err := asql.Update(node.tx, queryRejected, argsRejected...); err != nil {
		return err
	}

	// 将激活节点作废
	queryNode := "UPDATE wf_flow_node SET canceled_at_ = ?, status_ = ? WHERE flow_id_ = ? AND status_ = ?"
	argsNode := []interface{}{now, enum.FlowNodeStatusCanceled, flowId, enum.FlowNodeStatusExecuting}
	if err := asql.Update(node.tx, queryNode, argsNode...); err != nil {
		return err
	}

	return nil
}

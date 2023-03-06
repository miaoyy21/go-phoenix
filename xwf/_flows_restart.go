package xwf

//
//import (
//	"database/sql"
//	"encoding/json"
//	"errors"
//	"github.com/sirupsen/logrus"
//	"go-phoenix/asql"
//	"go-phoenix/base"
//	"go-phoenix/handle"
//	"go-phoenix/xwf/enum"
//	"go-phoenix/xwf/flow"
//)
//
//// PostRestart 流程重新启动
//func (r *Flows) PostRestart(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
//	instanceId := ctx.PostFormValue("instanceId") // 流程实例ID
//	diagramId := ctx.PostFormValue("diagramId")   // 流程ID
//	values := ctx.PostFormValue("values")         // 表单数据
//
//	// 后续节点
//	var backs []ExecuteBackward
//	if err := json.Unmarshal([]byte(ctx.PostFormValue("backwards")), &backs); err != nil {
//		return nil, err
//	}
//
//	// 流程实例是否已启动
//	var key int
//	if err := asql.SelectRow(tx, "SELECT start_key_ FROM wf_flow WHERE instance_id_ = ? AND status_ = ? AND create_user_id_ = ?", instanceId, enum.FlowStatusRejected, ctx.UserId()).Scan(&key); err != nil {
//		if err == sql.ErrNoRows {
//			return nil, errors.New("没有处理该待办事项权限")
//		}
//
//		return nil, err
//	}
//
//	exists, executed, activated := make(map[int]struct{}), base.NewIntSet([]int{}), base.NewIntSet([]int{})
//	for _, back := range backs {
//		for _, route := range back.Routes {
//			// 已经执行的节点忽略
//			if _, ok := exists[route]; ok {
//				continue
//			}
//
//			node, err := flow.NewNode(tx, ctx, diagramId, route)
//			if err != nil {
//				return nil, err
//			}
//
//			// Start
//			if start, ok := node.(flow.StartFlowable); ok {
//				executed.Append(route)
//				if err := start.Start(instanceId, values); err != nil {
//					return nil, err
//				}
//			}
//
//			// Execute Start
//			if execute, ok := node.(flow.ExecuteFlowable); ok {
//				if route != back.Key {
//					logrus.Panic("unreachable")
//				}
//
//				activated.Append(route)
//				if err := execute.ExecuteStart(instanceId, back.Executors); err != nil {
//					return nil, err
//				}
//			}
//
//			// Branch
//			if branch, ok := node.(flow.BranchFlowable); ok {
//				executed.Append(route)
//				if err := branch.Branch(instanceId); err != nil {
//					return nil, err
//				}
//			}
//
//			// End
//			if _, ok := node.(flow.EndFlowable); ok {
//				logrus.Panic("unreachable")
//			}
//
//			exists[route] = struct{}{}
//		}
//	}
//
//	// 更新流程实例状态
//	query := "UPDATE wf_flow SET values_ = ?, executed_keys_ = ?, activated_keys_ = ?, active_at_ = ?, status_ = ? WHERE instance_id_ = ?"
//	args := []interface{}{values, executed.String(), activated.String(), asql.GetNow(), enum.FlowStatusExecuting, instanceId}
//	if err := asql.Update(tx, query, args...); err != nil {
//		return nil, err
//	}
//
//	return map[string]interface{}{"status": "success"}, nil
//}

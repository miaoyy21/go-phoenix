package xwf

//
//import (
//	"database/sql"
//	"errors"
//	"go-phoenix/asql"
//	"go-phoenix/handle"
//	"go-phoenix/xwf/enum"
//)
//
//// PostRestartBackwards 重新启动流程的向后流程查询
//func (r *Flows) PostRestartBackwards(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
//	instanceId := ctx.PostFormValue("instanceId") // 流程实例ID
//	diagramId := ctx.PostFormValue("diagramId")   // 流程实例ID
//	values := ctx.PostFormValue("values")         // 表单数据
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
//	return backwards(tx, ctx, diagramId, key, instanceId, values)
//}

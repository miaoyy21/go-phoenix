package xwf

import (
	"database/sql"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/base"
	"go-phoenix/handle"
	"go-phoenix/xwf/enum"
	"strings"
)

type Flows struct {
}

func (o *Flows) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	diagramId := ctx.FormValue("diagram_id")
	status := ctx.FormValue("status")

	query, args := "invalid", make([]interface{}, 0)
	switch status {
	case "DraftRevokedRejected":
		query = `
			SELECT id, keyword_, values_md5_, executed_keys_, activated_keys_, 
				status_, status_text_, create_at_, start_at_, active_at_, end_at_ 
			FROM wf_flow 
			WHERE diagram_id_ = ? AND create_user_id_ = ? AND status_ IN(?,?,?)
			ORDER BY wf_flow.end_at_, wf_flow.active_at_ DESC, wf_flow.create_at_ DESC
		`
		args = append(args, diagramId, ctx.GetUserId(), enum.FlowStatusDraft, enum.FlowStatusRevoked, enum.FlowStatusRejected)
	default:
		query = `
			SELECT id, keyword_, values_md5_, executed_keys_, activated_keys_, 
				status_, status_text_, create_at_, start_at_, active_at_, end_at_ 
			FROM wf_flow 
			WHERE diagram_id_ = ? AND create_user_id_ = ? AND status_ = ?
			ORDER BY wf_flow.end_at_, wf_flow.active_at_ DESC, wf_flow.create_at_ DESC
		`
		args = append(args, diagramId, ctx.GetUserId(), status)
	}

	return asql.Select(tx, query, args...)
}

func (o *Flows) GetModelValues(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	id := ctx.FormValue("id")

	var model, values string

	query := `
		SELECT wf_options_diagram.model_,wf_flow.values_ 
		FROM wf_flow, wf_options_diagram 
		WHERE wf_options_diagram.diagram_id_ = wf_flow.diagram_id_ 
			AND wf_flow.id = ?
	`
	if err := asql.SelectRow(tx, query, id).Scan(&model, &values); err != nil {
		return nil, err
	}

	return map[string]string{"values": values, "model": model}, nil
}

type Summary struct {
	Id    string `json:"id,omitempty"`
	Icon  string `json:"icon,omitempty"`
	Value string `json:"value,omitempty"`
	Badge int    `json:"badge,omitempty"`

	Template string `json:"$template,omitempty"`
}

func (o *Flows) GetSummary(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	diagramId := ctx.FormValue("diagram_id")

	query := `
		SELECT status_, COUNT(1) AS count_ 
		FROM wf_flow 
		WHERE diagram_id_ = ? AND create_user_id_ = ?
		GROUP BY status_
	`
	res, err := asql.Select(tx, query, diagramId, ctx.GetUserId())
	if err != nil {
		return nil, err
	}

	// 汇总
	sBadge := base.ResAsMapStringInt(res, "status_", "count_")

	return []Summary{
		{
			Id:    "DraftRevokedRejected",
			Value: "草稿箱",
			Icon:  "mdi mdi-progress-wrench",
			Badge: sBadge[string(enum.FlowStatusDraft)] + sBadge[string(enum.FlowStatusRevoked)] + sBadge[string(enum.FlowStatusRejected)],
		},
		{Template: "Separator"},
		{Id: "Executing", Value: "执行中", Icon: "mdi mdi-progress-upload", Badge: sBadge[string(enum.FlowStatusExecuting)]},
		{Template: "Separator"},
		{Id: "Finished", Value: "已结束", Icon: "mdi mdi-progress-check", Badge: sBadge[string(enum.FlowStatusFinished)]},
		{Template: "Separator"},
		{Id: "Suspended", Value: "已挂起", Icon: "mdi mdi-progress-question", Badge: sBadge[string(enum.FlowStatusSuspended)]},
		{Template: "Separator"},
	}, nil
}

func (o *Flows) Post(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	operation := ctx.PostFormValue("operation")

	id := ctx.PostFormValue("id")
	values := ctx.PostFormValue("values_")
	valuesMd5 := ctx.PostFormValue("values_md5_")
	keyword := ctx.PostFormValue("keyword_")
	diagramId := ctx.PostFormValue("diagram_id_")

	switch operation {
	case "insert":
		now := asql.GetNow()
		newId := asql.GenerateId()

		// 流程配置
		var key int
		query := "SELECT start_key_ FROM wf_options_diagram WHERE diagram_id_ = ?"
		if err := asql.SelectRow(tx, query, diagramId).Scan(&key); err != nil {
			return nil, err
		}

		// 创建流程实例
		query = `
			INSERT INTO wf_flow(
				id, values_, values_md5_, diagram_id_, keyword_, 
				start_key_, status_, status_text_, order_, create_at_,
				create_depart_id_, create_depart_code_, create_depart_name_,
				create_user_id_, create_user_code_, create_user_name_
			)
			VALUES(?,?,?,?,?, ?,?,?,?,?, ?,?,?, ?,?,?)
		`
		args := []interface{}{
			newId, values, valuesMd5, diagramId, keyword,
			key, enum.FlowStatusDraft, "等待流程实例启动", asql.GenerateOrderId(), now,
			ctx.GetDepartId(), ctx.GetDepartCode(), ctx.GetDepartName(),
			ctx.GetUserId(), ctx.GetUserCode(), ctx.GetUserName(),
		}

		if err := asql.Insert(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "id": newId}, nil
	case "update":
		query := "UPDATE wf_flow SET values_ = ?, values_md5_ = ?, keyword_ = ? WHERE id = ?"
		args := []interface{}{values, valuesMd5, keyword, id}
		if err := asql.Update(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "id": id}, nil
	}

	return nil, fmt.Errorf("unrecognizable operation %s ", operation)
}

func (o *Flows) executors(tx *sql.Tx, flowId string) ([]string, error) {
	max := 3
	query := `
 		SELECT wf_options_node.name_, wf_flow_node.executor_user_name_ 
		FROM wf_flow_node, wf_options_node 
		WHERE wf_flow_node.diagram_id_ = wf_options_node.diagram_id_ 
			AND wf_flow_node.key_ = wf_options_node.key_ 
			AND wf_flow_node.flow_id_ = ? AND wf_flow_node.status_ = ?
	`
	res, err := asql.Select(tx, query, flowId, enum.FlowNodeStatusExecuting)
	if err != nil {
		return nil, err
	}

	// 无效实例ID
	if len(res) <= 0 {
		return nil, fmt.Errorf("没有找到流程实例ID%q的执行者", flowId)
	}

	isMore := false
	names := make(map[string][]string)
	for i, row := range res {
		if i >= max {
			isMore = true
			break
		}

		users, ok := names[row["name_"]]
		if !ok {
			users = make([]string, 0)
		}

		users = append(users, row["executor_user_name_"])
		names[row["name_"]] = users
	}

	texts := make([]string, 0, len(names))
	for name, users := range names {
		texts = append(texts, fmt.Sprintf("[ %s ] %s", name, strings.Join(users, ",")))
	}

	// 是否超过10个
	if isMore {
		texts = append(texts, fmt.Sprintf("等%d位执行者", len(res)-max))
	}

	return texts, nil
}

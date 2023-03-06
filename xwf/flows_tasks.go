package xwf

import (
	"database/sql"
	"go-phoenix/asql"
	"go-phoenix/handle"
)

func (o *Flows) GetTasks(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	status := ctx.FormValue("status")

	query := `
	SELECT wf_flow_task.id,wf_flow.id AS flow_id_, wf_options_diagram.diagram_id_, 
		wf_options_diagram.diagram_code_, wf_options_diagram.diagram_name_, 
		wf_flow_task.code_, wf_flow_task.name_, 
		wf_options_diagram.keyword_, wf_flow.keyword_ AS keyword_text_, 
		wf_flow.executed_keys_, wf_flow.activated_keys_, 
		wf_flow.status_, wf_flow.status_text_, 
		wf_flow_task.status_ AS task_status_, wf_flow_task.comment_, 
		wf_flow.create_depart_name_, wf_flow.create_user_name_,
		wf_flow.start_at_, wf_flow_task.activated_at_, wf_flow_task.executed_at_
	FROM wf_flow_task, wf_flow, wf_options_diagram
	WHERE wf_flow_task.flow_id_ = wf_flow.id AND wf_flow.diagram_id_ = wf_options_diagram.diagram_id_
		AND wf_flow_task.executor_user_id_ = ? AND wf_flow_task.status_ = ?
	ORDER BY wf_flow.end_at_ DESC, wf_flow.active_at_ DESC, wf_flow.create_at_ DESC
		`
	args := []interface{}{ctx.UserId(), status}

	return asql.Select(tx, query, args...)
}

package asql

import (
	"database/sql"
)

func QueryFlowsExecutingCount(tx *sql.Tx, userId string) (interface{}, error) {
	var count int

	query := "SELECT COUNT(1) AS count_ FROM wf_flow_node WHERE executor_user_id_ = ? AND status_ = ?"
	if err := SelectRow(tx, query, userId, "Executing").Scan(&count); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return count, nil
}

func QueryFlowsByExecuting(tx *sql.Tx, userId string, status string) (interface{}, error) {
	query := `
	SELECT wf_flow_node.id, wf_options_diagram.diagram_id_, wf_options_diagram.diagram_code_, wf_options_diagram.diagram_name_, wf_flow_node.name_, 
		wf_options_diagram.keyword_, wf_flow.keyword_ AS keyword_text_, wf_flow.values_md5_, wf_flow.executed_keys_, wf_flow.activated_keys_, 
		wf_flow.status_, wf_flow.status_text_, 
		wf_flow_node.status_ AS node_status_, wf_flow_node.comment_, 
		wf_flow.create_depart_name_, wf_flow.create_user_name_,
		wf_flow.start_at_, wf_flow_node.activated_at_, wf_flow_node.executed_at_
	FROM wf_flow_node, wf_flow, wf_options_diagram
	WHERE wf_flow_node.flow_id_ = wf_flow.id AND wf_flow.diagram_id_ = wf_options_diagram.diagram_id_
		AND wf_flow_node.executor_user_id_ = ? AND wf_flow_node.status_ = ?
	ORDER BY wf_flow.end_at_ DESC, wf_flow.active_at_ DESC, wf_flow.create_at_ DESC
		`
	args := []interface{}{userId, "Executing"}

	return Select(tx, query, args...)
}

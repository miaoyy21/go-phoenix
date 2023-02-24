package mdg

import (
	"database/sql"
	"errors"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/xwf/enum"
	"strings"
)

func Publish(tx *sql.Tx, id, code, name string, smo string, order int64, model Model, options Options) error {

	// 获取各节点的名称
	names := make(map[int]string)
	for _, node := range model.Nodes {
		if len(strings.TrimSpace(node.Text)) > 0 {
			names[node.Key] = strings.TrimSpace(node.Text)
		}
	}

	// 获取连接节点的连接线上的脚本
	linkScripts := make(map[int]map[int]string)
	for _, link := range options.Links {
		from, ok := linkScripts[link.From]
		if !ok {
			from = make(map[int]string)
		}
		from[link.To] = link.OnScript

		linkScripts[link.From] = from
	}

	start, end := 0, make([]int, 0)
	codes, nodes := make(map[string]struct{}), make(map[int]OptionsNode)
	for _, node := range options.Nodes {
		// 流程图的开始入口只能有一个
		if node.Category == enum.CategoryStart {
			if start != 0 {
				return errors.New("只允许存在1个开始节点")
			}

			start = node.Key
		}

		// 必须命名
		if _, ok := names[node.Key]; !ok {
			return fmt.Errorf("节点%q没有命名 ", node.Key)
		}

		// 是否存在相同编码的节点
		if _, ok := codes[node.Code]; ok {
			return fmt.Errorf("不同节点拥有相同编码%q", node.Code)
		}

		// 允许存在多个结束节点
		if node.Category == enum.CategoryEnd {
			end = append(end, node.Key)
		}

		// 是否设置执行者
		if node.Category == enum.CategoryExecute && node.ExecutorPolicy == enum.ExecutorPolicyNone &&
			len(node.ExecutorDeparts)+len(node.ExecutorUsers)+len(node.ExecutorRoles) == 0 {
			return fmt.Errorf("执行环节【%s】没有完成执行者设置", names[node.Key])
		}

		nodes[node.Key] = node
		if len(node.Code) > 0 {
			codes[node.Code] = struct{}{}
		}
	}

	// 是否存在开始节点
	if start == 0 {
		return errors.New("必须存在1个开始节点")
	}

	// 是否存在结束节点
	if len(end) < 1 {
		return errors.New("至少存在1个结束节点")
	}

	links := make(map[int]map[int]struct{})
	tos := make(map[int]struct{})
	for _, link := range model.Links {

		// 无效的连接线
		if _, ok := nodes[link.From]; !ok {
			return fmt.Errorf("存在无效的连接线%q", link.From)
		}
		if _, ok := nodes[link.To]; !ok {
			return fmt.Errorf("存在无效的连接线%q", link.To)
		}

		// 不允许来源节点和目标节点相同的连接线
		if _, ok := links[link.From][link.To]; ok {
			return fmt.Errorf("%q->%q: 连接线重复", names[link.From], names[link.To])
		}

		// 不允许来源节点和目标节点形成环路
		if _, ok := links[link.To][link.From]; ok {
			// 如果是条件分支参与，可能会打破环路条件
			if nodes[link.From].Category != enum.CategoryBranch && nodes[link.To].Category != enum.CategoryBranch {
				return fmt.Errorf("%q-%q: 连接线形成环路", names[link.From], names[link.To])
			}
		}

		from, ok := links[link.From]
		if !ok {
			from = make(map[int]struct{})
		}
		from[link.To] = struct{}{}

		links[link.From] = from
		tos[link.To] = struct{}{}
	}

	for _, node := range options.Nodes {
		// 开始节点
		if node.Category == enum.CategoryStart {
			if _, ok := tos[node.Key]; ok {
				return errors.New("开始节点不允许存在输入节点")
			}

			if _, ok := links[node.Key]; !ok {
				return errors.New("无效的开始节点")
			}
		}

		// 执行环节
		if node.Category == enum.CategoryExecute {
			if _, ok := tos[node.Key]; !ok {
				return fmt.Errorf("执行环节%q: 无输入节点", names[node.Key])
			}

			if _, ok := links[node.Key]; !ok {
				return fmt.Errorf("执行环节%q: 无输出节点", names[node.Key])
			}
		}

		// 条件分支
		if node.Category == enum.CategoryBranch {
			if _, ok := tos[node.Key]; !ok {
				return fmt.Errorf("条件分支%q: 无输入节点", names[node.Key])
			}

			if _, ok := links[node.Key]; !ok {
				return fmt.Errorf("条件分支%q: 无输出节点", names[node.Key])
			}
		}

		// 结束节点
		if node.Category == enum.CategoryEnd {
			if _, ok := tos[node.Key]; !ok {
				return errors.New("无效的结束节点")
			}

			if _, ok := links[node.Key]; ok {
				return errors.New("结束节点不允许存在输出节点")
			}
		}
	}

	// 删除流程图编译记录
	if err := asql.Delete(tx, "DELETE FROM wf_options_node WHERE diagram_id_ = ?", id); err != nil {
		return err
	}
	if err := asql.Delete(tx, "DELETE FROM wf_options_link WHERE diagram_id_ = ?", id); err != nil {
		return err
	}

	var query string
	var args []interface{}

	// Insert Node Options
	for _, node := range options.Nodes {
		query = `
			INSERT INTO wf_options_node(id, diagram_id_, key_, category_, code_, name_,
				rejectable_, require_reject_comment_, revocable_, 
				on_reject_script_,on_revoke_script_, on_remove_script_,
				on_before_script_, on_after_script_, 
				executor_custom_num_, executor_selectable_num_, executor_savable_,
				executor_users_, executor_name_users_, 
				executor_departs_, executor_name_departs_, 
				executor_roles_, executor_name_roles_, 
				executor_policy_, executor_script_)
			VALUES (?,?,?,?,?,?, ?,?,?, ?,?,?, ?,?, ?,?,?, ?,?, ?,?, ?,?, ?,?)`
		args = []interface{}{
			asql.GenerateId(), id, node.Key, node.Category, node.Code, names[node.Key],
			node.Rejectable, node.RequireRejectComment, node.Revocable,
			node.OnRejectScript, node.OnRevokeScript, node.OnRemoveScript,
			node.OnBeforeScript, node.OnAfterScript,
			node.ExecutorCustomNum, node.ExecutorSelectableNum, node.ExecutorSavable,
			node.ExecutorUsers, node.ExecutorNameUsers,
			node.ExecutorDeparts, node.ExecutorNameDeparts,
			node.ExecutorRoles, node.ExecutorNameRoles,
			node.ExecutorPolicy, node.ExecutorScript,
		}

		if err := asql.Insert(tx, query, args...); err != nil {
			return err
		}
	}

	// Insert Link Options
	for _, link := range model.Links {
		onScript, ok := linkScripts[link.From][link.To]
		if !ok || len(onScript) < 1 {
			onScript = "true"
		}

		query = "INSERT INTO wf_options_link(id, diagram_id_, from_key_, to_key_, on_script_) VALUES (?,?,?,?,?)"
		args = []interface{}{asql.GenerateId(), id, link.From, link.To, onScript}

		if err := asql.Insert(tx, query, args...); err != nil {
			return err
		}
	}

	var published string
	if err := asql.SelectRow(tx, "SELECT id FROM wf_options_diagram WHERE diagram_id_ = ? ", id).Scan(&published); err != nil {
		if err == sql.ErrNoRows {

			// Insert Diagram Options
			query = "INSERT INTO wf_options_diagram(id, diagram_id_, diagram_code_, diagram_name_, model_, keyword_, icon_, description_, exceed_days_, start_key_, order_) VALUES (?,?,?,?,?,?,?,?,?,?,?)"
			args = []interface{}{asql.GenerateId(), id, code, name, smo, options.Diagram.Keyword, options.Diagram.Icon, options.Diagram.Description, options.Diagram.ExceedDays, start, order}

			return asql.Insert(tx, query, args...)
		}

		return err
	}

	// Update Diagram Options
	query = `
		UPDATE wf_options_diagram 
		SET diagram_code_ = ?, diagram_name_ = ?, model_ = ?,
			keyword_ = ?, icon_ = ?, description_ = ?,
			exceed_days_ = ?, start_key_ = ?, order_ = ?
		WHERE diagram_id_ = ?
	`
	args = []interface{}{code, name, smo,
		options.Diagram.Keyword, options.Diagram.Icon, options.Diagram.Description,
		options.Diagram.ExceedDays, start, order,
		id,
	}

	return asql.Update(tx, query, args...)
}

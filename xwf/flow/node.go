package flow

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-phoenix/asql"
	"go-phoenix/base"
	"go-phoenix/handle"
	"go-phoenix/rujs"
	"go-phoenix/xwf/enum"
)

// 接口检测
var _ StartFlowable = &NodeStart{}
var _ ExecuteFlowable = &NodeExecute{}
var _ BranchFlowable = &NodeBranch{}
var _ EndFlowable = &NodeEnd{}

type Node struct {
	tx  *sql.Tx
	ctx *handle.Context

	diagramId string
	key       int
	category  enum.Category
	code      string
	name      string

	onBeforeScript string
	onAfterScript  string
	onRejectScript string
}

func NewNode(tx *sql.Tx, ctx *handle.Context, diagramId string, key int) (Flowable, error) {
	var category enum.Category
	var code, name string
	var rejectable, requireRejectComment, revocable bool
	var executorCustomNum, executorSelectableNum int
	var executorSavable bool
	var onRemoveScript string
	var onRejectScript, onRevokeScript string
	var onBeforeScript, onAfterScript string
	var executorUsers, executorNameUsers string
	var executorDeparts, executorNameDeparts string
	var executorRoles, executorNameRoles string
	var executorPolicy, executorScript string

	query := `
		SELECT category_, code_, name_, 
			rejectable_, require_reject_comment_, revocable_, on_remove_script_,
			on_before_script_, on_after_script_, on_reject_script_, on_revoke_script_,
			executor_custom_num_, executor_selectable_num_, executor_savable_,
			executor_users_, executor_name_users_, 
			executor_departs_, executor_name_departs_, 
			executor_roles_, executor_name_roles_, 
			executor_policy_, executor_script_
		FROM wf_options_node
		WHERE diagram_id_ = ? AND key_ = ?
	`
	args := []interface{}{diagramId, key}
	if err := asql.SelectRow(tx, query, args...).Scan(
		&category, &code, &name,
		&rejectable, &requireRejectComment, &revocable, &onRemoveScript,
		&onBeforeScript, &onAfterScript, &onRejectScript, &onRevokeScript,
		&executorCustomNum, &executorSelectableNum, &executorSavable,
		&executorUsers, &executorNameUsers,
		&executorDeparts, &executorNameDeparts,
		&executorRoles, &executorNameRoles,
		&executorPolicy, &executorScript,
	); err != nil {
		return nil, err
	}

	node := Node{
		tx:  tx,
		ctx: ctx,

		diagramId: diagramId,
		key:       key,
		category:  category,
		code:      code,
		name:      name,

		onBeforeScript: onBeforeScript,
		onAfterScript:  onAfterScript,
	}

	switch category {
	case enum.CategoryStart:
		return &NodeStart{
			Node:           node,
			revocable:      revocable,
			onRevokeScript: onRevokeScript,
			onRemoveScript: onRemoveScript,
		}, nil
	case enum.CategoryEnd:
		return &NodeEnd{Node: node}, nil
	case enum.CategoryBranch:
		return &NodeBranch{Node: node}, nil
	case enum.CategoryExecute:
		return &NodeExecute{
			Node: node,

			rejectable:           rejectable,
			requireRejectComment: requireRejectComment,

			executorCustomNum:     executorCustomNum,
			executorSelectableNum: executorSelectableNum,
			executorSavable:       executorSavable,

			onRejectScript: onRejectScript,

			executorUsers:       executorUsers,
			executorNameUsers:   executorNameUsers,
			executorDeparts:     executorDeparts,
			executorNameDeparts: executorNameDeparts,
			executorRoles:       executorRoles,
			executorNameRoles:   executorNameRoles,
			executorPolicy:      enum.ExecutorPolicy(executorPolicy),
			executorScript:      executorScript,
		}, nil
	default:
		return nil, fmt.Errorf("invalid category of %q", category)
	}
}

func (node *Node) Backwards(values string) ([]Flowable, error) {
	query := "SELECT to_key_, on_script_ FROM wf_options_link WHERE diagram_id_ = ? AND from_key_ = ? ORDER BY to_key_ ASC"
	res, err := asql.Select(node.tx, query, node.diagramId, node.key)
	if err != nil {
		return nil, err
	}
	keyScrips := base.ResAsMapInt(res, "to_key_", "on_script_")

	nodes := make([]Flowable, 0, len(keyScrips))
	for key, script := range keyScrips {
		if node.Category() == enum.CategoryBranch {
			// 是否符合执行条件
			value, err := rujs.Run(node.tx, node.ctx, script, 0, flowReg(node, values))
			if err != nil {
				return nil, err
			}

			// 必须明确返回Boolean类型
			if is := value.IsBoolean(); !is {
				return nil, fmt.Errorf("%d -> %d: 执行条件返回值必须为布尔类型（%q）", node.key, key, value.String())
			}

			// 是否符合分支条件
			ok, err := value.ToBoolean()
			if err != nil {
				return nil, err
			}

			logrus.Debugf("%q ::::::::: =>=>=>=>=>=>=>=> %v", script, ok)

			if !ok {
				continue
			}
		}

		n, err := NewNode(node.tx, node.ctx, node.diagramId, key)
		if err != nil {
			return nil, err
		}

		nodes = append(nodes, n)
	}

	return nodes, nil
}

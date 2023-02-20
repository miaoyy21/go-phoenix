package xwf

import (
	"database/sql"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"go-phoenix/handle"
	"go-phoenix/xwf/enum"
	"go-phoenix/xwf/flow"
)

type Backward struct {
	Key      int           `json:"key"`      // 节点
	Name     string        `json:"name"`     // 名称
	Category enum.Category `json:"category"` // 分类
	Routes   []int         `json:"routes"`   // 节点路由

	Executors    []flow.Executor `json:"executors"`    // 默认执行者
	Organization []string        `json:"organization"` // 组织ID（部门ID或用户ID）
}

func backwards(tx *sql.Tx, ctx *handle.Context, diagramId string, key int, flowId string, values string) (interface{}, error) {
	start, err := flow.NewNode(tx, ctx, diagramId, key)
	if err != nil {
		return nil, err
	}

	route := flow.NewBackwardsRoute(start)
	nodes, err := route.Backwards(values)
	if err != nil {
		return nil, err
	}

	backs := make([]Backward, 0, len(nodes))
	exists := make(map[int]struct{})
	for _, node := range nodes {
		// 忽略条件分支
		if _, ok := node.(flow.BranchFlowable); ok {
			continue
		}

		// 路由
		routes, ok := route.Routes()[node.Key()]
		if !ok {
			logrus.Panic("unreachable")
		}

		// 当前节点，只可能是执行环节或结束节点
		backward := Backward{
			Key:          node.Key(),
			Name:         node.Name(),
			Category:     node.Category(),
			Routes:       routes,
			Executors:    make([]flow.Executor, 0),
			Organization: make([]string, 0),
		}

		// 如果包含结束节点，那么就认为流程结束
		if _, ok := node.(flow.EndFlowable); ok {
			backs = append(backs, backward)
			break
		}

		// 去掉重复的节点
		if _, ok := exists[backward.Key]; ok {
			continue
		}

		// 获取执行环节的所有执行者
		if execute, ok := node.(flow.ExecuteFlowable); ok {
			scope, err := execute.ScopeExecutors(flowId, values)
			if err != nil {
				return nil, err
			}

			backward.Organization = append(backward.Organization, scope...)

			// 默认执行者
			executors, err := execute.CustomExecutors(scope)
			if err != nil {
				return nil, err
			}

			backward.Executors = executors
		}

		exists[backward.Key] = struct{}{}
		backs = append(backs, backward)
	}

	jbs, err := json.MarshalIndent(backs, "", "\t")
	if err != nil {
		return nil, err
	}

	logrus.Debugf("backs are %s", string(jbs))

	return backs, nil
}

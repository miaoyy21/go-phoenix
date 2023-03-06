package flow

import (
	"go-phoenix/asql"
	"go-phoenix/base"
)

func (node *NodeExecute) CustomExecutors(scope []string) ([]Executor, error) {
	num := 10 // 最多默认执行者数量
	if node.executorCustomNum > 0 {
		num = node.executorCustomNum
		if node.executorSelectableNum > 0 && node.executorCustomNum > node.executorSelectableNum {
			num = node.executorSelectableNum
		}
	} else {
		if node.executorSelectableNum > 0 {
			num = node.executorSelectableNum
		}
	}

	exists := make(map[string]struct{})
	executors := make([]Executor, 0, num)

	// 是否自动保存执行者
	if node.executorSavable {
		query := `
			SELECT executor_user_id_, executor_user_name_ 
			FROM wf_flow_executors 
			WHERE diagram_id_ = ? AND key_ = ? AND create_user_id_ = ? 
			ORDER BY order_ ASC
		`
		res, err := asql.Select(node.tx, query, node.diagramId, node.key, node.ctx.UserId())
		if err != nil {
			return nil, err
		}

		// 返回
		if len(res) > 0 {
			for _, row := range res {
				executors = append(executors, Executor{
					Id:   row["executor_user_id_"],
					Name: row["executor_user_name_"],
				})
			}
		}
	}

	// 随机序列
	ns := base.Config.Rand().Perm(len(scope))

BREAK:
	for _, i := range ns {
		query := "SELECT id, user_name_ FROM sys_user WHERE id = ? OR depart_id_ = ? ORDER BY order_ ASC"
		res, err := asql.Select(node.tx, query, scope[i], scope[i])
		if err != nil {
			return nil, err
		}

		for _, row := range res {
			if _, ok := exists[row["id"]]; ok {
				continue
			}

			exists[row["id"]] = struct{}{}
			executors = append(executors, Executor{Id: row["id"], Name: row["user_name_"]})
			if len(executors) == num {
				break BREAK
			}
		}
	}

	return executors, nil
}

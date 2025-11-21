package xwf

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-phoenix/asql"
	"go-phoenix/base"
	"go-phoenix/handle"
	"go-phoenix/xwf/mdg"
	"strings"
)

type Diagrams struct {
}

func (r *Diagrams) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	id := ctx.FormValue("id")
	scope := ctx.FormValue("scope")

	query, args := "invalid", make([]interface{}, 0)
	if len(id) > 0 && strings.EqualFold(scope, "EXTRA") {
		var model, options string
		if err := asql.SelectRow(tx, "SELECT model_, options_ FROM wf_diagram WHERE id = ?", id).Scan(&model, &options); err != nil {
			return nil, err
		}

		return map[string]string{"model": model, "options": options}, nil
	} else {
		query = "SELECT id, code_, name_, icon_, description_, create_at_, update_at_, publish_at_ FROM wf_diagram ORDER BY order_ ASC"
	}

	return asql.Select(tx, query, args...)
}

func (r *Diagrams) GetPublishPermission(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	userId, departId := ctx.UserId(), ctx.DepartId()

	// 获取用户的部门ID
	org, err := asql.QueryRelationParents(tx, "sys_depart", departId)
	if err != nil {
		return nil, err
	}

	org = append(org, userId)

	// 查询用户和所有上级部门的权限
	query := fmt.Sprintf("SELECT role_id_ FROM sys_organization_role WHERE organization_id_ IN (?%s)", strings.Repeat(", ?", len(org)-1))
	userRoles, err := asql.Select(tx, query, org...)
	if err != nil {
		return nil, err
	}

	if len(userRoles) < 1 {
		return make([]map[string]string, 0), nil
	}

	useRoles := base.ResAsSliceString(userRoles, "role_id_")

	roles := make([]interface{}, 0, len(userRoles))
	for _, userRole := range useRoles {
		roles = append(roles, userRole)
	}
	query = fmt.Sprintf(`
		SELECT t1.id AS id, t1.diagram_id_, 
			t1.diagram_code_, t1.diagram_name_, 
			t1.keyword_, t1.icon_, t1.description_ 
		FROM wf_options_diagram t1
			INNER JOIN sys_diagram_role t2 ON t2.diagram_id_ = t1.diagram_id_
		WHERE t2.role_id_ IN (?%s)
		ORDER BY t1.order_ ASC
	`, strings.Repeat(", ?", len(roles)-1))

	return asql.Select(tx, query, roles...)
}

func (r *Diagrams) GetPublish(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	id := ctx.FormValue("id")
	scope := ctx.FormValue("scope")

	query, args := "invalid", make([]interface{}, 0)
	if strings.EqualFold(scope, "SIMPLE") {
		query = `
			SELECT id, diagram_id_, diagram_code_, diagram_name_, 
				keyword_, icon_, description_ 
			FROM wf_options_diagram
			ORDER BY order_ ASC
		`

		return asql.Select(tx, query, args...)
	} else if len(id) > 0 && strings.EqualFold(scope, "MODEL") {
		var model string
		query = "SELECT model_ FROM wf_options_diagram WHERE diagram_id_ = ?"
		args = append(args, id)

		if err := asql.SelectRow(tx, query, args...).Scan(&model); err != nil {
			return nil, err
		}

		return model, nil
	}

	return asql.Select(tx, query, args...)
}

func (r *Diagrams) Post(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	operation := ctx.PostFormValue("operation")
	id := ctx.PostFormValue("id")

	moveId := ctx.PostFormValue("webix_move_id")
	moveIndex := ctx.PostFormValue("webix_move_index")
	moveParent := ctx.PostFormValue("webix_move_parent")

	switch operation {
	case "delete":
		// 删除连接节点的线
		if err := asql.Delete(tx, "DELETE FROM wf_options_link WHERE diagram_id_ = ?", id); err != nil {
			return nil, err
		}

		// 删除流程节点
		if err := asql.Delete(tx, "DELETE FROM wf_options_node WHERE diagram_id_ = ?", id); err != nil {
			return nil, err
		}

		// 删除发布流程
		if err := asql.Delete(tx, "DELETE FROM wf_options_diagram WHERE diagram_id_ = ?", id); err != nil {
			return nil, err
		}

		// 删除流程
		if err := asql.Delete(tx, "DELETE FROM wf_diagram WHERE id = ?", id); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "id": id}, nil
	case "order":
		if err := asql.Order(tx, "wf_diagram", id, moveId, moveIndex, moveParent); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success"}, nil
	}

	return nil, fmt.Errorf("unrecognizable operation %s ", operation)
}

func (r *Diagrams) PostSave(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	id := ctx.PostFormValue("id")
	operation := ctx.PostFormValue("operation")
	model := ctx.PostFormValue("model")
	options := ctx.PostFormValue("options")

	var op mdg.Options
	if err := json.Unmarshal([]byte(options), &op); err != nil {
		return nil, err
	}

	now := asql.GetNow()
	code, name := strings.TrimSpace(op.Diagram.Code), strings.TrimSpace(op.Diagram.Name)
	icon, description := op.Diagram.Icon, op.Diagram.Description

	switch operation {
	case "insert":
		newId := asql.GenerateId()
		query := "INSERT INTO wf_diagram(id, code_, name_, icon_, description_, model_, options_, order_, create_at_) VALUES (?,?,?,?,?,?,?,?,?)"
		args := []interface{}{newId, code, name, icon, description, model, options, asql.GenerateOrderId(), now}
		if err := asql.Insert(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "id": newId}, nil
	case "update":
		query := "UPDATE wf_diagram SET code_ = ?, name_ = ?, icon_ = ?, description_ = ?, model_ = ?, options_ = ?, update_at_ = ? WHERE id = ?"
		args := []interface{}{code, name, icon, description, model, options, now, id}
		if err := asql.Update(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "id": id}, nil
	}

	return nil, fmt.Errorf("unrecognizable operation %s ", operation)
}

func (r *Diagrams) PostPublish(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	id := ctx.PostFormValue("id")

	var code, name, model, options string
	var order int64
	err := asql.SelectRow(tx, "SELECT code_, name_, model_, options_, order_ FROM wf_diagram WHERE id = ?", id).Scan(&code, &name, &model, &options, &order)
	if err != nil {
		return nil, err
	}

	// 流程编码
	if len(code) < 1 {
		return nil, errors.New("请输入流程编码")
	}

	// 流程名称
	if len(name) < 1 {
		return nil, errors.New("请输入流程名称")
	}

	var mo mdg.Model
	if err := json.Unmarshal([]byte(model), &mo); err != nil {
		logrus.Error(err)
		return nil, err
	}

	var op mdg.Options
	if err := json.Unmarshal([]byte(options), &op); err != nil {
		return nil, err
	}

	// 构建流程图
	if err := mdg.Publish(tx, id, code, name, model, order, mo, op); err != nil {
		return nil, err
	}

	// 记录发布时间
	now := asql.GetNow()
	if err := asql.Update(tx, "UPDATE wf_diagram SET publish_at_ = ? WHERE id = ?", now, id); err != nil {
		return nil, err
	}

	return map[string]interface{}{"status": "success", "id": id, "publish_at_": now}, nil
}

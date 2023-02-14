package xwf

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-phoenix/asql"
	"go-phoenix/handle"
	"go-phoenix/xwf/mdg"
	"strings"
)

type Diagrams struct {
}

func (r *Diagrams) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	id := ctx.FormValue("id")
	scope := ctx.FormValue("scope")

	if len(id) > 0 && strings.EqualFold(scope, "EXTRA") {
		var model, options string

		query := "SELECT model_, options_ FROM wf_diagram WHERE id = ?"
		if err := asql.SelectRow(tx, query, id).Scan(&model, &options); err != nil {
			return nil, err
		}

		return map[string]string{"model": model, "options": options}, nil
	}

	query := "SELECT id, code_, name_, description_, create_at_, update_at_, publish_at_ FROM wf_diagram ORDER BY order_ ASC"
	return asql.Select(tx, query)
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
	code, name, description := strings.TrimSpace(op.Diagram.Code), strings.TrimSpace(op.Diagram.Name), op.Diagram.Description

	switch operation {
	case "insert":
		newId := asql.GenerateId()
		query := "INSERT INTO wf_diagram(id, code_, name_, description_, model_, options_, order_, create_at_) VALUES (?,?,?,?,?,?,?,?)"
		args := []interface{}{newId, code, name, description, model, options, asql.GenerateOrderId(), now}
		if err := asql.Insert(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "id": newId}, nil
	case "update":
		query := "UPDATE wf_diagram SET code_ = ?, name_ = ?, description_ = ?, model_ = ?, options_ = ?, update_at_ = ? WHERE id = ?"
		args := []interface{}{code, name, description, model, options, now, id}
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
	err := asql.SelectRow(tx, "SELECT code_, name_, model_, options_ FROM wf_diagram WHERE id = ?", id).Scan(&code, &name, &model, &options)
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
	if err := mdg.Publish(tx, id, code, name, mo, op); err != nil {
		return nil, err
	}

	// 记录发布时间
	now := asql.GetNow()
	if err := asql.Update(tx, "UPDATE wf_diagram SET publish_at_ = ? WHERE id = ?", now, id); err != nil {
		return nil, err
	}

	return map[string]interface{}{"status": "success", "id": id, "publish_at_": now}, nil
}

package xwf

import (
	"database/sql"
	"encoding/json"
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

	query, args := "invalid", make([]interface{}, 0)
	if len(id) > 0 && strings.EqualFold(scope, "EXTRA") {
		query = "SELECT model_, options_ FROM wf_diagram WHERE id = ?"
		args = append(args, id)
	} else {
		query = `
			SELECT id, code_, name_, description_, create_at_, update_at_
			FROM wf_diagram
			ORDER BY order_ ASC
		`
	}

	res, err := asql.Select(tx, query, args...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *Diagrams) Post(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	operation := ctx.PostFormValue("operation")

	id := ctx.PostFormValue("id")
	code := ctx.PostFormValue("code_")
	name := ctx.PostFormValue("name_")
	description := ctx.PostFormValue("description_")

	moveId := ctx.PostFormValue("webix_move_id")
	moveIndex := ctx.PostFormValue("webix_move_index")
	moveParent := ctx.PostFormValue("webix_move_parent")

	now := asql.GetNow()
	switch operation {
	case "insert":
		newId := asql.GenerateId()
		query := "INSERT INTO wf_diagram(id, code_, name_, description_, order_, create_at_) VALUES (?,?,?,?,?,?,?)"
		args := []interface{}{newId, code, name, description, asql.GenerateOrderId(), now}
		if err := asql.Insert(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "newid": newId, "create_at_": now}, nil
	case "update":
		query := "UPDATE wf_diagram SET code_ = ?, name_ = ?, description_ = ?, update_at_ = ? WHERE id = ?"
		args := []interface{}{code, name, description, now, id}
		if err := asql.Update(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "id": id, "update_at_": now}, nil
	case "delete":
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
	model := ctx.PostFormValue("model")
	options := ctx.PostFormValue("options")

	now := asql.GetNow()
	query := "UPDATE wf_diagram SET model_ = ?, options_ = ?, update_at_ = ? WHERE id = ?"
	args := []interface{}{model, options, now, id}
	if err := asql.Update(tx, query, args...); err != nil {
		return nil, err
	}

	return map[string]interface{}{"status": "success", "id": id, "update_at_": now}, nil
}

func (r *Diagrams) PostPublish(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	id := ctx.PostFormValue("id")

	var code, name, model, options string
	err := asql.SelectRow(tx, "SELECT code_, name_, model_, options_ FROM wf_diagram WHERE id = ?", id).Scan(&code, &name, &model, &options)
	if err != nil {
		return nil, err
	}

	var mo mdg.Model
	if err := json.Unmarshal([]byte(model), &mo); err != nil {
		logrus.Error(err)
		return nil, err
	}

	var op mdg.Options
	if err := json.Unmarshal([]byte(options), &op); err != nil {
		logrus.Error(err)
		return nil, err
	}

	// 构建流程图
	if err := mdg.Publish(tx, id, code, name, mo, op); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return map[string]interface{}{"status": "success", "id": id}, nil
}

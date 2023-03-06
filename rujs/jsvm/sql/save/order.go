package vm

import (
	"database/sql"
	"errors"
	"github.com/sirupsen/logrus"
	"go-phoenix/asql"
	"go-phoenix/handle"
)

// Order 排序操作（使用排序的依据必须是 order_ ASC）
func Order(tx *sql.Tx, ctx *handle.Context, table string) map[string]interface{} {
	values := ctx.Values()

	// 起始行
	source, ok := values["id"]
	if !ok {
		logrus.Panic(errors.New("提交参数必须包含id字段"))
	}

	// 目标行
	target, ok := values["webix_move_id"]
	if !ok {
		logrus.Panic(errors.New("提交参数必须包含webix_move_id字段"))
	}

	// 目标的索引
	targetIndex := values["webix_move_index"]
	if !ok {
		logrus.Panic(errors.New("提交参数必须包含webix_move_index字段"))
	}

	// 目标的上级
	targetParent := values["webix_move_parent"]
	if !ok {
		logrus.Panic(errors.New("提交参数必须包含webix_move_parent字段"))
	}

	if err := asql.Order(tx, table, source, target, targetIndex, targetParent); err != nil {
		logrus.Panic(err)
	}

	return map[string]interface{}{"status": "success", "id": source}
}

package vm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-phoenix/asql"
	"go-phoenix/handle"
	"strings"
)

// Delete 删除操作
func Delete(tx *sql.Tx, ctx *handle.Context, table string) map[string]interface{} {
	values := ctx.GetValues()

	// id
	id, ok := values["id"]
	if !ok {
		logrus.Panic(errors.New("提交参数必须包含id字段"))
	}

	ids := make([]interface{}, 0)

	// 无法解析为JSON，作为字符串处理
	if err := json.Unmarshal([]byte(id), &ids); err != nil {
		ids = append(ids, id)
	}

	// 返回值
	ret := make(map[string]interface{})

	// 除了删除自身，也要删除其子节点
	news, err := asql.QueryChildren(tx, table, ids)
	if err != nil {
		logrus.Panic(err)
	}

	// 级联删除
	if err := asql.Delete(tx, fmt.Sprintf("DELETE FROM %s WHERE id IN (?%s)", table, strings.Repeat(", ?", len(news)-1)), news...); err != nil {
		logrus.Panic(err)
	}

	ret["id"] = news
	ret["status"] = "success"
	return ret
}

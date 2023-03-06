package vm

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-phoenix/asql"
	"go-phoenix/base"
	"go-phoenix/handle"
	"strings"
)

// Update 更新操作
func Update(tx *sql.Tx, ctx *handle.Context, table string) map[string]interface{} {
	values := ctx.Values()

	query := `
		SELECT sys_table_column.code_
		FROM sys_table
			LEFT JOIN sys_table_column ON sys_table.id = sys_table_column.table_id_
		WHERE sys_table.code_ = ?
	`
	args := []interface{}{table}
	res, err := asql.Select(tx, query, args...)
	if err != nil {
		logrus.Panic(err)
	}
	cols := base.ResAsMap2(res, "code_")

	fields := make(map[string]interface{})
	for key, value := range values {
		if asql.IsSaveIgnore(key) {
			continue
		}

		if strings.EqualFold(key, "parent") {
			if strings.EqualFold(value, "0") {
				fields["parent_id_"] = nil
			} else {
				fields["parent_id_"] = value
			}

			continue
		}

		// 对零值的特殊处理
		if _, ok := cols[key]; ok {
			if len(value) < 1 {
				fields[key] = nil
				continue
			}

			fields[key] = value
		}
	}

	// id
	id, ok := values["id"]
	if !ok {
		logrus.Panic(errors.New("提交参数必须包含id字段"))
	}

	// 返回值
	ret := make(map[string]interface{})

	// update_at_
	if _, ok := cols["update_at_"]; ok {
		now := asql.GetNow()

		fields["update_at_"] = now
		ret["update_at_"] = now
	}

	ks := make([]string, 0, len(fields))
	vs := make([]interface{}, 0, len(fields))
	for k, v := range fields {
		if v == nil {
			ks = append(ks, fmt.Sprintf("%s = NULL", k))
		} else {
			ks = append(ks, fmt.Sprintf("%s = ?", k))
			vs = append(vs, v)
		}
	}
	vs = append(vs, id)

	query = fmt.Sprintf("UPDATE %s SET %s WHERE id = ?", table, strings.Join(ks, ","))
	if err := asql.Update(tx, query, vs...); err != nil {
		logrus.Panic(err)
	}

	ret["id"] = id
	ret["status"] = "success"
	return ret
}

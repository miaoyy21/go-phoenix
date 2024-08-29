package vm

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-phoenix/asql"
	"go-phoenix/base"
	"go-phoenix/handle"
	"strconv"
	"strings"
)

// Insert 插入操作
func Insert(tx *sql.Tx, ctx *handle.Context, table string, values map[string]string) map[string]interface{} {
	// 比较原始提交数据与待更新数据的差异，用于返回客户端
	changed := base.CompareMapChanged(base.GetURLValues(ctx.PostForm), values)

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

	fields := make(map[string]string)
	for key, value := range values {
		if strings.EqualFold(key, "parent") {
			fields["parent_id_"] = value
			continue
		}

		if _, ok := cols[key]; ok {
			fields[key] = value
		}
	}

	// 返回值
	ret := make(map[string]interface{})
	for col := range cols {
		isRet := true

		switch col {
		case "id":
			newid := asql.GenerateId()
			fields[col] = newid

			isRet = false
			ret[col] = values[col] // 返回客户端发送的id
			ret["newid"] = newid   // 服务器生成的newid
		case "order_":
			fields[col] = strconv.FormatInt(asql.GenerateOrderId(), 10)
		case "create_depart_id_":
			fields[col] = ctx.DepartId()
		case "create_depart_code_":
			fields[col] = ctx.DepartCode()
		case "create_depart_name_":
			fields[col] = ctx.DepartName()
		case "create_user_id_":
			fields[col] = ctx.UserId()
		case "create_user_code_":
			fields[col] = ctx.UserCode()
		case "create_user_name_":
			fields[col] = ctx.UserName()
		case "create_at_":
			fields[col] = asql.GetNow()
		default:
			if _, ok := changed[col]; !ok {
				isRet = false
			}
		}

		if isRet {
			ret[col] = fields[col]
		}
	}

	ks := make([]string, 0, len(fields))
	ps := make([]string, 0, len(fields))
	vs := make([]interface{}, 0, len(fields))
	for k, v := range fields {
		ks = append(ks, k)
		ps = append(ps, "?")

		if strings.EqualFold(k, "parent_id_") {
			if strings.EqualFold(v, "0") {
				vs = append(vs, nil)
				continue
			} else {
				ret["parent_id_"] = v
			}
		}

		vs = append(vs, v)
	}

	query = fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s)", table, strings.Join(ks, ","), strings.Join(ps, ","))
	if err := asql.Insert(tx, query, vs...); err != nil {
		logrus.Panic(err)
	}

	ret["status"] = "success"
	return ret
}

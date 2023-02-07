package vm

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-phoenix/asql"
	"go-phoenix/base"
	"strconv"
	"strings"
)

// Select 查询数据
func (s *Sql) Select(args ...string) map[string]interface{} {
	params := s.ctx.GetParams()

	var fields, from, where, groupBy, orderBy string
	for _, arg := range args {
		arg := strings.ToLower(strings.TrimSpace(arg))
		if strings.HasPrefix(arg, "select ") {
			fields = strings.TrimSpace(strings.TrimPrefix(arg, "select "))
		} else if strings.HasPrefix(arg, "from ") {
			from = strings.TrimSpace(strings.TrimPrefix(arg, "from "))
		} else if strings.HasPrefix(arg, "where ") {
			where = strings.TrimSpace(strings.TrimPrefix(arg, "where "))
		} else if strings.HasPrefix(arg, "group ") {
			groupBy = strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(strings.TrimPrefix(arg, "group ")), "by "))
		} else if strings.HasPrefix(arg, "order ") {
			orderBy = strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(strings.TrimPrefix(arg, "order ")), "by "))
		} else {
			logrus.Panic(fmt.Errorf("无法判定的子语句 %q", arg))
		}
	}

	// 至少包含SELECT
	if len(fields) < 1 {
		logrus.Panic(errors.New("至少包含SELECT子语句"))
	}

	// 只能简单分割,仅识别 T.a AS s || T.a
	mapFields := make(map[string]string)
	sFields := strings.Split(fields, ",")
	for _, field := range sFields {
		ls := strings.Split(strings.TrimSpace(field), " ")
		if len(ls) == 2 {
			ls1, ls0 := strings.TrimSpace(ls[1]), strings.TrimSpace(ls[0])
			if len(ls1) < 1 || len(ls0) < 1 {
				continue
			}

			mapFields[ls1] = ls0
		} else {
			ls := strings.Split(strings.TrimSpace(field), " as ")
			if len(ls) == 2 {
				ls1, ls0 := strings.TrimSpace(ls[1]), strings.TrimSpace(ls[0])
				if len(ls1) < 1 || len(ls0) < 1 {
					continue
				}

				mapFields[ls1] = ls0
			} else {
				// TODO SELECT中存在子查询
			}
		}
	}

	// 获取排序和过滤条件
	sorts, filters := s.ctx.GetSortsFilters(mapFields)

	// 获取参数，排除filter和sort
	newParams := make(map[string]string)
	ignore := map[string]struct{}{"service": {}, "is_parse_columns": {}, "start": {}, "count": {}, "continue": {}, "pager": {}} // 忽略的参数
	for key, value := range params {
		if _, ok := ignore[key]; ok {
			continue
		}

		if strings.HasPrefix(key, "sort[") && strings.HasSuffix(key, "]") {
			continue
		}

		if strings.HasPrefix(key, "filter[") && strings.HasSuffix(key, "]") {
			continue
		}

		newParams[key] = value
	}

	/************************** WHERE **************************/
	wheres := make([]string, 0, len(newParams)+len(filters)+1)

	// 自身的where条件
	if len(where) > 0 {
		wheres = append(wheres, fmt.Sprintf("(%s)", where))
	}

	// 传递的参数
	for key, value := range newParams {
		if newKey, ok := mapFields[key]; ok {
			wheres = append(wheres, fmt.Sprintf("%s = '%s'", newKey, value))
		} else {
			wheres = append(wheres, fmt.Sprintf("%s = '%s'", key, value))
		}
	}

	// 过滤条件
	if len(filters) > 0 {
		wheres = append(wheres, strings.Join(filters, " AND "))
	}

	/************************** ORDER BY **************************/
	if len(sorts) > 0 {
		orderBy = strings.Join(sorts, ",")
	}

	if len(orderBy) < 1 {
		orderBy = "order_ ASC"
	}

	sqs := make([]string, 0, 3)

	// FROM
	sqs = append(sqs, fmt.Sprintf("FROM %s", from))

	// WHERE
	if len(wheres) > 0 {
		sqs = append(sqs, fmt.Sprintf("WHERE %s", strings.Join(wheres, " AND ")))
	}

	// GROUP BY HAVING
	if len(groupBy) > 0 {
		sqs = append(sqs, fmt.Sprintf("GROUP BY %s", groupBy))
	}

	// 是否为仅获取数据列
	if _, ok := params["is_parse_columns"]; ok {
		query := fmt.Sprintf("\n\tSELECT %s \n\t%s \n\tLIMIT %d,%d \n", fields, strings.Join(sqs, "\n\t"), 0, 1)
		columns, err := s.ParseColumnConfigs(query)
		if err != nil {
			logrus.Panic(err)
		}

		return map[string]interface{}{"columns": columns}
	}

	if _, ok := params["pager"]; ok {
		sStart, sCount := params["start"], params["count"]

		// 有start和count参数时认为属于分页
		start, _ := strconv.Atoi(sStart)
		count, _ := strconv.Atoi(sCount)

		count = count * 5
		if count < 1 {
			count = 250
		}

		// 查询总记录数量
		var totalCount int
		if err := asql.SelectRow(s.tx, fmt.Sprintf("SELECT COUNT(1) %s", strings.Join(sqs, " "))).Scan(&totalCount); err != nil {
			logrus.Panic(err)
		}

		// 根据分页查询当前页
		query := fmt.Sprintf("\n\tSELECT %s \n\t%s \n\tORDER BY %s \n\tLIMIT %d,%d \n", fields, strings.Join(sqs, "\n\t"), orderBy, start, count)
		data, err := asql.Select(s.tx, query)
		if err != nil {
			logrus.Panic(err)
		}

		return map[string]interface{}{
			"data":        data,
			"pos":         start,
			"total_count": totalCount,
		}
	}

	// 不分页
	query := fmt.Sprintf("\n\tSELECT %s \n\t%s \n\tORDER BY %s \n", fields, strings.Join(sqs, "\n\t"), orderBy)
	data, err := asql.Select(s.tx, query)
	if err != nil {
		logrus.Panic(err)
	}

	return map[string]interface{}{
		"data": data,
	}
}

func (s *Sql) ParseColumnConfigs(query string) ([]map[string]string, error) {
	cols, err := asql.SelectColumns(s.tx, query)
	if err != nil {
		return nil, err
	}

	// 检查是否存在重复列
	zs, rs := make(map[string]int), make([]string, 0)
	for _, col := range cols {
		zs[strings.ToLower(col)]++

		if zs[col] == 2 {
			rs = append(rs, col)
		}
	}

	// 不允许列重复
	if len(rs) > 0 {
		return nil, fmt.Errorf("存在重复的查询列【%s】", strings.Join(rs, ","))
	}

	// 必须包含id字段
	if _, ok := zs["id"]; !ok {
		return nil, errors.New("缺少主键id")
	}

	// 查询所有列的类型
	qs := `
		SELECT t.code_, t.name_, t.type_
		FROM sys_table_column t 
			INNER JOIN (
				SELECT code_,MAX(id) as max_id
				FROM sys_table_column 
				WHERE code_ IN (?%s)  
				GROUP BY code_
			) m ON t.id = m.max_id
	`

	args := make([]interface{}, 0, len(cols))
	for _, arg := range cols {
		args = append(args, arg)
	}

	res, err := asql.Select(s.tx, fmt.Sprintf(qs, strings.Repeat(", ?", len(args)-1)), args...)
	if err != nil {
		return nil, err
	}

	configs := make([]map[string]string, 0, len(cols)+1)

	// 添加 索引
	configs = append(configs, map[string]string{"val.id": "index", "val.hidden": "true", "header.text": "№", "header.align": "center", "val.align": "center", "val.width": "60"})

	cts := base.ResAsMap2(res, "code_")
	for _, col := range cols {
		config := make(map[string]string)

		// 是否默认忽略列
		config["val.id"] = col
		config["val.hidden"] = "false"
		if asql.IsSaveIgnore(col) {
			config["val.hidden"] = "true"
		}

		// 默认页眉
		config["header.align"] = "center"
		config["header.filter"] = "disable"

		// 默认编辑器
		config["val.editor"] = "text"

		// 默认排序
		config["val.sort"] = "disable"

		// 默认单元格对齐
		config["val.align"] = "left"
		config["val.resize"] = "false"
		config["val.adjust"] = "true"

		ms, ok := cts[col]
		if ok {
			n, t := ms["name_"], ms["type_"]

			// 默认页眉
			config["header.text"] = n
			switch t {
			case "VARCHAR(256)":
			case "VARCHAR(1024)":
				config["val.editor"] = "popup"

				delete(config, "val.adjust")
				config["val.width"] = "360"
				config["val.minWidth"] = "360"
				config["val.maxWidth"] = "540"
			case "VARCHAR(4096)":
				config["val.editor"] = "popup"

				delete(config, "val.adjust")
				config["val.width"] = "480"
				config["val.minWidth"] = "480"
				config["val.maxWidth"] = "720"
			case "TINYINT":
				config["val.align"] = "right"
				config["val.format"] = "int"
			case "INT":
				config["val.align"] = "right"
				config["val.format"] = "int"
			case "BIGINT":
				config["val.align"] = "right"
				config["val.format"] = "int"
			case "NUMERIC(13,2)":
				config["val.align"] = "right"

				// 语义断定是否代表金钱
				symbol := strings.Split("额,财,货,资,赈,贺,贡,贷,赂,赠,赉,赏,赐,赢,贮,贸,赎,费,贵,金,银,钱,贪", ",")
				isPrice := false
				for _, s := range symbol {
					if strings.Contains(n, s) {
						isPrice = true
						break
					}
				}

				if isPrice {
					config["val.format"] = "price"
				} else {
					config["val.format"] = "number"
				}

				config["val.decimal.size"] = "2"
			case "NUMERIC(18,4)":
				config["val.align"] = "right"
				config["val.format"] = "number"
				config["val.decimal.size"] = "4"
			case "NUMERIC(23,6)":
				config["val.align"] = "right"
				config["val.format"] = "number"
				config["val.decimal.size"] = "6"
			case "DATE":
				config["val.align"] = "center"
				config["val.editor"] = "date"
				config["val.format"] = "date"
			case "DATETIME":
				config["val.align"] = "center"
				config["val.editor"] = "date"
				config["val.format"] = "datetime"
			case "TEXT":
				config["val.editor"] = "popup"

				delete(config, "val.adjust")
				config["val.width"] = "480"
				config["val.minWidth"] = "480"
				config["val.maxWidth"] = "720"
			default:
				// Nothing
			}

			// 忽略字段移除编辑器
			if asql.IsSaveIgnore(col) {
				config["val.editor"] = "disable"
			}
		} else {
			config["header.text"] = strings.Title(strings.TrimSpace(strings.ReplaceAll(col, "_", " ")))
		}

		configs = append(configs, config)
	}

	// 添加 复制按钮
	configs = append(configs, map[string]string{"val.id": "button_copy", "val.hidden": "true", "header.text": "复制按钮"})

	// 添加 删除按钮
	configs = append(configs, map[string]string{"val.id": "button_delete", "val.hidden": "true", "header.text": "删除按钮"})

	return configs, nil
}

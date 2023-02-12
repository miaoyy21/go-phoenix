package xsys

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/base"
	"go-phoenix/handle"
	"strings"
	"time"
)

type SysTables struct {
}

func (o *SysTables) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	query := "SELECT id, code_, name_, sync_status_, description_, create_at_, update_at_ FROM sys_table ORDER BY order_ ASC"
	res, err := asql.Select(tx, query)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (o *SysTables) Post(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	operation := ctx.PostFormValue("operation")

	id := ctx.PostFormValue("id")
	code := ctx.PostFormValue("code_")
	name := ctx.PostFormValue("name_")
	syncStatus := ctx.PostFormValue("sync_status_")
	description := ctx.PostFormValue("description_")

	moveId := ctx.PostFormValue("webix_move_id")
	moveIndex := ctx.PostFormValue("webix_move_index")
	moveParent := ctx.PostFormValue("webix_move_parent")

	now := asql.GetNow()
	switch operation {
	case "insert":
		tableId := asql.GenerateId()
		creating := "Creating"

		// 数据库表
		query := "INSERT INTO sys_table(id, code_, name_, sync_status_, description_, order_, create_at_) VALUES (?,?,?,?,?,?,?)"
		args := []interface{}{tableId, code, name, creating, description, asql.GenerateOrderId(), now}
		if err := asql.Insert(tx, query, args...); err != nil {
			return nil, err
		}

		// 数据库表的默认列
		for _, col := range asql.SysColumns {
			order := asql.GenerateOrderId()
			if !strings.EqualFold(col.Code, "id") {
				order = order + (time.Hour * 24 * 365).Nanoseconds()
			}

			query := "INSERT INTO sys_table_column(id, table_id_, is_sys_, code_, name_, type_, description_, order_, create_at_) VALUES (?,?,?,?,?,?,?,?,?)"
			args := []interface{}{asql.GenerateId(), tableId, true, col.Code, col.Name, col.Type, col.Description, order, now}
			if err := asql.Insert(tx, query, args...); err != nil {
				return nil, err
			}
		}

		return map[string]interface{}{"status": "success", "newid": tableId, "sync_status_": creating, "create_at_": now}, nil
	case "update":
		query := "UPDATE sys_table SET code_ = ?, name_ = ?, sync_status_ = ?, description_ = ?, update_at_ = ? WHERE id = ?"
		args := []interface{}{code, name, syncStatus, description, now, id}
		if err := asql.Update(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "id": id, "update_at_": now}, nil
	case "delete":
		var table string

		// 表名
		if err := asql.SelectRow(tx, "SELECT code_ FROM sys_table WHERE id = ?", id).Scan(&table); err != nil {
			if err == sql.ErrNoRows {
				return map[string]interface{}{"status": "success"}, nil
			}

			return nil, err
		}

		// 数据库表
		if err := asql.Delete(tx, "DELETE FROM sys_table WHERE id = ?", id); err != nil {
			return nil, err
		}

		// 数据库表的所有列
		if err := asql.Delete(tx, "DELETE FROM sys_table_column WHERE table_id_ = ?", id); err != nil {
			return nil, err
		}

		// 把数据库表名改掉（不实际删除）
		if exists := o.exists(tx, table); exists {
			query := fmt.Sprintf("ALTER TABLE %s RENAME TO  %s ;", table, fmt.Sprintf("_%s", table))
			if _, err := asql.Exec(tx, query); err != nil {
				return nil, err
			}
		}

		return map[string]interface{}{"status": "success"}, nil
	case "order":
		if err := asql.Order(tx, "sys_table", id, moveId, moveIndex, moveParent); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success"}, nil
	}

	return nil, fmt.Errorf("unrecognizable operation %s ", operation)
}

func (o *SysTables) exists(tx *sql.Tx, table string) bool {
	// TODO 暂时以能不能正常执行作为判定是否创建表
	if _, err := asql.Exec(tx, fmt.Sprintf("DESC %s", table)); err != nil {
		return false
	}

	return true
}

func (o *SysTables) PostSync(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	id := ctx.PostFormValue("id")

	// 查询表名
	var table string
	if err := asql.SelectRow(tx, "SELECT code_ FROM sys_table WHERE id = ?", id).Scan(&table); err != nil {
		return nil, err
	}

	// 表是否存在
	exists := o.exists(tx, table)

	// 查询所有列
	res, err := asql.Select(tx, "SELECT code_, type_ FROM sys_table_column WHERE table_id_ = ? ORDER BY order_ ASC", id)
	if err != nil {
		return nil, err
	}

	present, cols := base.ResAsMapSlice(res, false, "code_", "type_")
	if len(cols) < 1 {
		return nil, errors.New("empty table columns")
	}

	buf := new(bytes.Buffer)
	if !exists {
		// 创建数据库表
		buf.WriteString(fmt.Sprintf("\n CREATE TABLE %s ( \n", table))

		for _, col := range cols {
			if strings.EqualFold(col, "id") {
				buf.WriteString(fmt.Sprintf("\t %s %s COLLATE utf8mb4_general_ci NOT NULL, \n", col, present[col]))
			} else {
				buf.WriteString(fmt.Sprintf("\t %s %s COLLATE utf8mb4_general_ci DEFAULT NULL, \n", col, present[col]))
			}
		}

		buf.WriteString("\t PRIMARY KEY (`id`) \n")
		buf.WriteString(" ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci")

		if _, err := asql.Exec(tx, buf.String()); err != nil {
			return nil, err
		}
	} else {
		res, err := asql.Select(tx, fmt.Sprintf("DESC %s", table))
		if err != nil {
			return nil, err
		}

		// 将Decimal转为Numeric再比较
		sLatest, sCols := base.ResAsMapSlice(res, false, "Field", "Type")

		latest := make(map[string]string)
		xCols := make([]string, 0, len(sCols))
		for _, sCol := range sCols {
			if strings.HasPrefix(sCol, "_") {
				continue
			}

			latest[sCol] = strings.ReplaceAll(strings.ToLower(sLatest[sCol]), "decimal", "numeric")
			xCols = append(xCols, sCol)
		}

		added, changed, removed := base.CompareMap(latest, present)
		if len(added)+len(changed)+len(removed) < 1 {
			if len(cols) != len(xCols) {
				return nil, fmt.Errorf("compare map is empty, but there size is not equal %d <> %d", len(cols), len(xCols))
			}

			// 是否顺序变化
			for i, col := range cols {
				if !strings.EqualFold(col, xCols[i]) {
					changed[col] = present[col]
				}
			}

			// 没有变化？
			if len(changed) < 0 {
				return map[string]interface{}{"status": "success"}, nil
			}
		}

		syntax := make([]string, 0, len(added)+len(changed)+len(removed))
		buf.WriteString(fmt.Sprintf("\n ALTER TABLE %s ", table))

		var lastCol string
		for _, col := range cols {
			// 添加
			if _, ok := added[col]; ok {
				if len(lastCol) < 1 {
					syntax = append(syntax, fmt.Sprintf("\n\t ADD COLUMN %s %s COLLATE utf8mb4_general_ci DEFAULT NULL FIRST", col, present[col]))
				} else {
					syntax = append(syntax, fmt.Sprintf("\n\t ADD COLUMN %s %s COLLATE utf8mb4_general_ci DEFAULT NULL AFTER %s", col, present[col], lastCol))
				}
			}

			// 修改
			if _, ok := changed[col]; ok {
				if len(lastCol) < 1 {
					syntax = append(syntax, fmt.Sprintf("\n\t CHANGE COLUMN %s %s %s COLLATE utf8mb4_general_ci DEFAULT NULL FIRST", col, col, present[col]))
				} else {
					syntax = append(syntax, fmt.Sprintf("\n\t CHANGE COLUMN %s %s %s COLLATE utf8mb4_general_ci DEFAULT NULL AFTER %s", col, col, present[col], lastCol))
				}
			}

			lastCol = col
		}

		// 删除（不实际删除列），并且移动至表最后
		for field, xType := range removed {
			syntax = append(syntax, fmt.Sprintf("\n\t CHANGE COLUMN %s %s %s COLLATE utf8mb4_general_ci DEFAULT NULL AFTER %s", field, fmt.Sprintf("_%s", field), xType, lastCol))
		}

		buf.WriteString(strings.Join(syntax, ","))
		buf.WriteByte(';')

		if _, err := asql.Exec(tx, buf.String()); err != nil {
			return nil, err
		}
	}

	done := "Done"
	if err := asql.Update(tx, "UPDATE sys_table SET sync_status_ = ? WHERE id = ?", done, id); err != nil {
		return nil, err
	}

	return map[string]interface{}{"status": "success", "sync_status_": done}, nil
}

package xsys

import (
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

	return asql.Select(tx, query)
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

		// 数据服务
		if err := asql.Delete(tx, "DELETE FROM sys_data_service WHERE table_id_ = ?", id); err != nil {
			return nil, err
		}

		// 数据库表的所有列
		if err := asql.Delete(tx, "DELETE FROM sys_table_column WHERE table_id_ = ?", id); err != nil {
			return nil, err
		}

		// 数据库表
		if err := asql.Delete(tx, "DELETE FROM sys_table WHERE id = ?", id); err != nil {
			return nil, err
		}

		// 把数据库表名改掉（不实际删除）
		ddl := ctx.DDL(tx, table, nil, nil) // 数据定义语法
		if ddl.Exists() {
			if err := ddl.Drop(); err != nil {
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

func (o *SysTables) PostSync(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	id := ctx.PostFormValue("id")

	// 查询表名
	var table string
	if err := asql.SelectRow(tx, "SELECT code_ FROM sys_table WHERE id = ?", id).Scan(&table); err != nil {
		return nil, err
	}

	// 查询所有列
	res, err := asql.Select(tx, "SELECT code_, type_ FROM sys_table_column WHERE table_id_ = ? ORDER BY order_ ASC", id)
	if err != nil {
		return nil, err
	}

	present, cols := base.ResAsMapSlice(res, false, "code_", "type_")
	if len(cols) < 1 {
		return nil, errors.New("empty table columns")
	}

	ddl := ctx.DDL(tx, table, cols, present) // 数据定义语法
	if !ddl.Exists() {

		// 执行创建数据库表
		if err := ddl.Create(); err != nil {
			return nil, err
		}

		// 自动生成数据查询服务
		qQuery := "INSERT INTO sys_data_service(id, table_id_, code_, name_, method_, timeout_, source_, order_, create_at_) VALUES (?,?,?,?,?,?,?,?,?)"
		qArgs := []interface{}{asql.GenerateId(), id, "query", "查询服务(检索)", "GET", 0, fmt.Sprintf("/* 查询实现：按字段查询(=)、排序、分页、过滤(LIKE) 等常见需求 */\nsql.Select(\"SELECT * \", \"FROM %s\");", table), asql.GenerateOrderId(), asql.GetNow()}
		if err := asql.Insert(tx, qQuery, qArgs...); err != nil {
			return nil, err
		}

		// 自动生成数据查询服务
		sQuery := "INSERT INTO sys_data_service(id, table_id_, code_, name_, method_, timeout_, source_, order_, create_at_) VALUES (?,?,?,?,?,?,?,?,?)"
		sArgs := []interface{}{asql.GenerateId(), id, "save", "保存服务(新增 修改 删除 排序)", "POST", 0, fmt.Sprintf("/* 保存实现：插入、更新、删除、排序 等操作*/\nsql.Save(%q);", table), asql.GenerateOrderId(), asql.GetNow()}
		if err := asql.Insert(tx, sQuery, sArgs...); err != nil {
			return nil, err
		}
	} else {
		sCols, sLatest, err := ddl.Desc()
		if err != nil {
			return nil, err
		}

		latest := make(map[string]string)
		xCols := make([]string, 0, len(sCols))
		for _, sCol := range sCols {
			if strings.HasPrefix(sCol, "_") {
				continue
			}

			// 将Decimal转为Numeric再比较
			latest[sCol] = strings.ReplaceAll(strings.ToLower(sLatest[sCol]), "decimal", "numeric")
			xCols = append(xCols, sCol)
		}

		added, changed, removed := base.CompareMap(latest, present)
		if len(added)+len(changed)+len(removed) < 1 {
			if len(cols) != len(xCols) {
				return nil, fmt.Errorf("compare map is empty, but there size is not equal %d <> %d", len(cols), len(xCols))
			}

			// 是否顺序变化，达梦不支持修改列顺序
			if ddl.IsSupportSequence() {
				for i, col := range cols {
					if !strings.EqualFold(col, xCols[i]) {
						changed[col] = present[col]
					}
				}
			}

			// 没有变化？
			if len(changed) <= 0 {
				return map[string]interface{}{"status": "success"}, nil
			}
		}

		// 修改数据库表
		if err := ddl.Alter(added, changed, removed); err != nil {
			return nil, err
		}
	}

	done := "Done"
	if err := asql.Update(tx, "UPDATE sys_table SET sync_status_ = ? WHERE id = ?", done, id); err != nil {
		return nil, err
	}

	return map[string]interface{}{"status": "success", "sync_status_": done}, nil
}

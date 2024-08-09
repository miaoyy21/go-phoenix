package asql

import (
	"bytes"
	"database/sql"
	"fmt"
	"go-phoenix/base"
	"strings"
)

type DDL interface {
	Exists() bool            // 判断表是否存在
	IsSupportSequence() bool // 是否支持修改列顺序

	Desc() ([]string, map[string]string, error)                          // 获取数据库表结构描述信息
	Create() error                                                       // 创建数据库表结构
	Alter(map[string]string, map[string]string, map[string]string) error // 修改数据库表结构

	Drop() error // 删除数据库表结构的SQL语句
}

type DDLBase struct {
	tx *sql.Tx

	table   string
	cols    []string
	present map[string]string
}

func NewDDLBase(tx *sql.Tx, table string, cols []string, present map[string]string) *DDLBase {
	return &DDLBase{
		tx: tx,

		table:   table,
		cols:    cols,
		present: present,
	}
}

func (o *DDLBase) Exists() bool {
	var id string

	if err := SelectRow(o.tx, fmt.Sprintf("SELECT TOP 1 'X' AS id FROM %s", o.table)).Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return true
		}

		return false
	}

	return true
}

func (o *DDLBase) Drop() error {
	query := fmt.Sprintf("ALTER TABLE %s RENAME TO  %s ;", o.table, fmt.Sprintf("_%s", o.table))
	if _, err := Exec(o.tx, query); err != nil {
		return err
	}

	return nil
}

type MySqlDDL struct {
	*DDLBase
}

func (o *MySqlDDL) IsSupportSequence() bool {
	return true
}

func (o *MySqlDDL) Desc() (cols []string, present map[string]string, err error) {
	res, err := Select(o.tx, fmt.Sprintf("DESC %s", o.table))
	if err != nil {
		return nil, nil, err
	}

	present, cols = base.ResAsMapSlice(res, false, "Field", "Type")
	return
}

func (o *MySqlDDL) Create() error {
	buf := new(bytes.Buffer)

	buf.WriteString(fmt.Sprintf("\n CREATE TABLE %s ( \n", o.table))
	for _, col := range o.cols {
		if strings.EqualFold(col, "id") {
			buf.WriteString(fmt.Sprintf("\t %s %s COLLATE utf8mb4_general_ci NOT NULL, \n", col, o.present[col]))
		} else {
			buf.WriteString(fmt.Sprintf("\t %s %s COLLATE utf8mb4_general_ci DEFAULT NULL, \n", col, o.present[col]))
		}
	}

	buf.WriteString("\t PRIMARY KEY (`id`) \n")
	buf.WriteString(" ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;")

	if _, err := Exec(o.tx, buf.String()); err != nil {
		return err
	}

	return nil
}

func (o *MySqlDDL) Alter(added, changed, removed map[string]string) error {
	buf := new(bytes.Buffer)

	syntax := make([]string, 0, len(added)+len(changed)+len(removed))
	buf.WriteString(fmt.Sprintf("\n ALTER TABLE %s ", o.table))

	var lastCol string
	for _, col := range o.cols {
		// 添加
		if _, ok := added[col]; ok {
			if len(lastCol) < 1 {
				syntax = append(syntax, fmt.Sprintf("\n\t ADD COLUMN %s %s COLLATE utf8mb4_general_ci DEFAULT NULL FIRST", col, o.present[col]))
			} else {
				syntax = append(syntax, fmt.Sprintf("\n\t ADD COLUMN %s %s COLLATE utf8mb4_general_ci DEFAULT NULL AFTER %s", col, o.present[col], lastCol))
			}
		}

		// 修改
		if _, ok := changed[col]; ok {
			if len(lastCol) < 1 {
				syntax = append(syntax, fmt.Sprintf("\n\t CHANGE COLUMN %s %s %s COLLATE utf8mb4_general_ci DEFAULT NULL FIRST", col, col, o.present[col]))
			} else {
				syntax = append(syntax, fmt.Sprintf("\n\t CHANGE COLUMN %s %s %s COLLATE utf8mb4_general_ci DEFAULT NULL AFTER %s", col, col, o.present[col], lastCol))
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

	if _, err := Exec(o.tx, buf.String()); err != nil {
		return err
	}

	return nil
}

type DmDDL struct {
	*DDLBase
}

func (o *DmDDL) IsSupportSequence() bool {
	return true
}

func (o *DmDDL) Desc() (cols []string, present map[string]string, err error) {
	query := `
		SELECT column_name AS column_name,
			CASE data_type 
				WHEN 'VARCHAR' THEN CONCAT('varchar','(',CAST(data_length AS varchar),')') 
				WHEN 'NUMERIC' THEN CONCAT('decimal','(',CAST(data_precision AS varchar),',',CAST(data_scale AS varchar),')') 
				ELSE lower(data_type)
			END AS data_type 
		FROM user_tab_columns 
		WHERE table_name = ?
		`
	res, err := Select(o.tx, query, o.table)
	if err != nil {
		return nil, nil, err
	}

	present, cols = base.ResAsMapSlice(res, false, "column_name", "data_type")
	return
}

func (o *DmDDL) Create() error {
	buf := new(bytes.Buffer)

	buf.WriteString(fmt.Sprintf("\n CREATE TABLE %s ( \n", o.table))
	for _, col := range o.cols {
		if strings.EqualFold(col, "id") {
			buf.WriteString(fmt.Sprintf("\t %s %s NOT NULL, \n", col, o.present[col]))
		} else {
			buf.WriteString(fmt.Sprintf("\t %s %s NULL, \n", col, o.present[col]))
		}
	}

	buf.WriteString(fmt.Sprintf("\t CONSTRAINT %s_PK PRIMARY KEY (id) \n", o.table))
	buf.WriteString(" );")

	if _, err := Exec(o.tx, buf.String()); err != nil {
		return err
	}

	return nil
}

func (o *DmDDL) Alter(added, changed, removed map[string]string) error {
	// 添加
	for _, col := range o.cols {
		if _, ok := added[col]; ok {
			if _, err := Exec(o.tx, fmt.Sprintf("ALTER TABLE %s ADD %s %s NULL;", o.table, col, o.present[col])); err != nil {
				return err
			}
		}

		// 修改，改变数据列的数据类型
		if _, ok := changed[col]; ok {
			if _, err := Exec(o.tx, fmt.Sprintf("ALTER TABLE %s MODIFY %s %s NULL;", o.table, col, o.present[col])); err != nil {
				return err
			}
		}
	}

	// 删除，不实际删除列
	for field := range removed {
		if _, err := Exec(o.tx, fmt.Sprintf("ALTER TABLE %s RENAME COLUMN %s TO %s;", o.table, field, fmt.Sprintf("_%s", field))); err != nil {
			return err
		}
	}

	return nil
}

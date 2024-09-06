package asql

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-phoenix/base"
)

type DDL interface {
	Exists() bool            // 判断表是否存在
	IsSupportSequence() bool // 是否支持修改列顺序

	Desc() ([]string, map[string]string, error)                          // 获取数据库表结构描述信息
	Create() error                                                       // 创建数据库表结构
	Alter(map[string]string, map[string]string, map[string]string) error // 修改数据库表结构
	Drop() error                                                         // 删除数据库表

	LimitOffset(int, int) string // 分页语法
}

func NewDDL(tx *sql.Tx, table string, cols []string, present map[string]string) DDL {
	ddl := NewDDLBase(tx, table, cols, present)

	switch base.Config.DBDriver {
	case "mysql":
		return &MySqlDDL{DDLBase: ddl}
	case "dm":
		return &DmDDL{DDLBase: ddl}
	case "mssql":
		return &MsSqlDDL{DDLBase: ddl}
	default:
		logrus.Panicf("unsupport database driver %q", base.Config.DBDriver)
	}

	return nil
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

package asql

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-phoenix/base"
	"strings"
)

type DDL struct {
	tx *sql.Tx

	table   string            // Table
	columns map[string]string // Physical  Columns
}

func NewDDL(tx *sql.Tx, table string) (*DDL, error) {

	// 查询实际定义的字段和类型
	columns := make(map[string]string)
	fields := make([]string, 0)
	sColumns, err := Select(tx, fmt.Sprintf("DESC %s", table))
	if err != nil {
		// 极大概率是因为不存在表
		logrus.Errorf("Get Description Table %s ERROR :: %s , Treat as Table not Exists .", table, err.Error())
	} else {
		columns, fields = base.ResAsMapSlice(sColumns, false, "Field", "Type")

		// 去掉以_开头的字段
		for _, field := range fields {
			if strings.HasPrefix(field, "_") {
				delete(columns, field)
			}

			// 在数据库中decimal(13,2)和numeric(13,2)是等效，统一用numeric处理
			if strings.Contains(columns[field], "decimal") {
				columns[field] = strings.ReplaceAll(columns[field], "decimal", "numeric")
			}
		}
	}

	logrus.Debugf("Columns %#v", columns)

	return &DDL{tx: tx, table: table, columns: columns}, nil
}

package asql

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-phoenix/base"
	"strings"
)

type MsSqlDDL struct {
	*DDLBase
}

func (o *MsSqlDDL) IsSupportSequence() bool {
	return false
}

func (o *MsSqlDDL) Desc() (cols []string, present map[string]string, err error) {
	query := `
		SELECT X1.name AS column_name,
			CASE X2.name 
				WHEN 'varchar' THEN X2.name+'('+CONVERT(VARCHAR(10),X1.length)+')' 
				WHEN 'numeric' THEN X2.name+'('+CONVERT(VARCHAR(2),X1.xprec)+','+CONVERT(VARCHAR(2),X1.xscale)+')' 
				ELSE X2.name
			END AS data_type
		FROM dbo.sysobjects T 
			INNER JOIN dbo.syscolumns X1 ON X1.id = T.id
			INNER JOIN dbo.systypes X2 ON X2.xusertype = X1.xtype
		WHERE T.id = X1.id AND T.xtype = 'U' AND T.status >= 0 AND T.name = ?
		ORDER BY X1.colid ASC
	`

	res, err := Select(o.tx, query, o.table)
	if err != nil {
		return nil, nil, err
	}

	present, cols = base.ResAsMapSlice(res, false, "column_name", "data_type")
	return
}

func (o *MsSqlDDL) Create() error {
	buf := new(bytes.Buffer)

	buf.WriteString(fmt.Sprintf("\n CREATE TABLE %s ( \n", o.table))
	for _, col := range o.cols {
		if strings.EqualFold(col, "id") {
			buf.WriteString(fmt.Sprintf("\t %s %s NOT NULL, \n", col, o.present[col]))
		} else {
			buf.WriteString(fmt.Sprintf("\t %s %s  DEFAULT NULL, \n", col, o.present[col]))
		}
	}

	buf.WriteString(fmt.Sprintf("\t CONSTRAINT %s_PK PRIMARY KEY (id) \n", o.table))
	buf.WriteString(" ) ;")

	if _, err := Exec(o.tx, buf.String()); err != nil {
		return err
	}

	return nil
}

func (o *MsSqlDDL) Alter(added, changed, removed map[string]string) error {
	for _, col := range o.cols {
		// 添加
		if _, ok := added[col]; ok {
			if _, err := Exec(o.tx, fmt.Sprintf("ALTER TABLE %s ADD %s %s NULL;", o.table, col, o.present[col])); err != nil {
				return err
			}
		}

		// 修改
		if aaa, ok := changed[col]; ok {
			logrus.Infof("%s =>  %s :: %s", col, aaa, o.present[col])
			if _, err := Exec(o.tx, fmt.Sprintf("ALTER TABLE %s ALTER COLUMN %s %s NULL;", o.table, col, o.present[col])); err != nil {
				return err
			}
		}
	}

	// 删除（不实际删除列），并且移动至表最后
	for field := range removed {
		if strings.EqualFold(field, "id") {
			continue
		}

		if _, err := Exec(o.tx, fmt.Sprintf("EXEC sp_rename '%s.%s', '%s', 'COLUMN';", o.table, field, fmt.Sprintf("_%s", field))); err != nil {
			return err
		}
	}

	return nil
}

func (o *MsSqlDDL) LimitOffset(start, count int) string {
	return fmt.Sprintf("OFFSET %d ROWS FETCH NEXT %d ROWS ONLY", start, count)
}

func (o *MsSqlDDL) Drop() error {
	query := fmt.Sprintf("EXEC sp_rename %s,%s ;", o.table, fmt.Sprintf("_%s", o.table))
	if _, err := Exec(o.tx, query); err != nil {
		return err
	}

	return nil
}

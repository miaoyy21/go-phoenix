package asql

import (
	"bytes"
	"fmt"
	"go-phoenix/base"
	"strings"
)

type DmDDL struct {
	*DDLBase
}

func (o *DmDDL) IsSupportSequence() bool {
	return false
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
		if strings.EqualFold(field, "id") {
			continue
		}

		if _, err := Exec(o.tx, fmt.Sprintf("ALTER TABLE %s RENAME COLUMN %s TO %s;", o.table, field, fmt.Sprintf("_%s", field))); err != nil {
			return err
		}
	}

	return nil
}

func (o *DmDDL) LimitOffset(start, count int) string {
	return fmt.Sprintf("LIMIT %d,%d", start, count)
}

func (o *DmDDL) Drop() error {
	query := fmt.Sprintf("ALTER TABLE %s RENAME TO  %s ;", o.table, fmt.Sprintf("_%s", o.table))
	if _, err := Exec(o.tx, query); err != nil {
		return err
	}

	return nil
}

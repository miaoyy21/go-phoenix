package asql

import (
	"bytes"
	"fmt"
	"go-phoenix/base"
	"strings"
)

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
		if strings.EqualFold(field, "id") {
			continue
		}

		syntax = append(syntax, fmt.Sprintf("\n\t CHANGE COLUMN %s %s %s COLLATE utf8mb4_general_ci DEFAULT NULL AFTER %s", field, fmt.Sprintf("_%s", field), xType, lastCol))
	}

	buf.WriteString(strings.Join(syntax, ","))
	buf.WriteByte(';')

	if _, err := Exec(o.tx, buf.String()); err != nil {
		return err
	}

	return nil
}

package asql

import (
	"bytes"
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

func AutoNo(tx *sql.Tx, kindCode string, num int, values map[string]string) ([]string, error) {
	query := `
			SELECT sys_auto_no_item.kind_id_, sys_auto_no_item.code_, sys_auto_no_item.value_
			FROM sys_auto_no_item 
				LEFT JOIN sys_auto_no_kind ON sys_auto_no_item.kind_id_ = sys_auto_no_kind.id
			WHERE sys_auto_no_kind.code_ = ?
			ORDER BY sys_auto_no_item.order_ ASC
		`
	res, err := Select(tx, query, kindCode)
	if err != nil {
		return nil, err
	}

	allNos := make([]string, 0, num)

	buf := &bytes.Buffer{}
	for _, s := range res {
		kindId, code, value := s["kind_id_"], s["code_"], s["value_"]
		switch code {
		case "STRING":
			if len(allNos) > 0 {
				for idx := range allNos {
					allNos[idx] = fmt.Sprintf("%s%s", allNos[idx], value)
				}
			} else {
				buf.WriteString(value)
			}
		case "VALUES":
			newValue, ok := values[value]
			if !ok {
				return nil, fmt.Errorf("缺少可变参数%s", value)
			}

			if len(allNos) > 0 {
				for idx := range allNos {
					allNos[idx] = fmt.Sprintf("%s%s", allNos[idx], newValue)
				}
			} else {
				buf.WriteString(newValue)
			}
		case "DATETIME":
			if len(allNos) > 0 {
				for idx := range allNos {
					allNos[idx] = fmt.Sprintf("%s%s", allNos[idx], time.Now().Format(value))
				}
			} else {
				buf.WriteString(time.Now().Format(value))
			}
		case "SEQ":
			var seq string

			prefix := buf.String()
			row := SelectRow(tx, "SELECT value_ FROM sys_auto_no WHERE kind_id_ = ? AND prefix_ = ?", kindId, prefix)
			if err := row.Scan(&seq); err != nil {
				if err != sql.ErrNoRows {
					return nil, err
				}

				// 转换为整数
				index, err := strconv.Atoi(value)
				if err != nil {
					return nil, err
				}

				index = index + num
				query := "INSERT INTO sys_auto_no(id, kind_id_, prefix_, value_, create_at_) VALUES (?,?,?,?,?)"
				args := []interface{}{GenerateId(), kindId, buf.String(), index, GetNow()}
				if err := Insert(tx, query, args...); err != nil {
					return nil, err
				}

				format := fmt.Sprintf("%%0%dd", len(value))
				for idx := 0; idx < num; idx++ {
					allNos = append(allNos, fmt.Sprintf("%s%s", prefix, fmt.Sprintf(format, index+idx)))
				}
			} else {
				index, err := strconv.Atoi(seq)
				if err != nil {
					return nil, err
				}

				index = index + num
				query := "UPDATE sys_auto_no SET value_ = ?, update_at_ = ? WHERE kind_id_ = ? AND prefix_ = ?"
				args := []interface{}{index, GetNow(), kindId, buf.String()}
				if err := Update(tx, query, args...); err != nil {
					return nil, err
				}

				format := fmt.Sprintf("%%0%dd", len(value))
				for idx := 1; idx <= num; idx++ {
					allNos = append(allNos, fmt.Sprintf("%s%s", prefix, fmt.Sprintf(format, index+idx)))
				}
			}
		}
	}

	if len(allNos) != num {
		return nil, fmt.Errorf("generate %d No,but Expect %d No", len(allNos), num)
	}

	return allNos, nil
}

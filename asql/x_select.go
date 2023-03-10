package asql

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"time"
)

func Select(tx *sql.Tx, query string, args ...interface{}) ([]map[string]string, error) {
	logrus.Debugf("Select SQL with Arguments [%s] %s", Arguments(args...), query)

	// Rows
	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Columns
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	values := make([]sql.RawBytes, len(columns))
	valuePts := make([]interface{}, len(columns))
	for i := 0; i < len(columns); i++ {
		valuePts[i] = &values[i]
	}

	entries := make([]map[string]string, 0)
	for rows.Next() {
		if err := rows.Scan(valuePts...); err != nil {
			return nil, err
		}

		entry := make(map[string]string)
		for i, col := range columns {
			if values[i] != nil {
				value := string(values[i])

				// 隐式转换时间格式
				if len(value) == len(time.RFC3339) {
					dt, err := time.Parse(time.RFC3339, value)
					if err == nil {
						if dt.Hour()+dt.Minute()+dt.Second() == 0 {
							value = dt.Format("2006-01-02")
						} else {
							value = dt.Format("2006-01-02 15:04:05")
						}
					}
				}

				entry[col] = value
			}
		}

		entries = append(entries, entry)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}

func SelectRow(tx *sql.Tx, query string, args ...interface{}) *sql.Row {
	logrus.Debugf("Get Row SQL with Arguments [%s] %s", Arguments(args...), query)

	return tx.QueryRow(query, args...)
}

func SelectColumns(tx *sql.Tx, query string, args ...interface{}) ([]string, error) {
	logrus.Debugf("SelectColumns SQL with Arguments [%s] %s", Arguments(args...), query)

	// Rows
	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Columns
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	return columns, nil
}

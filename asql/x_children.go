package asql

import (
	"database/sql"
	"fmt"
	"strings"
)

func QueryChildren(tx *sql.Tx, table string, ids []interface{}) ([]interface{}, error) {
	if len(ids) == 0 {
		return []interface{}{}, nil
	}

	query := fmt.Sprintf("SELECT id FROM %s WHERE parent_id_ IN (?%s) ", table, strings.Repeat(", ?", len(ids)-1))
	res, err := Select(tx, query, ids...)
	if err != nil {
		return nil, err
	}

	ns := make([]interface{}, 0, len(res))
	for _, kv := range res {
		id, ok := kv["id"]
		if ok {
			ns = append(ns, id)
		}
	}

	ms, err := QueryChildren(tx, table, ns)
	if err != nil {
		return nil, err
	}

	news := make([]interface{}, 0, len(ids)+len(ms))
	news = append(news, ids...)
	news = append(news, ms...)

	return news, nil
}

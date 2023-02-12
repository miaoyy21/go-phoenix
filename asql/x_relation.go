package asql

import (
	"database/sql"
	"fmt"
	"strings"
)

// QueryRelationChildren 根据ParentId查询对应的Id，返回的Id列表
func QueryRelationChildren(tx *sql.Tx, table string, ids []interface{}) ([]interface{}, error) {
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

	ms, err := QueryRelationChildren(tx, table, ns)
	if err != nil {
		return nil, err
	}

	news := make([]interface{}, 0, len(ids)+len(ms))
	news = append(news, ids...)
	news = append(news, ms...)

	return news, nil
}

// QueryRelationParents 根据Id查询ParentId，返回的ParentId列表
func QueryRelationParents(tx *sql.Tx, table string, id string) ([]interface{}, error) {
	var parentId sql.NullString

	if err := SelectRow(tx, fmt.Sprintf("SELECT parent_id_ FROM %s WHERE id = ? ", table), id).Scan(&parentId); err != nil {
		if err == sql.ErrNoRows {
			return []interface{}{}, nil
		}

		return nil, err
	}

	if !parentId.Valid {
		return []interface{}{id}, nil
	}

	supers, err := QueryRelationParents(tx, table, parentId.String)
	if err != nil {
		if err == sql.ErrNoRows {
			return []interface{}{id}, nil
		}

		return nil, err
	}

	news := make([]interface{}, 0, len(supers)+1)
	news = append(news, id)
	news = append(news, supers...)

	return news, nil
}

package vm

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"go-phoenix/asql"
	"go-phoenix/handle"
)

type Sql struct {
	tx  *sql.Tx
	ctx *handle.Context
}

func NewSql(tx *sql.Tx, ctx *handle.Context) *Sql {
	return &Sql{tx, ctx}
}

// AutoNo 自动编码
func (s *Sql) AutoNo(code string, values map[string]string) string {
	ano, err := asql.AutoNo(s.tx, code, values)
	if err != nil {
		logrus.Panic(err)
	}

	return ano
}

// Query 执行查询SQL
func (s *Sql) Query(query string, args ...interface{}) []map[string]string {
	res, err := asql.Select(s.tx, query, args...)
	if err != nil {
		logrus.Panic(err)
	}

	return res
}

// QueryRow 执行查询单结果SQL
func (s *Sql) QueryRow(query string, args ...interface{}) map[string]string {
	res, err := asql.Select(s.tx, query, args...)
	if err != nil {
		logrus.Panic(err)
	}

	if len(res) == 0 {
		return make(map[string]string)
	}

	return res[0]
}

// Exec 执行 插入、更新、删除 等SQL
func (s *Sql) Exec(query string, args ...interface{}) int64 {
	n, err := asql.Exec(s.tx, query, args...)
	if err != nil {
		logrus.Panic(err)
	}

	return n
}

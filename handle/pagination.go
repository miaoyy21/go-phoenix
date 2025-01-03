package handle

import (
	"database/sql"
	"fmt"
	"go-phoenix/asql"
	"strconv"
	"strings"
)

// Pagination 自带过滤与排序的分页器
type Pagination struct {
	start int
	count int

	sorts        []string
	filters      []string
	filtered     []interface{}
	fullFilters  []string      // TODO
	fullFiltered []interface{} // TODO

	// 直接导出为JSON
	Pos        int                 `json:"pos"`
	TotalCount int                 `json:"total_count"`
	Data       []map[string]string `json:"data"`
}

func NewPagination(ctx *Context, defaultSorts []string) *Pagination {
	start, _ := strconv.Atoi(ctx.FormValue("start"))
	count, _ := strconv.Atoi(ctx.FormValue("count"))

	count = count * 5
	if count < 1 {
		count = 250
	}

	// 如果没进行排序，那么使用默认排序
	sorts, filters, filtered, fullFilters, fullFiltered := ctx.SortFilters(map[string]string{})
	if len(sorts) < 1 && len(defaultSorts) > 0 {
		sorts = append(sorts, defaultSorts...)
	}

	return &Pagination{
		start:        start,
		count:        count,
		sorts:        sorts,
		filters:      filters,
		filtered:     filtered,
		fullFilters:  fullFilters,
		fullFiltered: fullFiltered,
		Pos:          start,
	}
}

func (p *Pagination) setRowCount(tx *sql.Tx, query string) error {
	var totalCount int

	// 构建新的SQL语句
	newQuery := fmt.Sprintf("SELECT COUNT(1) FROM (%s) _PHOENIX_", query)
	if len(p.filters) > 0 {
		newQuery = fmt.Sprintf("SELECT COUNT(1) FROM (%s) _PHOENIX_ WHERE %s", query, strings.Join(p.filters, " AND "))
	}

	if err := asql.SelectRow(tx, newQuery, p.filtered...).Scan(&totalCount); err != nil {
		return err
	}

	p.TotalCount = totalCount
	return nil
}

func (p *Pagination) SetData(tx *sql.Tx, query string, args ...interface{}) error {
	// 计算总记录数量
	if err := p.setRowCount(tx, query); err != nil {
		return err
	}

	ddl := asql.NewDDL(tx)
	newQuery := fmt.Sprintf(" %sORDER BY %s \n\t\t%s ", query, strings.Join(p.sorts, ","), ddl.LimitOffset(p.start, p.count))
	if len(p.filters) > 0 {
		newQuery = fmt.Sprintf("SELECT * FROM (%s) _PHOENIX_ WHERE %s ORDER BY %s \n\t\t %s ", query, strings.Join(p.filters, " AND "), strings.Join(p.sorts, ","), ddl.LimitOffset(p.start, p.count))
	}

	args = append(args, p.filtered...)
	data, err := asql.Select(tx, newQuery, args...)
	if err != nil {
		return err
	}

	p.Data = data
	return nil
}

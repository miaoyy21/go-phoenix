package xsys

import (
	"database/sql"
	"go-phoenix/asql"
	"go-phoenix/handle"
	"strings"
)

type SysTimeTasks struct {
}

func (o *SysTimeTasks) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	query := `
		SELECT id, name_, type_, once_at_, description_, is_start_, 
			frequency_, frequency_day_, frequency_day_repeat_, frequency_day_repeat_unit_,
			frequency_day_start_, frequency_day_end_, frequency_start_at_, frequency_end_at_,
			create_at_, update_at_
		FROM sys_time_task 
		ORDER BY order_ ASC
	`

	return asql.Select(tx, query)
}

func (o *SysTimeTasks) GetSource(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	id := ctx.FormValue("id")
	scope := ctx.FormValue("scope")

	query, args := "invalid", make([]interface{}, 0)

	query = "SELECT id,name_,type_,description_,is_start_,create_at_,update_at_ FROM sys_time_task ORDER BY order_ ASC"
	if strings.EqualFold(scope, "EXTRA") && len(id) > 0 {
		query = `
			SELECT once_at_, frequency_, frequency_day_, 
				frequency_day_repeat_, frequency_day_repeat_unit_,
				frequency_day_start_, frequency_day_end_, 
				frequency_start_at_, frequency_end_at_
			FROM sys_time_task
			WHERE id = ?
			ORDER BY order_ ASC
		`
		args = append(args, id)
	}

	return asql.Select(tx, query, args...)
}

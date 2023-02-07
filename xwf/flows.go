package xwf

import (
	"database/sql"
	"go-phoenix/asql"
	"go-phoenix/handle"
)

type Flows struct {
}

func (r *Flows) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	//id := ctx.FormValue("id")
	//scope := ctx.FormValue("scope")

	query, args := "invalid", make([]interface{}, 0)
	//if len(id) > 0 && strings.EqualFold(scope, "EXTRA") {
	//	query = "SELECT model_, options_ FROM wf_diagram WHERE id = ?"
	//	args = append(args, id)
	//} else {
	//	query = `
	//		SELECT id, code_, name_, description_, create_at_, update_at_
	//		FROM wf_diagram
	//		ORDER BY order_ ASC
	//	`
	//}

	res, err := asql.Select(tx, query, args...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

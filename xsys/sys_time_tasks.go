package xsys

import (
	"database/sql"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/handle"
)

type SysTimeTasks struct {
}

func (o *SysTimeTasks) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	query := `
		SELECT id, name_, type_, description_, is_start_, 
			create_at_, update_at_ 
		FROM sys_time_task 
		ORDER BY order_ ASC
	`

	return asql.Select(tx, query)
}

func (o *SysTimeTasks) GetExtra(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	id := ctx.FormValue("id")

	query, args := "invalid", make([]interface{}, 0)
	if len(id) > 0 {
		query = `
			SELECT source_, once_at_, frequency_, frequency_day_, 
				frequency_day_repeat_, frequency_day_repeat_unit_,
				frequency_day_start_, frequency_day_end_, 
				frequency_start_at_, frequency_end_at_
			FROM sys_time_task
			WHERE id = ?
		`
		args = append(args, id)
	}

	return asql.Select(tx, query, args...)
}

func (o *SysTimeTasks) Post(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	operation := ctx.PostFormValue("operation")

	id := ctx.PostFormValue("id")
	name := ctx.PostFormValue("name_")
	xType := ctx.PostFormValue("type_")
	description := ctx.PostFormValue("description_")

	source := ctx.PostFormValue("source_")
	onceAt := ctx.PostFormValue("once_at_")
	frequency := ctx.PostFormValue("frequency_")
	frequencyDay := ctx.PostFormValue("frequency_day_")
	frequencyDayRepeat := ctx.PostFormNullableValue("frequency_day_repeat_")
	frequencyDayRepeatUnit := ctx.PostFormValue("frequency_day_repeat_unit_")
	frequencyDayStart := ctx.PostFormValue("frequency_day_start_")
	frequencyDayEnd := ctx.PostFormValue("frequency_day_end_")
	frequencyStartAt := ctx.PostFormValue("frequency_start_at_")
	frequencyEndAt := ctx.PostFormValue("frequency_end_at_")

	moveId := ctx.PostFormValue("webix_move_id")
	moveIndex := ctx.PostFormValue("webix_move_index")
	moveParent := ctx.PostFormValue("webix_move_parent")

	now := asql.GetNow()
	switch operation {
	case "insert":
		newId := asql.GenerateId()
		query := `
			INSERT INTO sys_time_task(
				id, name_, type_, description_, 
				source_, once_at_, frequency_, frequency_day_, 
				frequency_day_repeat_, frequency_day_repeat_unit_,
				frequency_day_start_, frequency_day_end_, 
				frequency_start_at_, frequency_end_at_, 
				order_, create_at_
			) VALUES (?,?,?,?,?, ?,?,?,?, ?,?, ?,?, ?,?, ?,?)`
		args := []interface{}{
			newId, name, xType, description,
			source, onceAt, frequency, frequencyDay,
			frequencyDayRepeat, frequencyDayRepeatUnit,
			frequencyDayStart, frequencyDayEnd,
			frequencyStartAt, frequencyEndAt,
			asql.GenerateOrderId(), now,
		}
		if err := asql.Insert(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "newid": newId, "create_at_": now}, nil
	case "update":
		query := `
			UPDATE sys_time_task 
			SET name_ = ?, type_ = ?, description_ = ?, 
				source_ = ?, once_at_ = ?, frequency_ = ?, frequency_day_ = ?,
				frequency_day_repeat_ = ?, frequency_day_repeat_unit_ = ?,
				frequency_day_start_ = ?, frequency_day_end_ = ?,
				frequency_start_at_ = ?, frequency_end_at_ = ?,
				update_at_ = ? 
			WHERE id = ?
		`
		args := []interface{}{
			name, xType, description,
			source, onceAt, frequency, frequencyDay,
			frequencyDayRepeat, frequencyDayRepeatUnit,
			frequencyDayStart, frequencyDayEnd,
			frequencyStartAt, frequencyEndAt,
			now, id,
		}
		if err := asql.Update(tx, query, args...); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "id": id, "update_at_": now}, nil
	case "delete":
		if err := asql.Delete(tx, "DELETE FROM sys_time_task WHERE id = ?", id); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success", "id": id}, nil
	case "order":
		if err := asql.Order(tx, "sys_time_task", id, moveId, moveIndex, moveParent); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success"}, nil
	}

	return nil, fmt.Errorf("unrecognizable operation %s ", operation)
}

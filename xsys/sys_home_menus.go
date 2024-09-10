package xsys

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-phoenix/asql"
	"go-phoenix/base"
	"go-phoenix/handle"
)

type SysHomeMenus struct {
}

func (o *SysHomeMenus) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	query := "SELECT T.id AS id, T.menu_id_ AS menu_id_, X.name_ AS menu_name_, X.icon_ AS menu_icon_ FROM sys_home_menu T INNER JOIN sys_menu X ON X.id = T.menu_id_ WHERE T.user_id_ = ? ORDER BY T.order_ ASC"

	return asql.Select(tx, query, ctx.UserId())
}

func (o *SysHomeMenus) Post(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	operation := ctx.PostFormValue("operation")

	switch operation {
	case "save":
		rawMenus := ctx.PostFormValue("menus")

		menus := make([]string, 0)
		if err := json.Unmarshal([]byte(rawMenus), &menus); err != nil {
			return nil, err
		}

		newMenus := make(map[string]string)
		for _, menu := range menus {
			newMenus[menu] = "NULL"
		}

		res, err := asql.Select(tx, "SELECT menu_id_, 'NULL' AS empty_ FROM sys_home_menu WHERE user_id_ = ?", ctx.UserId())
		if err != nil {
			return nil, err
		}

		// 两个Map进行比较
		oldMenus := base.ResAsMap(res, true, "menu_id_", "empty_")
		added, _, removed := base.CompareMap(oldMenus, newMenus)

		now := asql.GetNow()

		// 新增
		for menu := range added {
			query := "INSERT INTO sys_home_menu(id, user_id_, menu_id_, order_, create_at_) VALUES (?,?,?,?,?)"
			args := []interface{}{asql.GenerateId(), ctx.UserId(), menu, asql.GenerateOrderId(), now}

			if err := asql.Insert(tx, query, args...); err != nil {
				return nil, err
			}
		}

		// 删除
		for menu := range removed {
			if err := asql.Delete(tx, "DELETE FROM sys_home_menu WHERE user_id_ = ? AND menu_id_ = ?", ctx.UserId(), menu); err != nil {
				return nil, err
			}
		}

		return map[string]interface{}{"status": "success"}, nil
	case "order":
		id := ctx.PostFormValue("id")

		moveId := ctx.PostFormValue("webix_move_id")
		moveIndex := ctx.PostFormValue("webix_move_index")
		moveParent := ctx.PostFormValue("webix_move_parent")

		if err := asql.Order(tx, "sys_home_menu", id, moveId, moveIndex, moveParent); err != nil {
			return nil, err
		}

		return map[string]interface{}{"status": "success"}, nil
	}

	return nil, fmt.Errorf("unrecognizable operation %s ", operation)
}

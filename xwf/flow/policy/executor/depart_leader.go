package flow

import (
	"database/sql"
	"go-phoenix/asql"
	"go-phoenix/base"
)

func GetDepartLeaderByDepartId(tx *sql.Tx, departId string) ([]string, error) {
	query := "SELECT id FROM sys_user WHERE depart_id_ = ? AND is_depart_leader_ = 'Yes' ORDER BY order_ ASC"
	res, err := asql.Select(tx, query, departId)
	if err != nil {
		return nil, err
	}

	// 如果本部门有部门领导，那么返回这些部门领导，否则取上级部门领导
	leaders := base.ResAsSliceString(res, "id")
	if len(leaders) > 0 {
		return leaders, nil
	}

	var parentDepartId string
	erx := asql.SelectRow(tx, "SELECT parent_id_ FROM sys_depart WHERE id = ?", departId).Scan(&parentDepartId)
	if erx != nil {
		if erx == sql.ErrNoRows {
			return make([]string, 0), nil
		}

		return nil, err
	}

	parentLeaders, err := GetDepartLeaderByDepartId(tx, parentDepartId)
	if err != nil {
		return nil, err
	}

	return parentLeaders, nil
}

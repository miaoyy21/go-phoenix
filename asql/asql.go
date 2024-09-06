package asql

import (
	"sync"
)

type SysColumn struct {
	SaveIgnore bool

	Code        string
	Name        string
	Type        string
	Description string
}

var SysColumns = []SysColumn{
	{true, "id", "ID", "VARCHAR(32)", "【系统自动创建】该条记录的全局唯一ID"},

	{true, "parent_id_", "父级ID", "VARCHAR(32)", "【系统自动创建】自动建树服务的父级ID"},
	{true, "order_", "排序号", "BIGINT", "【系统自动创建】拖拽排序服务的排序号"},
	{true, "create_depart_id_", "部门ID", "VARCHAR(32)", "【系统自动创建】该条记录的创建者的部门ID"},
	{true, "create_depart_code_", "部门编码", "VARCHAR(32)", "【系统自动创建】该条记录的创建者的部门编码"},
	{true, "create_depart_name_", "部门名称", "VARCHAR(256)", "【系统自动创建】该条记录的创建者的部门名称"},
	{true, "create_user_id_", "用户ID", "VARCHAR(32)", "【系统自动创建】该条记录的创建者ID"},
	{true, "create_user_code_", "工号", "VARCHAR(32)", "【系统自动创建】该条记录的创建者工号"},
	{true, "create_user_name_", "用户名", "VARCHAR(32)", "【系统自动创建】该条记录的创建者用户名"},
	{true, "create_at_", "创建时间", "DATETIME", "【系统自动创建】该条记录的创建时间"},
	{true, "update_at_", "更新时间", "DATETIME", "【系统自动创建】该条记录的最新更新时间"},
}

var ignore = map[string]struct{}{}

var ignoreOnce = sync.Once{}

func IsSaveIgnore(field string) bool {
	ignoreOnce.Do(func() {
		for _, col := range SysColumns {
			if col.SaveIgnore {
				ignore[col.Code] = struct{}{}
			}
		}
	})

	_, ok := ignore[field]
	return ok
}
